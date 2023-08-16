package lib

func AlertSet(config *Config, battery *Battery, state bool) {
	batteryConfig := config.Get(battery)
	if batteryConfig != nil {
		batteryConfig.Alerted = state
	}
}

func AlertIfCritical(config *Config, battery *Battery) {
	batteryConfig := config.Get(battery)
	if batteryConfig != nil {
		if battery.Api.Capacity <= batteryConfig.Critical &&
			!batteryConfig.Alerted {
			PlaySafe(config.Alerts.Critical)
			batteryConfig.Alerted = true
		}
	}
}

func AlertSound(config *Config, event *Event) {
	switch event.Kind {
	case Added:
		PlaySafe(config.Alerts.Added)
	case Removed:
		PlaySafe(config.Alerts.Removed)
	case Charging:
		PlaySafe(config.Alerts.Charging)
	case Discharging:
		PlaySafe(config.Alerts.Discharging)
	case Full:
		PlaySafe(config.Alerts.Full)
	case Idle:
		PlaySafe(config.Alerts.Idle)
	case Empty:
		PlaySafe(config.Alerts.Empty)
	}
}

func AlertEvent(config *Config, event *Event) {
	AlertSound(config, event)
	AlertNotify(config, event)
}
