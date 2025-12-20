package gordle

import "strings"

type hint byte

const (
	absentCharacter hint = iota
	wrongPosition
	correctPosition
)

// String implements the Strnger interface.
func (h hint) String() string {
	switch h {
	case absentCharacter:
		return "🩶"
	case wrongPosition:
		return "🟡"
	case correctPosition:
		return "💚"
	default:
		return "💔"
	}
}

type feedback []hint

// StringConcat is a naive implementation to build feedback as a string.
// It is used only to benchmark it against the strings.Builder version.
func (fb feedback) StringConcat() string {
	var output string
	for _, h := range fb {
		output += h.String()
	}
	return output
}

// String implements the Stringer interface for a slice of hints.
func (fb feedback) String() string {
	sb := strings.Builder{}
	for _, h := range fb {
		sb.WriteString(h.String())
	}
	return sb.String()
}

// Equal determines equality of two feedbacks.
func (fb feedback) Equal(other feedback) bool {
	if len(fb) != len(other) {
		return false
	}
	for index, value := range fb {
		if value != other[index] {
			return false
		}
	}
	return true
}