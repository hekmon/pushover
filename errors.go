package pushover

import (
	"fmt"

	"github.com/gregdel/pushover"
)

type ServerError struct {
	limitReached bool
	Errors       pushover.Errors
}

func (e ServerError) Error() string {
	if e.limitReached {
		e.Errors = append(e.Errors, "messages limit reached")
	}
	return fmt.Sprintf("pushover server errors: %+v", e.Errors)
}

func (e ServerError) IsLimitReached() bool {
	return e.limitReached
}
