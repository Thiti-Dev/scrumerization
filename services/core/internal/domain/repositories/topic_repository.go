package repositories

import (
	jetModel "github.com/Thiti-Dev/scrumerization-core-service/.gen/scrumerization/public/model"
	"github.com/Thiti-Dev/scrumerization-core-service/graph/model"
)

type TopicRepository interface {
	CreateTopic(input *model.CreateTopicInput) (*jetModel.Topics, error)
}
