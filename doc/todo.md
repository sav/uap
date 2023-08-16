
# To Do

## Version 0.1

   ---
   - **DONE** write list_test.go
   - **DONE** libnotify notification
   - **DONE** alert critical level
   - **DONE** create alert module
   - **DONE** alert added/removed batteries
   - **DONE** create package lib
   - **DONE** use lib from main
   - **DONE** verbose/debug arg
   - **DONE** add tests
   - **DONE** output log to file
   - **DONE** getopt/cmd-args
   - **DONE** config and ~/.uaprc
   - **DONE** play alarm sound
   - **DONE** check percentage
   - **DONE** ogg/vorbis support
   - **DONE** unicode rep of states
   - **DONE** detect plug/unplug

## Version 0.2

   - **TODO** log capacity change every `n-threshold'
   - **TODO** use the `slog' package instead of `log'
   - **TODO** syslog support
   - **TODO** c-rate, time to (dis)charge, based on capacity/mAh
   - **TODO** c-rate, based on capacity change over time
   - **TODO** call script on alert
   - **TODO** health (curr/design)
   - **TODO** historical health (sql)
   - **TODO** system tray
   - **TODO** simple gui panel
   - **TODO** test coverage report (plus ~50% coverage)
   - **TODO** interface{Next(),...} instead of *list.List
   - **TODO** add godocs

## Issues

   ---
   - **DONE** removed batteries lead to runtime errors
   - **DONE** newly added batteries are ignored


## Contrib


## distatus/battery

   - **TODO** attribute CAPACITY (windows)
   - **TODO** attribute NAME (windows)
   - **TODO** attribute CAPACITY (netbsd)
   - **TODO** attribute NAME (netbsd)
   ---
   - **DONE** attribute CAPACITY (linux)
   - **DONE** attribute NAME (linux)
