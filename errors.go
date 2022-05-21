package dongle

import (
	"fmt"
)

// returns an invalid file error
var invalidFileError = func(file string) error {
	return fmt.Errorf("invalid file %q, please make sure the file exists", file)
}

// returns an invalid key error
var invalidKeyError = func(length int, size ...int) error {
	if len(size) == 0 {
		return fmt.Errorf("invalid key size %d, key must be 16, 24 or 32 bytes", length)
	}
	return fmt.Errorf("invalid key size %d, key must be %d bytes", length, size[0])
}

// returns a overflow key error
var overflowKeyError = func(length int) error {
	return fmt.Errorf("invalid key size %d, key at least 1 byte and at most 256 bytes", length)
}

// returns an invalid IV error
var invalidIVError = func(length int, size int) error {
	return fmt.Errorf("invalid IV size %d, IV size must be %d", length, size)
}

// returns an invalid encryption mode or padding error
var invalidModeOrPaddingError = func(mode, padding string) error {
	return fmt.Errorf("invalid encryption mode %q or padding %q", mode, padding)
}
