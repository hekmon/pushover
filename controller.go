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
func New(applicationKey, userKey *string) (c *Controller) {
	c = new(Controller)
	if applicationKey != nil {
		c.app = pushover.New(*applicationKey)
	}
	if userKey != nil {
		c.dest = pushover.NewRecipient(*userKey)
	}
	return
}

func (c *Controller) initialized() bool {
	return c.app != nil && c.dest != nil
}
