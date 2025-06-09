package pushover

import (
	"github.com/gregdel/pushover"
)

// Controller will wrap access to pushover notifications
type Controller struct {
	app  *pushover.Pushover
	dest *pushover.Recipient
}

// New will return an initialized and ready to use Controller.
// If applicationKey or userKey is nil, all methods will have no effect.
func New(applicationKey, userKey string) (c *Controller) {
	return &Controller{
		app:  pushover.New(applicationKey),
		dest: pushover.NewRecipient(userKey),
	}
}
