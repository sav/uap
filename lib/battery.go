package lib

import (
	"container/list"
	"fmt"
	"time"

	impl "github.com/sav/battery"
)

type Snapshot struct {
	State impl.AgnosticState
	When  time.Time
}

type Battery struct {
	Id   int
	Name string
	Last Snapshot
	Api  *impl.Battery
}

type EventType int8

const (
	Charging EventType = iota
	Discharging
	Idle
	Empty
	Full
	Added
	Removed
)

var eventTypeString = map[EventType]string{
	Charging:    "Charging",
	Discharging: "Discharging",
	Idle:        "Idle",
	Empty:       "Empty",
	Full:        "Full",
	Added:       "Added",
	Removed:     "Removed",
}

func (eventType EventType) String() string {
	return eventTypeString[eventType]
}

type Event struct {
	Kind    EventType
	Battery *Battery
}

func (event Event) String() string {
	return fmt.Sprintf("%v: %#v", event.Kind, event.Battery)
}

func (battery *Battery) Glyph() string {
	var glyph = ""
	switch battery.Api.State.Raw {
	case impl.Charging:
		glyph = "🗲"
	case impl.Discharging:
		glyph = "🔋"
	case impl.Full:
		glyph = "😎"
	case impl.Idle:
		glyph = "⏲"
	case impl.Empty:
		glyph = "⚠"
	case impl.Unknown:
		glyph = "�"
	case impl.Undefined:
		glyph = "⁇"
	}
	return glyph
}

func (battery *Battery) String() string {
	return fmt.Sprintf("%s: %s %.1f%% (%s)",
		battery.Name, battery.Glyph(), battery.Capacity(), battery.Api.State)
}

func (battery *Battery) GoString() string {
	return fmt.Sprintf("%s", battery.Name)
}

func (battery *Battery) Charging() bool {
	return battery.Api.State.Raw == impl.Charging
}

func (battery *Battery) Discharging() bool {
	return battery.Api.State.Raw == impl.Discharging
}

func (battery *Battery) Full() bool {
	return battery.Api.State.Raw == impl.Full
}

func (battery *Battery) Capacity() float64 {
	return battery.Api.Capacity
}

func (battery *Battery) Compare(c float64) float64 {
	return c - battery.Capacity()
}

func (battery *Battery) Update() (bool, error) {
	var err error
	battery.Api, err = impl.Get(battery.Id)
	if battery.Api == nil && err != nil {
		return false, err
	}
	if battery.Api.State.Raw == battery.Last.State {
		return false, nil
	}
	battery.Last.State = battery.Api.State.Raw
	battery.Last.When = time.Now()
	return true, nil
}

func Update(batteries *list.List, callback func(Event)) {
	current, err := GetBatteries()
	if err != nil {
		panic(err)
	}

	ListMirror(batteries, current,
		func(a *Battery, b *Battery) bool {
			return a.Name == b.Name
		},
		func(event ListEvent[*Battery]) {
			switch event.Kind {
			case ElementAdded:
				callback(Event{Added, event.Element})
			case ElementRemoved:
				callback(Event{Removed, event.Element})
			}
		})

	for element := batteries.Front(); element != nil; element = element.Next() {
		battery := element.Value.(*Battery)
		changed, _ := battery.Update()
		if changed {
			switch battery.Api.State.Raw {
			case impl.Charging:
				callback(Event{Charging, battery})
			case impl.Discharging:
				callback(Event{Discharging, battery})
			case impl.Idle:
				callback(Event{Idle, battery})
			case impl.Empty:
				callback(Event{Empty, battery})
			case impl.Full:
				callback(Event{Full, battery})
			default:
				Log.Warn("Uknown battery state: %v", battery.Api.State.Raw)
			}
		}
	}
}

func newBattery(id int, api *impl.Battery) (*Battery, error) {
	if api == nil {
		var err error
		api, err = impl.Get(id)
		if err != nil {
			return nil, err
		}
	}
	return &Battery{
		id,
		api.Name,
		Snapshot{api.State.Raw, time.Now()},
		api,
	}, nil
}

func GetBatteries() (*list.List, error) {
	batteries, err := impl.GetAll()
	if err != nil {
		if fatal, ok := err.(impl.ErrFatal); ok {
			return nil, fatal
		}
	}
	var ready = list.New()
	for i, api := range batteries {
		if api.Full > 0 || api.Capacity > 0 {
			var battery, err = newBattery(i, api)
			if err != nil {
				Log.Err("Invalid battery: %d", i)
				continue
			}
			ready.PushBack(battery)
		}
	}
	return ready, nil
}
