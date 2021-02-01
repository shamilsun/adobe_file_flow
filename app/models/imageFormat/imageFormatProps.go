package imageFormat

// Alias hide the real type of the enum
// and users can use it to define the var for accepting enum
type AImageFormat = string

const (
	JPG  AImageFormat = "jpg"
	JPG1 AImageFormat = "jpg1"
)

type lImageFormat struct {
	JPG AImageFormat
	PNG AImageFormat
	PSD AImageFormat
}

// Enum for public use
var EImageFormat = &lImageFormat{
	JPG: "JPG",
	PNG: "PNG",
	PSD: "PSD",
}
