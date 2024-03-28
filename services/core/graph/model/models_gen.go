// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type CreateTopicInput struct {
	RoomID uuid.UUID `json:"roomID"`
	Name   string    `json:"name"`
}

type CreateTopicVoteInput struct {
	TopicID   uuid.UUID `json:"topicID"`
	Voted     int       `json:"voted"`
	VotedDesc *string   `json:"votedDesc,omitempty"`
}

type CreateUserInput struct {
	Username string `json:"username" jsonschema:"minLength=6,maxLength=32,required"`
	Password string `json:"password" jsonschema:"minLength=8,maxLength=64,required"`
	Name     string `json:"name" jsonschema:"minLength=3,required"`
}

type LoginUserInput struct {
	Username string `json:"username" jsonschema:"required"`
	Password string `json:"password" jsonschema:"required"`
}

type LoginUserResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

type Mutation struct {
}

type OnGoingTopic struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type PaginatedRoomResult struct {
	Data       []*Room `json:"data"`
	TotalCount int     `json:"totalCount"`
	Count      int     `json:"count"`
}

type PaginationInput struct {
	Limit  *int `json:"limit,omitempty"`
	Offset *int `json:"offset,omitempty"`
}

type PrivateVersionResponse struct {
	Version string `json:"version"`
}

type Query struct {
}

type Room struct {
	ID        uuid.UUID `json:"id"`
	CreatorID uuid.UUID `json:"creator_id"`
	RoomName  *string   `json:"room_name,omitempty"`
	Password  *string   `json:"password,omitempty"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Creator   *User     `json:"creator,omitempty"`
}

type RoomCreationInput struct {
	Name     *string `json:"name,omitempty"`
	Password *string `json:"password,omitempty"`
}

type RoomState struct {
	Clients      []uuid.UUID   `json:"clients"`
	Active       bool          `json:"active"`
	OnGoingTopic *OnGoingTopic `json:"onGoingTopic,omitempty"`
}

type RoomWhereClause struct {
	ID        *uuid.UUID `json:"id,omitempty"`
	RoomName  *string    `json:"room_name,omitempty"`
	CreatorID *uuid.UUID `json:"creator_id,omitempty"`
}

type Subscription struct {
}

type Topic struct {
	ID        uuid.UUID `json:"id"`
	RoomID    uuid.UUID `json:"roomID"`
	Name      string    `json:"name"`
	AvgScore  *float64  `json:"avgScore,omitempty"`
	IsActive  bool      `json:"isActive"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type TopicVote struct {
	TopicID   uuid.UUID `json:"topicID"`
	UserID    uuid.UUID `json:"userID"`
	Voted     int       `json:"voted"`
	VotedDesc *string   `json:"votedDesc,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	User      *User     `json:"User,omitempty"`
}

type TopicVoteQueryWhereClause struct {
	TopicID *uuid.UUID `json:"topicID,omitempty"`
}

type User struct {
	ID        uuid.UUID `json:"id"`
	Username  *string   `json:"username,omitempty"`
	Name      *string   `json:"name,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Role string

const (
	RoleAdmin Role = "ADMIN"
	RoleUser  Role = "USER"
)

var AllRole = []Role{
	RoleAdmin,
	RoleUser,
}

func (e Role) IsValid() bool {
	switch e {
	case RoleAdmin, RoleUser:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
