package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/Thiti-Dev/scrumerization-core-service/graph/model"
	context_type "github.com/Thiti-Dev/scrumerization-core-service/internal/domain/context"
	"github.com/Thiti-Dev/scrumerization-core-service/pkg/tokenizer"
	"github.com/google/uuid"
)

// CreateTopic is the resolver for the createTopic field.
func (r *mutationResolver) CreateTopic(ctx context.Context, input *model.CreateTopicInput) (*model.Topic, error) {
	userPayload := ctx.Value(context_type.UserDataCtxKey).(*tokenizer.Payload)
	room, err := r.RoomRepository.FindRoomByID(input.RoomID)
	if err != nil {
		return nil, err
	}

	if room == nil {
		return nil, fmt.Errorf("room with id:\"%v\" doesn't exist", input.RoomID)
	}

	if room.CreatorID != userPayload.UUID {
		return nil, fmt.Errorf("requires room owner to create topic")
	}

	topic, err := r.TopicRepository.CreateTopic(input)
	if err != nil {
		return nil, err
	}

	// Broadcasting to room [async]
	// Set the current Topic
	go func() {
		room := r.RoomHub.MustGetRoomFromRoomID(input.RoomID)
		room.SetCurrentTopic(topic.ID, topic.Name)
	}()

	return &model.Topic{
		ID:        topic.ID,
		RoomID:    topic.RoomID,
		Name:      topic.Name,
		AvgScore:  topic.AvgScore,
		IsActive:  topic.IsActive,
		CreatedAt: topic.CreatedAt,
		UpdatedAt: topic.UpdatedAt,
	}, nil
}

// CreateTopicVote is the resolver for the createTopicVote field.
func (r *mutationResolver) CreateTopicVote(ctx context.Context, input *model.CreateTopicVoteInput) (*model.TopicVote, error) {
	userPayload := ctx.Value(context_type.UserDataCtxKey).(*tokenizer.Payload)
	topicVote, err := r.TopicRepository.CreateTopicVote(userPayload.UUID, input)
	if err != nil {
		return nil, err
	}

	// go func() {
	// 	topic, err := r.TopicRepository.FindOneTopic(input.TopicID)
	// 	if err == nil {
	// 		room := r.RoomHub.MustGetRoomFromRoomID(topic.RoomID)
	// 		room.EmitIsVote(userPayload.UUID, true)
	// 	}
	// 	// room := r.MustGetRoomFromRoomID(input.)
	// }()

	return &model.TopicVote{
		TopicID:   topicVote.TopicID,
		UserID:    topicVote.UserID,
		Voted:     int(topicVote.Voted),
		VotedDesc: topicVote.VotedDescription,
		CreatedAt: topicVote.CreatedAt,
		UpdatedAt: topicVote.UpdatedAt,
	}, nil
}

// TerminateTopic is the resolver for the terminateTopic field.
func (r *mutationResolver) TerminateTopic(ctx context.Context, topicID uuid.UUID) (bool, error) {
	topic, err := r.TopicRepository.FindOneTopic(topicID) // check if the topic is indeed existed
	if err != nil {
		return false, err
	}

	_, err = r.TopicRepository.TerminateTopic(topicID)
	if err != nil {
		return false, err
	}

	room := r.RoomHub.MustGetRoomFromRoomID(topic.RoomID)

	var wg sync.WaitGroup

	for uuid, client := range room.Clients {
		fmt.Println(uuid)
		wg.Add(1)
		userID := uuid
		givenPoint := int(client.GivenPoint)
		givenDesc := client.GivenDesc
		go func() {
			defer wg.Done()
			r.TopicRepository.CreateTopicVote(userID, &model.CreateTopicVoteInput{
				TopicID:   topicID,
				Voted:     givenPoint,
				VotedDesc: givenDesc,
			})
		}()
	}

	wg.Wait() // wait all pending votes to be created on actual db

	room.SetCurrentTopicToNull()

	return true, nil // return the result
}

