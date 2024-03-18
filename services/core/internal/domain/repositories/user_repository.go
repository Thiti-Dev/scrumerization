package repositories

import (
	jetModel "github.com/Thiti-Dev/scrumerization-core-service/.gen/scrumerization/public/model"
	"github.com/Thiti-Dev/scrumerization-core-service/graph/model"
)

type UserRepository interface {
	Create(input model.CreateUserInput) (*jetModel.Users, error)
	FindAll() ([]jetModel.Users, error)
}
