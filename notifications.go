package pushover

import (
	"strings"
	"time"

	"github.com/gregdel/pushover"
)

// SendEmergencyPriorityMsg sends a message with a title as emergency notification
func (c *Controller) SendEmergencyPriorityMsg(msg, title, logprefix string) {
	c.SendCustomMsg(Message{
		Message:   msg,
		Title:     title,
		Priority:  PriorityEmergency,
		Timestamp: time.Now().Unix(),
	}, logprefix)
}

// SendHighPriorityMsg sends a message with a title as high priority notification
func (c *Controller) SendHighPriorityMsg(msg, title, logprefix string) {
	c.SendCustomMsg(Message{
		Message:   msg,
		Title:     title,
		Priority:  PriorityHigh,
		Timestamp: time.Now().Unix(),
	}, logprefix)
}

// SendNormalPriorityMsg sends a message with a title as normal notification
func (c *Controller) SendNormalPriorityMsg(msg, title, logprefix string) {
	c.SendCustomMsg(Message{
		Message:   msg,
		Title:     title,
		Priority:  PriorityNormal,
		Timestamp: time.Now().Unix(),
	}, logprefix)
}

// SendLowPriorityMsg sends a message with a title as low priority notification
func (c *Controller) SendLowPriorityMsg(msg, title, logprefix string) {
	c.SendCustomMsg(Message{
		Message:   msg,
		Title:     title,
		Priority:  PriorityLow,
		Timestamp: time.Now().Unix(),
	}, logprefix)
}

// SendLowestPriorityMsg sends a message with a title as lowest priority notification
func (c *Controller) SendLowestPriorityMsg(msg, title, logprefix string) {
	c.SendCustomMsg(Message{
		Message:   msg,
		Title:     title,
		Priority:  PriorityLowest,
		Timestamp: time.Now().Unix(),
	}, logprefix)
}

// SendCustomMsg allow to send a custom message
func (c *Controller) SendCustomMsg(msg Message, logprefix string) {
	if !c.initialized() {
		if c.loggingEnabled() {
			c.logger.Debugf("[Pushover] %s: pushover unitialized: msg won't be sent", logprefix)
		}
		return
	}
	response, err := c.app.SendMessage(&pushover.Message{
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
	}, c.dest)
	if err != nil {
		if c.loggingEnabled() {
			c.logger.Errorf("[Pushover] %s: can't send msg: %v", logprefix, err)
		}
		return
	}
	if len(response.Errors) > 0 {
		if c.loggingEnabled() {
			c.logger.Errorf("[Pushover] %s: msg sent but server returned %d error()s: %+v",
				logprefix, len(response.Errors), response.Errors)
		}
		return
	}
	if c.loggingEnabled() {
		c.logger.Debugf("[Pushover] %s: message sent successfully: %s",
			logprefix, strings.Replace(response.String(), "\n", " | ", -1))
	}
}
