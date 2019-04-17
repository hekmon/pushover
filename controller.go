package pushover

import (
	"github.com/gregdel/pushover"
	"github.com/hekmon/hllogger"
)

// Controller will wrap access to pushover notifications
type Controller struct {
	app    *pushover.Pushover
	dest   *pushover.Recipient
	logger *hllogger.HlLogger
}

// New will return an initialized and ready to use Controller. If logger is nil
// logging will be disabled. If applicationKey or userKey is nil, all methods will
// have no effect.
func New(applicationKey, userKey *string, logger *hllogger.HlLogger) *Controller {
	var c Controller
	if applicationKey != nil {
		c.app = pushover.New(*applicationKey)
	}
	if userKey != nil {
		c.dest = pushover.NewRecipient(*userKey)
	}
	if logger != nil {
		c.logger = logger
	}
	return &c
}

func (c *Controller) initialized() bool {
	return c.app != nil && c.dest != nil
}

func (c *Controller) loggingEnabled() bool {
	return c.logger != nil
}
