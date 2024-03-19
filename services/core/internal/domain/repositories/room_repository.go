package repositories

import "github.com/Thiti-Dev/scrumerization-core-service/internal/domain/entities"

type RoomRepository interface {
	FindAll(populateUser bool) ([]entities.PopulatedRoom, error)
}
