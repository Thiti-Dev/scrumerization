package repositories

import (
	jetModel "github.com/Thiti-Dev/scrumerization-core-service/.gen/scrumerization/public/model"
	"github.com/Thiti-Dev/scrumerization-core-service/graph/model"
	"github.com/google/uuid"
)

type TopicRepository interface {
	CreateTopic(input *model.CreateTopicInput) (*jetModel.Topics, error)
	GetTopicsFromSpecificRoom(uuid uuid.UUID) ([]jetModel.Topics, error)
	CreateTopicVote(userID uuid.UUID, input *model.CreateTopicVoteInput) (*jetModel.TopicVotes, error)
}
