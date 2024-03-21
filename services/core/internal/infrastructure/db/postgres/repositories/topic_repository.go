package repositories

import (
	"database/sql"
	"fmt"
	"strings"

	jetModel "github.com/Thiti-Dev/scrumerization-core-service/.gen/scrumerization/public/model"
	"github.com/Thiti-Dev/scrumerization-core-service/.gen/scrumerization/public/table"
	"github.com/Thiti-Dev/scrumerization-core-service/graph/model"
	repository "github.com/Thiti-Dev/scrumerization-core-service/internal/domain/repositories"
	"github.com/Thiti-Dev/scrumerization-core-service/internal/infrastructure/utils"
)

type TopicRepository struct {
	SqlConnection *sql.DB
	Config        *utils.Config
}

func NewTopicRepository(dbConn *sql.DB, config *utils.Config) repository.TopicRepository {
	return &TopicRepository{
		SqlConnection: dbConn,
		Config:        config,
	}
}

func (repo *TopicRepository) CreateTopic(input *model.CreateTopicInput) (*jetModel.Topics, error) {
	stmt := table.Topics.INSERT(table.Topics.RoomID, table.Topics.Name).
		VALUES(input.RoomID, input.Name).RETURNING(table.Topics.AllColumns)

	topic := jetModel.Topics{}
	err := stmt.Query(repo.SqlConnection, &topic)

	if err != nil {
		if strings.Contains(err.Error(), "idx_unique_room_id_with_active_room") {
			err = fmt.Errorf("creating topic is available if no topic is on-going")
		}
	}

	return &topic, err
}
