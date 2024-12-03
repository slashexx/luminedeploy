package configs

import (
	"fmt"
)

func FormatError(err error) string {
	return fmt.Sprintf("Error: %v", err)
}
