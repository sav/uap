package lib

import (
	"fmt"

	"github.com/codegoalie/golibnotify"
)

func AlertNotify(config *Config, event *Event) {
	notifier := golibnotify.NewSimpleNotifier("UAP")
	Log.ErrWrap(notifier.Show(
		fmt.Sprintf("Battery %s", event.Kind),
		fmt.Sprintf("%#v", event.Battery), ""))
}
