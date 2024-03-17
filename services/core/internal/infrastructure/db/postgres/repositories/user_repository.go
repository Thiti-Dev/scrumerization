package repositories

import (
	"database/sql"

	jetModel "github.com/Thiti-Dev/scrumerization-core-service/.gen/scrumerization/public/model"
	"github.com/Thiti-Dev/scrumerization-core-service/.gen/scrumerization/public/table"
	"github.com/Thiti-Dev/scrumerization-core-service/graph/model"
	repository "github.com/Thiti-Dev/scrumerization-core-service/internal/domain/repositories"
	jet "github.com/go-jet/jet/v2/postgres"
)

type UserRepository struct {
	SqlConnection *sql.DB
}

func NewUserRepository(dbConn *sql.DB) repository.UserRepository {
	// Initialize any dependencies and return a new UserRepository instance.
	return &UserRepository{
		SqlConnection: dbConn,
	}
}

func (repo *UserRepository) Create(input model.NewUser) (*jetModel.Users, error) {
	stmt := table.Users.INSERT(table.Users.Username, table.Users.Password, table.Users.Name).
		VALUES(input.Username, input.Username, input.Name).RETURNING(table.Users.AllColumns)

	user := jetModel.Users{}
	err := stmt.Query(repo.SqlConnection, &user)

	return &user, err
}

func (repo *UserRepository) FindAll() ([]jetModel.Users, error) {
	stmt := jet.SELECT(table.Users.AllColumns).FROM(table.Users)

	users := []jetModel.Users{}
	err := stmt.Query(repo.SqlConnection, &users)
	return users, err
}
