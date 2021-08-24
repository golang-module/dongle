package dongle

import (
	"fmt"
)

// invalidFileError returns an invalid file error
var invalidFileError = func(file string) error {
	return fmt.Errorf("invalid file %q, please make sure the file exists", file)
}

// invalidEncodingModeError returns an invalid decoding mode error
var invalidDecodingModeError = func(mode string) error {
	return fmt.Errorf("invalid decoding mode %q", mode)
}

// invalidGroupModeError returns an invalid group mode error
var invalidGroupModeError = func(mode string) error {
	return fmt.Errorf("invalid group mode %q", mode)
}

// invalidPaddingModeError returns an invalid padding mode error
var invalidPaddingModeError = func(mode string) error {
	return fmt.Errorf("invalid padding mode %q", mode)
}
