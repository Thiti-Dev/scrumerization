package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/Thiti-Dev/scrumerization-core-service/graph/model"
	context_type "github.com/Thiti-Dev/scrumerization-core-service/internal/domain/context"
	"github.com/Thiti-Dev/scrumerization-core-service/internal/domain/rooms"
	"github.com/Thiti-Dev/scrumerization-core-service/pkg/tokenizer"
	"github.com/google/uuid"
)

// CreateRoom is the resolver for the createRoom field.
func (r *mutationResolver) CreateRoom(ctx context.Context, input model.RoomCreationInput) (*model.Room, error) {
	userPayload := ctx.Value(context_type.UserDataCtxKey).(*tokenizer.Payload)
	room, err := r.RoomRepository.CreateRoom(userPayload.UUID, &input)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	return &model.Room{
		ID:        room.ID,
		CreatorID: room.CreatorID,
		RoomName:  room.RoomName,
		Password:  room.Password,
		IsActive:  room.IsActive,
		CreatedAt: room.CreatedAt,
		UpdatedAt: room.UpdatedAt,
	}, nil
}

// Rooms is the resolver for the rooms field.
func (r *queryResolver) Rooms(ctx context.Context, where *model.RoomWhereClause, paginate *model.PaginationInput) (*model.PaginatedRoomResult, error) {
	populateUser := false

	_ = graphql.GetOperationContext(ctx)
	fields := graphql.CollectFieldsCtx(ctx, nil)
	for _, field := range fields {
		if field.Name == "creator" {
			populateUser = true
			break
		}
	}
	paginatedResult, err := r.RoomRepository.FindAll(populateUser, where, paginate)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	userPayload := ctx.Value(context_type.UserDataCtxKey).(*tokenizer.Payload)

	var res []*model.Room
	for _, room := range paginatedResult.Data {
		roomPassword := room.Password
		if room.Password != nil {
			if userPayload.UUID != room.CreatorID {
				censored := "**********"
				roomPassword = &censored
			}
		}
		res = append(res, &model.Room{
			ID:        room.ID,
			CreatorID: room.CreatorID,
			RoomName:  room.RoomName,
			Password:  roomPassword,
			IsActive:  room.IsActive,
			CreatedAt: room.CreatedAt,
			UpdatedAt: room.UpdatedAt,
			Creator: &model.User{
				ID:        room.Creator.ID,
				Username:  room.Creator.Username,
				Name:      room.Creator.Name,
				CreatedAt: room.Creator.CreatedAt,
				UpdatedAt: room.Creator.UpdatedAt,
			},
		})
	}

	return &model.PaginatedRoomResult{
		Data:       res,
		TotalCount: paginatedResult.TotalCount,
		Count:      paginatedResult.Count,
	}, nil
}

// ConnectToRoom is the resolver for the connectToRoom field.
func (r *subscriptionResolver) ConnectToRoom(ctx context.Context, roomID uuid.UUID) (<-chan *model.RoomState, error) {
	// TODO: Check if the room is really created in the database
	//     : BACKNOTE:
	//				 : Beforehand should call roomHub.CrateRoom(uuid.MustParse(roomID)
	//								: Right after when the user successfully create the room
	//     : Check if room requires password or not
	//     : If all rights are valid
	//     : return the created Channel
	//     : otherwise giving them no right to enter error

	// Check if the room has been created yet

	// Check if room is existed
	actualRoomResultFromJet, err := r.RoomRepository.FindRoomByID(roomID)
	if err != nil {
		return nil, err
	}
	if actualRoomResultFromJet == nil {
		return nil, fmt.Errorf("this room doesn't exist")
	} else if !actualRoomResultFromJet.IsActive {
		return nil, fmt.Errorf("this room has already been archived")
	}

	var room *rooms.RoomState
	room = r.RoomHub.GetRoomStateFromRoomID(roomID)
	if room == nil {
		// return nil, fmt.Errorf("this room doesn't exist or already ended")

		// create the room for the server (in-memory)
		createdRoom, err := r.RoomHub.CrateRoom(roomID) // create the room
		if err != nil {
			return nil, err
		}
		room = createdRoom
	}

	ch := make(chan *model.RoomState)
	mockUUID := uuid.New()
	room.InitializeClient(mockUUID, ch) // register our channel

	go func() {
		defer close(ch)
		for {
			<-ctx.Done() // Wait for context to be finished
			fmt.Printf("client: %v disconnected from room %v\n", mockUUID, roomID)
			go room.DisconnectClient(mockUUID)
			return
		}
	}()

	return ch, nil
}
