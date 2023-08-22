package main

import (
	"container/list"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	. "github.com/sav/uap/lib"
)

type Monitor func(*Config, *list.List)

type Status func(*Config, *list.List)

func start(
	config *Config,
	monitor Monitor,
	status Status,
) {
	tick := time.Tick(1500 * time.Millisecond)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, []os.Signal{
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGUSR1,
	}...)

	bats, err := GetBatteries()
	if err != nil {
		Log.Err("Could not get battery info: %v", err)
		os.Exit(1)
	}

	Log.Dbg("Starting UAP...")
	status(config, bats)

	var quit = false
	for !quit {
		select {
		case <-tick:
			monitor(config, bats)
		case sig := <-sigs:
			if sig == syscall.SIGUSR1 {
				status(config, bats)
			} else {
				Log.Dbg("Received signal: %v", sig)
				quit = true
			}
		default:
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func monitor(config *Config, batteries *list.List) {
	Update(batteries, func(event Event) {
		switch event.Kind {
		case Added:
			Log.Info("%v: ☑ (Added)", event.Battery.Name)
			Log.Info("%v", event.Battery)
		case Removed:
			Log.Info("%v: ☒ (Removed)", event.Battery.Name)
		default:
			Log.Info("%v", event.Battery)
		}
		// Don't reset when state's changing to either Charging or Removed.
		AlertSet(config, event.Battery,
			event.Kind == Charging || event.Kind == Removed)
		AlertEvent(config, &event)
	})
	ListMap(batteries, func(battery *Battery) {
		AlertIfCritical(config, battery)
	})
}

func status(_ *Config, batteries *list.List) {
	if batteries.Len() == 0 {
		Log.Info("No batteries found.")
	}
	ListMap(batteries, func(battery *Battery) {
		Log.Info("%v", battery)
	})
}

var (
	config     Config
	configFile string
	logFile    string
	verbose    bool
	showVers   bool
)

func init() {
	flag.StringVar(&configFile, "c", DefaultConfig(),
		"Load configuration from file.")
	flag.StringVar(&logFile, "l", "", "Write logs to file.")
	flag.BoolVar(&verbose, "v", false, "Enable verbose output.")
	flag.BoolVar(&showVers, "V", false, "Print version information.")
}

func main() {
	flag.Parse()

	if showVers {
		printVersion()
		os.Exit(0)
	}

	config, err := LoadConfig(configFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		config = Config{}
	}

	err = nil
	if logFile != "" {
		err = SetLogFile(logFile)
	} else if config.LogFile != "" {
		err = SetLogFile(config.LogFile)
	}
	if err != nil {
		panic(err)
	}
	SetLogDebug(verbose || config.Verbose)

	start(&config, monitor, status)
}
