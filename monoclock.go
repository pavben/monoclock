package monoclock

import (
	"syscall"
	"unsafe"
)

type MonoTimer struct {
	startTime int64
}

// Creates a new monotonic timer and starts it.
func NewMonoTimer() *MonoTimer {
	startTime := GetMonoTime()

	return &MonoTimer {
		startTime: startTime,
	}
}

// Returns the number of milliseconds that have elapsed since this MonoTimer was created.
func (self *MonoTimer) Value() int64 {
	currentTime := GetMonoTime()

	return currentTime - self.startTime
}

// Returns the time of the system monotonic clock in milliseconds.
func GetMonoTime() int64 {
	sec, nsec := getRawMonoTime()

	// to milliseconds
	return sec * 1000 + (nsec / (1 * 1000 * 1000))
}

const CLOCK_MONOTONIC_RAW = 4

func getRawMonoTime() (int64, int64) {
	var timeSpec syscall.Timespec

	syscall.Syscall(syscall.SYS_CLOCK_GETTIME, CLOCK_MONOTONIC_RAW, uintptr(unsafe.Pointer(&timeSpec)), 0)

	sec, nsec := timeSpec.Unix()

	return sec, nsec
}
