package rooms

import (
	"fmt"

	"github.com/Thiti-Dev/scrumerization-core-service/graph/model"
	"github.com/google/uuid"
)

type RoomHub struct {
	roomStateByRoomId map[uuid.UUID]*RoomState
}

func NewRoomHub() *RoomHub {
	return &RoomHub{
		roomStateByRoomId: make(map[uuid.UUID]*RoomState),
	}
}

func (hub *RoomHub) CrateRoom(roomID uuid.UUID) (*RoomState, error) {
	if hub.DoesRoomExist(roomID) {
		return nil, fmt.Errorf("room does already exists")
	}

	// ch := make(chan *RoomState)
	roomState := &RoomState{
		Active:         true,
		Clients:        make(map[uuid.UUID]*ConnectedClient),
		TopicListeners: []chan *model.Topic{},
		RoomID:         roomID,
	}
	hub.roomStateByRoomId[roomID] = roomState

	return roomState, nil
}

func (hub *RoomHub) DoesRoomExist(roomID uuid.UUID) bool {
	_, ok := hub.roomStateByRoomId[roomID]
	return ok
}

func (hub *RoomHub) GetRoomStateFromRoomID(roomID uuid.UUID) *RoomState {
	if hub.DoesRoomExist(roomID) {
		return hub.roomStateByRoomId[roomID]
	}
	return nil
}

func (hub *RoomHub) MustGetRoomFromRoomID(roomID uuid.UUID) *RoomState {
	if hub.DoesRoomExist(roomID) {
		return hub.roomStateByRoomId[roomID]
	}
	// create one
	createdRoom, _ := hub.CrateRoom(roomID) // No need to check for error as this would always be successful
	return createdRoom
}

// func (hub *RoomHub) AddClientChannelToRoom(roomID uuid.UUID, userID uuid.UUID, ch chan *RoomState) bool {
// 	room := hub.GetRoomStateFromRoomID(roomID)
// 	if room == nil {
// 		return false
// 	}

// 	return room.InitializeClient(userID, ch)
// }

// func (hub *RoomHub) GetGoChannelFromRoom(roomID uuid.UUID) chan *RoomState {
// 	room := hub.GetRoomStateFromRoomID(roomID)
// 	if room != nil {
// 		return room.Channel
// 	}
// 	return nil
// }
