package dongle

import (
	"fmt"
)

// returns an invalid file error
var invalidFileError = func(file string) error {
	return fmt.Errorf("invalid file %q, please make sure the file exists", file)
}

// returns a decode src error
var decodeSrcError = func(mode string) error {
	return fmt.Errorf("invalid src, the src can't be decoded by %s", mode)
}

// returns an invalid src error
var invalidSrcError = func(size int) error {
	return fmt.Errorf("invalid src size %d, the src must be multiple of 16", size)
}

// returns a overflow key error
var overflowKeyError = func(size int) error {
	return fmt.Errorf("invalid key size %d, the key at least 1 byte and at most 256 bytes", size)
}

// returns an invalid IV error
var invalidIVError = func(length int, size int) error {
	return fmt.Errorf("invalid iv size %d, the iv size must be %d", length, size)
}

// returns an invalid encryption mode or padding error
var invalidModeOrPaddingError = func(mode, padding string) error {
	return fmt.Errorf("invalid encryption mode %q or padding %q", mode, padding)
}
