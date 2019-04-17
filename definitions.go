package pushover

import (
	"time"

	"github.com/gregdel/pushover"
)

// Priority represents a given message priority
type Priority int

const (
	// PriorityEmergency will be shown as red notification bypassing silence settings on dest.
	PriorityEmergency Priority = pushover.PriorityEmergency
	// PriorityHigh will be shown as red notifications
	PriorityHigh Priority = pushover.PriorityHigh
	// PriorityNormal will create a notification tile and sound
	PriorityNormal Priority = pushover.PriorityNormal
	// PriorityLow will create a notification tile only (no sound or vibration)
	PriorityLow Priority = pushover.PriorityLow
	// PriorityLowest will send the message which will show up on application but won't generate any notification
	PriorityLowest Priority = pushover.PriorityLowest
)

// Message represents all the usable fields for a message
type Message struct {
	Message     string
	Title       string
	Priority    Priority
	URL         string
	URLTitle    string
	Timestamp   int64
	Retry       time.Duration
	Expire      time.Duration
	CallbackURL string
	DeviceName  string
	Sound       string
	HTML        bool
}
