// Example using the i2s hardware interface on the Adafruit Circuit Playground Express
// to read data from the onboard MEMS microphone.
//
// Uses ideas from the https://github.com/adafruit/Adafruit_CircuitPlayground repo.
//
package main

import (
	"machine"
)

const (
	SAMPLERATE_HZ = 22000
	DECIMATION    = 64
)

var (
	// the current sample data buffer
	data = make([]uint32, 64)

	// the sum for the current set of samples
	sum uint16
)

func main() {
	machine.I2S0.Configure(machine.I2SConfig{})

	for {
		// get the next group of samples
		machine.I2S0.Read(data)

		// process the samples
		sum = 0
		for i := 0; i < len(data); i += (DECIMATION / 16) {
			for j := 0; j < (DECIMATION / 16); j++ {
				// takes only the low order 16-bits
				sum += applySincFilter(uint16(data[i+j] & 0xffff))
			}
		}

		// adjust to 10 bit value
		s := int32(sum >> 6)

		// make it close to 0-offset signed
		s -= 512

		println("s", s)
	}
}

// a windowed sinc filter for 44 khz with 64 samples
var sincfilter = [DECIMATION]uint16{0, 2, 9, 21, 39, 63, 94, 132, 179, 236, 302, 379, 467, 565, 674, 792, 920, 1055, 1196, 1341, 1487, 1633, 1776, 1913, 2042, 2159, 2263, 2352, 2422,
	2474, 2506, 2516, 2506, 2474, 2422, 2352, 2263, 2159, 2042, 1913, 1776, 1633, 1487, 1341, 1196, 1055, 920, 792, 674, 565, 467, 379, 302, 236, 179, 132, 94, 63, 39, 21, 9, 2, 0, 0}

// applySincFilter uses a sinc filter to process a single sample value.
//
// For more information: https://en.wikipedia.org/wiki/Sinc_filter
func applySincFilter(sample uint16) (result uint16) {
	for i := 0; i < 16; i++ {
		if (sample & 0x01) > 0 {
			result += sincfilter[i]
		}
		sample >>= 1
	}
	return
}
