//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"github.com/google/uuid"
	"time"
)

type Rooms struct {
	ID        uuid.UUID `sql:"primary_key"`
	CreatorID uuid.UUID
	RoomName  *string
	Password  *string
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}