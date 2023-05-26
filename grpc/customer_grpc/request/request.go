package request

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/google/uuid"
)

type CustomerRequest struct {
	ID      uuid.UUID           `json:"id,omitempty"`
	Name    string              `json:"name,omitempty"`
	Phone   string              `json:"phone,omitempty"`
	Address string              `json:"address,omitempty"`
	Gender  string              `json:"gender,omitempty"`
	DoB     timestamp.Timestamp `json:"DoB,omitempty"`
}
