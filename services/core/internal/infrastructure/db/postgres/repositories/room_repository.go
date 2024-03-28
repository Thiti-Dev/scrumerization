package repositories

import (
	"database/sql"
	"log"
	"strings"

	jetModel "github.com/Thiti-Dev/scrumerization-core-service/.gen/scrumerization/public/model"
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

func (repo *RoomRepository) FindAll(populateUser bool, where *model.RoomWhereClause, paginate *model.PaginationInput) (entities.PaginatedRoomResult, error) {
	var roomFromClause jet.ReadableTable
	roomFromClause = table.Rooms
	projectionList := []jet.Projection{}
	if populateUser {
		roomFromClause = roomFromClause.INNER_JOIN(table.Users, table.Users.ID.EQ(table.Rooms.CreatorID))
		projectionList = append(projectionList, table.Users.AllColumns)
	}

	var whereBuilder jet.BoolExpression
	if where != nil {
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
		if where.CreatorID != nil {
			if whereBuilder != nil {
				whereBuilder = whereBuilder.AND(table.Rooms.CreatorID.EQ(jet.UUID(where.CreatorID)))
			} else {
				whereBuilder = table.Rooms.CreatorID.EQ(jet.UUID(where.CreatorID))
			}
		}
	}

	stmt := jet.SELECT(table.Rooms.AllColumns, projectionList...).FROM(roomFromClause).WHERE(whereBuilder)

	// Count is needed before the offset & limit
	countStmt := jet.SELECT(jet.COUNT(table.Rooms.ID).AS("total_count")).FROM(roomFromClause).WHERE(whereBuilder)

	if paginate != nil {
		// If paginate is present
		if paginate.Limit != nil {
			// Limit
			stmt = stmt.LIMIT(int64(*paginate.Limit))
		}
		if paginate.Offset != nil {
			stmt = stmt.OFFSET(int64(*paginate.Offset))
		}
	}

	countChannel := make(chan entities.PaginationResult)
	searchResultChannel := make(chan []entities.PopulatedRoom)

	// RUNNING THE THE SAME TIME

	go func() {
		// anonymous struct needed, otherwise the serialization process will not complete
		var countingResult struct {
			TotalCount int
		}
		err := countStmt.Query(repo.SqlConnection, &countingResult)
		if err != nil {
			log.Fatal(err)
		}

		countChannel <- entities.PaginationResult{
			TotalCount: countingResult.TotalCount,
		}
	}()

	go func() {
		rooms := []entities.PopulatedRoom{}
		err := stmt.Query(repo.SqlConnection, &rooms)
		if err != nil {
			log.Fatal(err)
		}
		searchResultChannel <- rooms
	}()

	/* -------------------------------------------------------------------------- */

	countRes, roomRes := <-countChannel, <-searchResultChannel

	return entities.PaginatedRoomResult{
		Data: roomRes,
		PaginationResult: entities.PaginationResult{
			TotalCount: countRes.TotalCount,
			Count:      len(roomRes),
		},
	}, nil
}

func (repo *RoomRepository) CreateRoom(creatorId uuid.UUID, input *model.RoomCreationInput) (*jetModel.Rooms, error) {
	stmt := table.Rooms.INSERT(table.Rooms.RoomName, table.Rooms.Password, table.Rooms.CreatorID).
		VALUES(input.Name, input.Password, creatorId).RETURNING(table.Rooms.AllColumns)

	room := jetModel.Rooms{}
	err := stmt.Query(repo.SqlConnection, &room)

	return &room, err
}

func (repo *RoomRepository) FindRoomByID(roomID uuid.UUID) (*entities.PopulatedRoom, error) {
	stmt := table.Rooms.SELECT(table.Rooms.AllColumns).FROM(table.Rooms).WHERE(table.Rooms.ID.EQ(jet.UUID(roomID))).LIMIT(1)
	room := entities.PopulatedRoom{}
	err := stmt.Query(repo.SqlConnection, &room)
	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			// No row has been found
			return nil, nil // returns out the nil instead
		}
	}

	return &room, err
}
