package common

import uuid "github.com/satori/go.uuid"

//DetailRequest : Message to get Doctor Detail
type DetailRequest struct {
	ID uuid.UUID `json:"id" binding:"required"`
}
