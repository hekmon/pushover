package pushover

import (
	"fmt"
	"time"

	"github.com/gregdel/pushover"
)

// SendEmergencyPriorityMsg sends a message with a title as emergency notification
func (c *Controller) SendEmergencyPriorityMsg(msg, title string) (err error) {
	return c.SendCustomMsg(Message{
		Message:   msg,
		Title:     title,
		Priority:  PriorityEmergency,
		Timestamp: time.Now().Unix(),
	})
}

// SendHighPriorityMsg sends a message with a title as high priority notification
func (c *Controller) SendHighPriorityMsg(msg, title string) (err error) {
	return c.SendCustomMsg(Message{
		Message:   msg,
		Title:     title,
		Priority:  PriorityHigh,
		Timestamp: time.Now().Unix(),
	})
}

// SendNormalPriorityMsg sends a message with a title as normal notification
func (c *Controller) SendNormalPriorityMsg(msg, title string) (err error) {
	return c.SendCustomMsg(Message{
		Message:   msg,
		Title:     title,
		Priority:  PriorityNormal,
		Timestamp: time.Now().Unix(),
	})
}

// SendLowPriorityMsg sends a message with a title as low priority notification
func (c *Controller) SendLowPriorityMsg(msg, title string) (err error) {
	return c.SendCustomMsg(Message{
		Message:   msg,
		Title:     title,
		Priority:  PriorityLow,
		Timestamp: time.Now().Unix(),
	})
}

// SendLowestPriorityMsg sends a message with a title as lowest priority notification
func (c *Controller) SendLowestPriorityMsg(msg, title string) (err error) {
	return c.SendCustomMsg(Message{
		Message:   msg,
		Title:     title,
		Priority:  PriorityLowest,
		Timestamp: time.Now().Unix(),
	})
}

// SendCustomMsg allow to send a custom message
func (c *Controller) SendCustomMsg(msg Message) (err error) {
	// prepare message
	raw := pushover.Message{
		Message:     msg.Message,
		Title:       msg.Title,
		Priority:    int(msg.Priority),
		URL:         msg.URL,
		URLTitle:    msg.URLTitle,
		Timestamp:   msg.Timestamp,
		Retry:       msg.Retry,
		Expire:      msg.Expire,
		CallbackURL: msg.CallbackURL,
		DeviceName:  msg.DeviceName,
		Sound:       msg.Sound,
		HTML:        msg.HTML,
	}
	if msg.Attachment != nil {
		if err = raw.AddAttachment(msg.Attachment); err != nil {
			err = fmt.Errorf("failed to add attachment: %w", err)
			return
		}
	}
	// send it
	response, err := c.app.SendMessage(&raw, c.dest)
	if err != nil {
		return fmt.Errorf("sending fail: %w", err)
	} else if response != nil && response.Status != 1 {
		return ServerError{
			Errors:       response.Errors,
			limitReached: response.Limit != nil && response.Limit.Remaining == 0,
		}
	}
	return
}
