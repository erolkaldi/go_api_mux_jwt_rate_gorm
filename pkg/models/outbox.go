package models

import "time"

type OutBox struct {
	ID            int         `gorm:"primary_key" json:"id"`
	RecordState   RecordState `json:"record_state"`
	Data          string      `json:"data"`
	BoxType       BoxType     `json:"box_type"`
	CreatedOn     time.Time   `json:"created_on"`
	ProcessedOn   time.Time   `json:"processed_on"`
	AttemptsCount int         `json:"attempts_count"`
	LastAttemptOn time.Time   `json:"last_attempt_on"`
	Error         string      `json:"error"`
}
type RecordState int

const (
	PendingDelivery RecordState = iota
	Processing
	Delivered
	MaxAttemptsReached
)

type BoxType int

const (
	Mail BoxType = iota
	Sms
	Notification
)
