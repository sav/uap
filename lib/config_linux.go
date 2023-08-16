package lib

import "os"

func DefaultConfig() string {
	home := os.Getenv("HOME")
	return home + "/.uaprc"
}
