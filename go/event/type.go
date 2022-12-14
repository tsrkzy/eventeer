package event

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/tsrkzy/jump_in/models"
)

type ListRequest struct{}

type Event struct {
	models.Event
}
type ListResponse struct {
	EventsOwns    []Event `json:"events_owns"`
	EventsJoins   []Event `json:"events_joins"`
	EventsRunning []Event `json:"events_running"`
}

type CreateRequest struct {
	Name string
}

func (r CreateRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(
			&r.Name,
			validation.Required.Error("イベント名は必須です"),
			validation.RuneLength(5, 40).Error("イベント名は5〜40文字で指定してください"),
		),
	)
}

type CreateResponse struct {
}

type DetailRequest struct {
	EventId int `query:"event_id"`
}

type DetailResponse struct {
	Event
}
