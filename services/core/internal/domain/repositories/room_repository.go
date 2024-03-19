package repositories

import (
	"github.com/Thiti-Dev/scrumerization-core-service/graph/model"
	"github.com/Thiti-Dev/scrumerization-core-service/internal/domain/entities"
)

type RoomRepository interface {
	FindAll(populateUser bool, where *model.RoomWhereClause) ([]entities.PopulatedRoom, error)
}
