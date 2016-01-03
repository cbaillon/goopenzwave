package goopenzwave

// #include "manager.h"
// #include <stdlib.h>
import "C"

// GetPollInterval returns the time period between polls of a node's state.
func GetPollInterval() int32 {
	return int32(C.manager_getPollInterval(cmanager))
}

// SetPollInterval will set the time period between polls of a node's state.
//
// Due to patent concerns, some devices do not report state changes
// automatically to the controller. These devices need to have their state
// polled at regular intervals. The length of the interval is the same for all
// devices. To even out the Z-Wave network traffic generated by polling,
// OpenZWave divides the polling interval by the number of devices that have
// polling enabled, and polls each in turn. It is recommended that if possible,
// the interval should not be set shorter than the number of polled devices in
// seconds (so that the network does not have to cope with more than one poll
// per second).
func SetPollInterval(milliseconds int32, intervalBetweenPolls bool) {
	C.manager_setPollInterval(cmanager, C.int32_t(milliseconds), C.bool(intervalBetweenPolls))
}

// EnablePoll enables the polling of a device's state. Returns true if polling
// was enabled.
func EnablePoll(homeID uint32, valueID uint64, intensity uint8) bool {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	return bool(C.manager_enablePoll(cmanager, cvalueid, C.uint8_t(intensity)))
}

// DisablePoll disables the polling of a device's state. Returns true if polling
// was disabled.
func DisablePoll(homeID uint32, valueID uint64) bool {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	return bool(C.manager_disablePoll(cmanager, cvalueid))
}

// IsPolled returns true if the device's state is being polled.
func IsPolled(homeID uint32, valueID uint64) bool {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	return bool(C.manager_isPolled(cmanager, cvalueid))
}

// SetPollIntensity sets the frequency of polling.
//
//  - 0 = none
//  - 1 = every time through the list
//  - 2 = every other time
//  - etc.
func SetPollIntensity(homeID uint32, valueID uint64, intensity uint8) {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	C.manager_setPollIntensity(cmanager, cvalueid, C.uint8_t(intensity))
}

// GetPollIntensity returns the polling intensity of a device's state.
func GetPollIntensity(homeID uint32, valueID uint64) uint8 {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	return uint8(C.manager_getPollIntensity(cmanager, cvalueid))
}
