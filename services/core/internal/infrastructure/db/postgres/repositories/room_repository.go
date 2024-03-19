package repositories

import (
	"database/sql"
	"log"

	"github.com/Thiti-Dev/scrumerization-core-service/.gen/scrumerization/public/table"
	"github.com/Thiti-Dev/scrumerization-core-service/graph/model"
	"github.com/Thiti-Dev/scrumerization-core-service/internal/domain/entities"
	repository "github.com/Thiti-Dev/scrumerization-core-service/internal/domain/repositories"
	"github.com/Thiti-Dev/scrumerization-core-service/internal/infrastructure/utils"
	jet "github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
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

func (repo *RoomRepository) FindAll(populateUser bool, where *model.RoomWhereClause) ([]entities.PopulatedRoom, error) {

	var roomFromClause jet.ReadableTable
	roomFromClause = table.Rooms
	projectionList := []jet.Projection{}
	if populateUser {
		roomFromClause = roomFromClause.INNER_JOIN(table.Users, table.Users.ID.EQ(table.Rooms.CreatorID))
		projectionList = append(projectionList, table.Users.AllColumns)
	}

	var whereBuilder jet.BoolExpression
	if where.ID != nil {
		whereBuilder = table.Rooms.ID.EQ(jet.UUID(uuid.MustParse(where.ID.String())))
	}
	if where.RoomName != nil {
		if whereBuilder != nil {
			whereBuilder = whereBuilder.AND(table.Rooms.RoomName.LIKE(jet.String(*where.RoomName)))
		} else {
			whereBuilder = table.Rooms.RoomName.LIKE(jet.String(*where.RoomName))
		}
	}

	stmt := jet.SELECT(table.Rooms.AllColumns, projectionList...).FROM(roomFromClause).WHERE(whereBuilder)
	rooms := []entities.PopulatedRoom{}
	err := stmt.Query(repo.SqlConnection, &rooms)
	if err != nil {
		log.Fatal(err)
	}

	return rooms, nil
}
