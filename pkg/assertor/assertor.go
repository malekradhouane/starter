package assertor

import (
	"fmt"
	"strings"
)

// Assertor exposes Separator used between message.
type Assertor struct {
	Separator string
	msgs      []string
}

// New returns a new assertor.
// with char ';' as the default separator between error messages.
func New() *Assertor {
	return &Assertor{
		msgs:      []string{},
		Separator: "; ",
	}
}

// Assert returns true when assertion is logically true else false.
func (a *Assertor) Assert(ok bool, format string, args ...any) bool {
	if !ok && format != "" {
		a.msgs = append(a.msgs, fmt.Errorf(format, args...).Error()) //nolint
	}

	return ok
}

// Validate returns an error if at least one Assert() has failed.
// Error message contains the list of unsatisfied assertions.
func (a *Assertor) Validate() error {
	if len(a.msgs) == 0 {
		return nil // No errors
	}

	return fmt.Errorf("%w: %d unsatisfied requirement(s): %s",
		ErrValidate,
		len(a.msgs),
		strings.Join(a.msgs, a.Separator))
}
