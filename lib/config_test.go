package lib

import (
	"os"
	"testing"
)

const exampleConfig1 = `batteries:
    - name: ps-controller-battery-90:89:5f:2a:81:5c
      critical: 30
    - name: DELL FW8KR9A
      critical: 60
    - name: Bluetooth Mouse M336/M337/M535
      critical: 10
alerts:
    added: /usr/share/sounds/freedesktop/stereo/power-plug.oga
    removed: /usr/share/sounds/freedesktop/stereo/power-unplug.oga
    charging: /usr/share/sounds/freedesktop/stereo/power-plug.oga
    discharging: /usr/share/sounds/freedesktop/stereo/power-unplug.oga
    full: /usr/share/sounds/freedesktop/stereo/complete.oga
    critical: /usr/share/sounds/freedesktop/stereo/alarm-clock-elapsed.oga
logfile: /tmp/uap.log
verbose: true`

func check(t *testing.T, err error) {
	if err != nil {
		t.Errorf("(test) error: %s\n", err)
	}
}

func TestConfig(t *testing.T) {
	f, err := os.CreateTemp("/tmp", "config_test_")
	check(t, err)

	f.WriteString(exampleConfig1)
	rc, err := LoadConfig(f.Name())
	check(t, err)

	if len(rc.Batteries) != 3 {
		t.Errorf("expected/resulting batteries length: 3/%d", len(rc.Batteries))
	}

	os.Remove(f.Name())

}
