package dto

import "github.com/google/uuid"

type IDRequest struct {
	ID uuid.UUID `json:"id"`
}
