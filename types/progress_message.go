package types

type ProgressStatus int

const (
	Idle = iota
	Started
	Finished
	Error
	Cancelled
	NotFound
)

type ProgressMessage struct {
	Filename     string         `json:"fileName, omitempty"`
	JobName      string         `json:"jobName, omitempty"`
	Percentage   float64        `json:"percent"`
	Status       ProgressStatus `json:"status"`
	ErrorMessage string         `json:"errorMessage, omitempty"`
}