// Vote is the resolver for the vote field.
func (r *mutationResolver) Vote(ctx context.Context, topicID uuid.UUID, input model.VoteInput) (bool, error) {
	// This method will not do any repository action
	// Only emit the state to the server

	userPayload := ctx.Value(context_type.UserDataCtxKey).(*tokenizer.Payload)

	topic, err := r.TopicRepository.FindOneTopic(topicID) // check if the topic is indeed existed
	if err != nil {
		return false, err
	}

	room := r.RoomHub.MustGetRoomFromRoomID(topic.RoomID)
	room.EmitIsVote(userPayload.UUID, int8(input.Point), input.Description)

	/* ---------------------------- TOPIC TERMINATION --------------------------- */

	shouldTerminate := true
	for _, client := range room.Clients {
		if !client.IsVoted {
			shouldTerminate = false
			break
		}
	}

	if shouldTerminate {
		fmt.Printf("Topic:%v should be terminated\n", topicID)
		go func() {
			r.TopicRepository.TerminateTopic(topicID) // later on should push this task to message-broker, to be consistent

			var wg sync.WaitGroup

			for uuid, client := range room.Clients {
				fmt.Println(uuid)
				wg.Add(1)
				userID := uuid
				givenPoint := int(client.GivenPoint)
				givenDesc := client.GivenDesc
				go func() {
					defer wg.Done()
					r.TopicRepository.CreateTopicVote(userID, &model.CreateTopicVoteInput{
						TopicID:   topicID,
						Voted:     givenPoint,
						VotedDesc: givenDesc,
					})
				}()
			}

			wg.Wait() // wait all pending votes to be created on actual db

			room.SetCurrentTopicToNull()

		}()
	}

	/* -------------------------------------------------------------------------- */

	return true, nil
}

// Topic is the resolver for the topic field.
func (r *queryResolver) Topic(ctx context.Context, topicID uuid.UUID, password *string) (*model.Topic, error) {
	topic, err := r.TopicRepository.FindOneTopic(topicID)
	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			return nil, nil
		}
		return nil, err
	}

	return &model.Topic{
		ID:        topic.ID,
		RoomID:    topic.RoomID,
		Name:      topic.Name,
		AvgScore:  topic.AvgScore,
		IsActive:  topic.IsActive,
		CreatedAt: topic.CreatedAt,
		UpdatedAt: topic.UpdatedAt,
	}, nil
}

// Topics is the resolver for the topics field.
func (r *queryResolver) Topics(ctx context.Context, roomID uuid.UUID, password *string) ([]*model.Topic, error) {
	room, err := r.RoomRepository.FindRoomByID(roomID)
	if err != nil {
		return nil, err
	}
	if room == nil {
		return nil, fmt.Errorf("the room with id:\"%v\" doesn't exist", roomID)
	}
	if room.Password != nil {
		// If password is required
		if room.Password != password {
			return nil, fmt.Errorf("invalid room password given")
		}
	}

	topics, err := r.TopicRepository.GetTopicsFromSpecificRoom(roomID)
	if err != nil {
		return nil, err
	}
	var res []*model.Topic
	for _, topic := range topics {
		res = append(res, &model.Topic{
			ID:        topic.ID,
			RoomID:    topic.RoomID,
			Name:      topic.Name,
			AvgScore:  topic.AvgScore,
			IsActive:  topic.IsActive,
			CreatedAt: topic.CreatedAt,
			UpdatedAt: topic.UpdatedAt,
		})
	}

	return res, nil
}

// TopicVotes is the resolver for the topicVotes field.
func (r *queryResolver) TopicVotes(ctx context.Context, where *model.TopicVoteQueryWhereClause) ([]*model.TopicVote, error) {
	// userPayload := ctx.Value(context_type.UserDataCtxKey).(*tokenizer.Payload)
	topics, err := r.TopicRepository.FindAll(where)
	if err != nil {
		return nil, err
	}

	var res []*model.TopicVote
	for _, topicVote := range topics {
		res = append(res, &model.TopicVote{
			TopicID:   topicVote.TopicID,
			UserID:    topicVote.UserID,
			Voted:     int(topicVote.Voted),
			VotedDesc: topicVote.VotedDescription,
			CreatedAt: topicVote.CreatedAt,
			UpdatedAt: topicVote.UpdatedAt,
			User: &model.User{
				ID:        topicVote.User.ID,
				Username:  topicVote.User.Username,
				Name:      topicVote.User.Name,
				CreatedAt: topicVote.User.CreatedAt,
				UpdatedAt: topicVote.User.UpdatedAt,
			},
		})
	}

	return res, nil
}
