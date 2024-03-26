package helpers

import "runtime"

func Goos() string {
	return runtime.GOOS
}
