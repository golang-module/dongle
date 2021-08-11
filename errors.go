package dongle

import (
	"fmt"
)

// invalidFileError returns an invalid file error
var invalidFileError = func(file string) error {
	return fmt.Errorf("invalid file %q, please make sure the file exists", file)
}
