package outbox

type EventStatus string

const (
	EventStatusCreated    EventStatus = "created"
	EventStatusProcessing EventStatus = "processing"
	EventStatusSent       EventStatus = "sent"
	EventStatusError      EventStatus = "error"
)

type Event struct {
	ID     uint64
	Topic  string
	Data   []byte
	Status string
}
