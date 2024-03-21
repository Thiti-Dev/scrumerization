package rooms

import (
	"github.com/Thiti-Dev/scrumerization-core-service/graph/model"
	"github.com/google/uuid"
)

type ConnectedClient struct {
	Channel chan *model.RoomState
}

type CurrentTopic struct {
	ID   uuid.UUID
	Name string
}

type RoomState struct {
	Active       bool
	Clients      map[uuid.UUID]*ConnectedClient
	CurrentTopic *CurrentTopic
}

func (rs *RoomState) BroadcastCurrentState() {
	for _, client := range rs.Clients {
		// initial emittion
		clients := []uuid.UUID{}

		for userUUID := range rs.Clients {
			clients = append(clients, userUUID)
		}

		var onGoingTopicData *model.OnGoingTopic
		if rs.CurrentTopic != nil {
			onGoingTopicData = &model.OnGoingTopic{
				ID:   rs.CurrentTopic.ID,
				Name: rs.CurrentTopic.Name,
			}
		}

		client.Channel <- &model.RoomState{
			Active:       rs.Active,
			Clients:      clients,
			OnGoingTopic: onGoingTopicData,
		}

		// fmt.Printf("broadcasted to: %v\n", uid)
	}
}

func (rs *RoomState) DisconnectClient(userID uuid.UUID) {
	delete(rs.Clients, userID) // remove user out
	go rs.BroadcastCurrentState()
}

func (rs *RoomState) InitializeClient(userID uuid.UUID, ch chan *model.RoomState) bool {
	rs.Clients[userID] = &ConnectedClient{
		Channel: ch,
	}

	go rs.BroadcastCurrentState() // broadcast the state to all connected client after initialization is fone
	return true
}

func (rs *RoomState) SetCurrentTopic(uuid uuid.UUID, name string) {
	rs.CurrentTopic = &CurrentTopic{
		ID:   uuid,
		Name: name,
	}
	go rs.BroadcastCurrentState()
}