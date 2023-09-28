package notifications

import (
	"github.com/IOTech17/neo-rport/share/refs"
)

type NotificationID refs.Identifiable

type NotificationSummary struct {
	State          ProcessingState `db:"state"`
	NotificationID string          `db:"notification_id"`
	Transport      string          `db:"transport"`
	Timestamp      string          `db:"timestamp"`
	Out            string          `db:"out"`
	Err            string          `db:"err"`
}
