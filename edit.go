package diff

import (
	"bytes"
	"fmt"
)

// Edit represents an edit of string.
type Edit struct {
	Action int
	Value  string
	From   string
}

// EditScript is slice of Edit.
type EditScript = []Edit

// String shows Value with prefix if added or deleted.
func (e *Edit) String() string {
	switch e.Action {
	case 2:
		return fmt.Sprintf("-%s\n+%s", e.From, e.Value)
	case 1:
		return fmt.Sprintf("+%s", e.Value)
	case -1:
		return fmt.Sprintf("-%s", e.Value)
	default:
		return e.Value
	}
}

// String shows each Edit.String line by line
func String(es EditScript) string {

	b := bytes.Buffer{}

	for i, e := range es {
		b.WriteString(e.String())
		if i+1 < len(es) {
			b.WriteString("\n")
		}
	}
	return b.String()
}
