// +build sam

package machine

type I2SMode uint8
type I2SStandard uint8
type I2SClockSource uint8
type I2SDataFormat uint8

const (
	I2SModeMaster I2SMode = iota
	I2SModeSlave
)

const (
	I2StandardPhilips I2SStandard = iota
	I2SStandardMSB
	I2SStandardLSB
)

const (
	I2SClockSourceInternal I2SClockSource = iota
	I2SClockSourceExternal
)

const (
	I2SDataFormatDefault I2SDataFormat = 0
	I2SDataFormat8bit                  = 8
	I2SDataFormat16bit                 = 16
	I2SDataFormat24bit                 = 24
	I2SDataFormat32bit                 = 32
)

// All fields are optional and may not be required or used on a particular platform.
type I2SConfig struct {
	SCK               uint8
	WS                uint8
	SD                uint8
	Mode              I2SMode
	Standard          I2SStandard
	ClockSource       I2SClockSource
	DataFormat        I2SDataFormat
	AudioFrequency    uint32
	MasterClockOutput bool
}

// var I2S0 = ... // implementation defined

// Not a real exported type, just here to serve as example.
// Supports the io.ReadWriteCloser interface
// type I2SReadWriteCloser interface {
//     Configure(I2SConfig)
//     Read(p []byte) (n int, err error)
//     Write(p []byte) (n int, err error)
//	   ReadByte() (byte, error)
//     WriteByte(c byte) error
//     Close() error
// }
