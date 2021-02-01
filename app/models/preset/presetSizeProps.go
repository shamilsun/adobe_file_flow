package preset

// Alias hide the real type of the enum
// and users can use it to define the var for accepting enum
type APresetSize = string

type LPresetSize struct {
	Big   APresetSize
	Small APresetSize
}

// Enum for public use
var EPresetSize = &LPresetSize{
	Big:   "BIG",
	Small: "SMALL",
}
