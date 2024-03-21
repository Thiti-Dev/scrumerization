package rooms

import (
	"github.com/Thiti-Dev/scrumerization-core-service/graph/model"
	"github.com/google/uuid"
)

type ConnectedClient struct {
	Channel chan *model.RoomState
}

type RoomState struct {
	Active  bool
	Clients map[uuid.UUID]*ConnectedClient
}

func (rs *RoomState) BroadcastCurrentState() {
	for _, client := range rs.Clients {
		// initial emittion
		clients := []uuid.UUID{}

		for userUUID := range rs.Clients {
			clients = append(clients, userUUID)
		}

		client.Channel <- &model.RoomState{
			Active:  rs.Active,
			Clients: clients,
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
