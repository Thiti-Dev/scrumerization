package rooms

import (
	"github.com/Thiti-Dev/scrumerization-core-service/graph/model"
	"github.com/google/uuid"
)

type ConnectedClient struct {
	Channel    chan *model.RoomState
	Name       string
	IsVoted    bool
	GivenPoint int8
	GivenDesc  *string
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
		// clients := []uuid.UUID{}

		// for userUUID := range rs.Clients {
		// 	clients = append(clients, userUUID)
		// }

		clientStates := []*model.ClientState{}

		for cuuid, client := range rs.Clients {
			clientStates = append(clientStates, &model.ClientState{
				UUID:    cuuid,
				Name:    client.Name,
				IsVoted: client.IsVoted,
			})
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
			Clients:      clientStates,
			OnGoingTopic: onGoingTopicData,
		}

		// fmt.Printf("broadcasted to: %v\n", uid)
	}
}

func (rs *RoomState) DisconnectClient(userID uuid.UUID) {
	delete(rs.Clients, userID) // remove user out
	go rs.BroadcastCurrentState()
}

func (rs *RoomState) InitializeClient(userID uuid.UUID, name string, ch chan *model.RoomState) bool {
	rs.Clients[userID] = &ConnectedClient{
		Channel: ch,
		Name:    name,
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

func (rs *RoomState) SetCurrentTopicToNull() {
	rs.CurrentTopic = nil
	go rs.BroadcastCurrentState()
}

func (rs *RoomState) EmitIsVote(userID uuid.UUID, point int8, description *string) {
	// ignore isVoted for now, we only care for this call
	client := rs.getClientFromUUID(userID)
	if client != nil {
		client.IsVoted = true
		client.GivenPoint = point
		client.GivenDesc = description
	}
	go rs.BroadcastCurrentState()
}

func (rs *RoomState) getClientFromUUID(uuid uuid.UUID) *ConnectedClient {
	client, ok := rs.Clients[uuid]
	if !ok {
		return nil
	}

	return client
}
