package request

import "regexp"

// Constants for request validation
var (
	emailRegex        = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	maxNameLength     = 50
	maxEmailLength    = 100
	maxPositionLength = 50
)
