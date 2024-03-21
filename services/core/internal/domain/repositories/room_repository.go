package repositories

import (
	jetModel "github.com/Thiti-Dev/scrumerization-core-service/.gen/scrumerization/public/model"
	"github.com/Thiti-Dev/scrumerization-core-service/graph/model"
	"github.com/Thiti-Dev/scrumerization-core-service/internal/domain/entities"
	"github.com/google/uuid"
)

type RoomRepository interface {
	FindAll(populateUser bool, where *model.RoomWhereClause) ([]entities.PopulatedRoom, error)
	FindRoomByID(roomID uuid.UUID) (*entities.PopulatedRoom, error)
	CreateRoom(creatorId uuid.UUID, input *model.RoomCreationInput) (*jetModel.Rooms, error)
}
