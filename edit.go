package diff

import (
	"fmt"
	"strings"
)

// Edit represents an edit of string.
type Edit struct {
	Action int
	Value  string
}

// EditScript is slice of Edit.
type EditScript = []Edit

// String shows Value with prefix if added or deleted.
func (e *Edit) String() string {
	switch e.Action {
	case 1:
		return fmt.Sprintf("+%s", e.Value)
	case -1:
		return fmt.Sprintf("-%s", e.Value)
	default:
		return fmt.Sprintf("%s", e.Value)
	}
}

// String shows each Edit.String line by line
func String(es EditScript) string {
	var b []string

	for _, e := range es {
		b = append(b, e.String())
	}
	return strings.Join(b, "\n")
}
