package repositories

import (
	"database/sql"
	"log"

	"github.com/Thiti-Dev/scrumerization-core-service/.gen/scrumerization/public/table"
	"github.com/Thiti-Dev/scrumerization-core-service/internal/domain/entities"
	repository "github.com/Thiti-Dev/scrumerization-core-service/internal/domain/repositories"
	"github.com/Thiti-Dev/scrumerization-core-service/internal/infrastructure/utils"
	jet "github.com/go-jet/jet/v2/postgres"
)

type RoomRepository struct {
	SqlConnection *sql.DB
	Config        *utils.Config
}

func NewRoomRepository(dbConn *sql.DB, config *utils.Config) repository.RoomRepository {
	return &RoomRepository{
		SqlConnection: dbConn,
		Config:        config,
	}
}

func (repo *RoomRepository) FindAll(populateUser bool) ([]entities.PopulatedRoom, error) {

	var roomFromClause jet.ReadableTable
	roomFromClause = table.Rooms
	projectionList := []jet.Projection{}
	if populateUser {
		roomFromClause = roomFromClause.INNER_JOIN(table.Users, table.Users.ID.EQ(table.Rooms.CreatorID))
		projectionList = append(projectionList, table.Users.AllColumns)
	}
	stmt := jet.SELECT(table.Rooms.AllColumns, projectionList...).FROM(roomFromClause)
	rooms := []entities.PopulatedRoom{}
	err := stmt.Query(repo.SqlConnection, &rooms)
	if err != nil {
		log.Fatal(err)
	}

	return rooms, nil
}
