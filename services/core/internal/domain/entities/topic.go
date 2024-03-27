package entities

import (
	jetModel "github.com/Thiti-Dev/scrumerization-core-service/.gen/scrumerization/public/model"
)

type PopulatedTopicVote struct {
	jetModel.TopicVotes

	User jetModel.Users
}
