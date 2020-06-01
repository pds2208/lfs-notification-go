package producer

import (
	"github.com/ONSDigital/lfs-notification-go/types"
)

type Progress struct {
	Message *types.ProgressMessage
}

func NewFileProgress(fileName string) *Progress {
	return &Progress{
		Message: &types.ProgressMessage{
			Filename:     fileName,
			JobName:      "",
			Percentage:   0,
			Status:       types.Idle,
			ErrorMessage: "",
		},
	}
}

func NewJobProgress(jobName string) *Progress {
	return &Progress{
		Message: &types.ProgressMessage{
			Filename:     "",
			JobName:      jobName,
			Percentage:   0,
			Status:       types.Idle,
			ErrorMessage: "",
		},
	}
}

func (up *Progress) SetPercentage(percentage float64) {
	up.Message.Percentage = percentage
}

func (up *Progress) SetUploadStarted() {
	up.Message.Status = types.Started
}

func (up *Progress) SetUploadFinished() {
	up.Message.Status = types.Finished
	up.Message.Percentage = 100
}

func (up *Progress) SetUploadError(errorMessage string) {
	up.Message.Status = types.Error
	up.Message.ErrorMessage = errorMessage
}

func (up *Progress) SetUploadCancelled() {
	up.Message.Status = types.Cancelled
}
