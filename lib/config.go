package lib

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type BatteryConfig struct {
	Name     string  `yaml:"name"`
	Critical float64 `yaml:"critical"`
	Alerted  bool
}

type AlertConfig struct {
	Added       string `yaml:"added"`
	Removed     string `yaml:"removed"`
	Charging    string `yaml:"charging"`
	Discharging string `yaml:"discharging"`
	Full        string `yaml:"full"`
	Idle        string `yaml:"idle"`
	Empty       string `yaml:"empty"`
	Critical    string `yaml:"critical"`
}

type Config struct {
	Batteries []*BatteryConfig `yaml:"batteries"`
	Alerts    AlertConfig      `yaml:"alerts"`
	LogFile   string           `yaml:"logfile"`
	Verbose   bool             `yaml:"verbose"`
}

func (config *Config) Load(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		return err
	}
	Log.Info("Loaded configuration file: %s", filename)
	return nil
}

func LoadConfig(filename string) (Config, error) {
	config := Config{}
	return config, config.Load(filename)
}

func (config *Config) Get(battery *Battery) *BatteryConfig {
	for _, each := range config.Batteries {
		if battery.Name == each.Name {
			return each
		}
	}
	return nil
}
