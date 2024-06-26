//go:generate go run generate.go

package graph

import (
	"database/sql"

	"github.com/Thiti-Dev/scrumerization-core-service/internal/domain/repositories"
	"github.com/Thiti-Dev/scrumerization-core-service/internal/domain/rooms"
	infraUtils "github.com/Thiti-Dev/scrumerization-core-service/internal/infrastructure/utils"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	SqlConnection   *sql.DB
	UserRepository  repositories.UserRepository
	RoomRepository  repositories.RoomRepository
	TopicRepository repositories.TopicRepository
	RoomHub         *rooms.RoomHub
	Config          *infraUtils.Config
}
