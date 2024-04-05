package repositories

import (
	"database/sql"
	"errors"
	"log"
	"strings"
	"time"

	jetModel "github.com/Thiti-Dev/scrumerization-core-service/.gen/scrumerization/public/model"
	"github.com/Thiti-Dev/scrumerization-core-service/.gen/scrumerization/public/table"
	"github.com/Thiti-Dev/scrumerization-core-service/graph/model"
	repository "github.com/Thiti-Dev/scrumerization-core-service/internal/domain/repositories"
	"github.com/Thiti-Dev/scrumerization-core-service/internal/infrastructure/utils"
	"github.com/Thiti-Dev/scrumerization-core-service/pkg/tokenizer"
	"github.com/alexedwards/argon2id"
	jet "github.com/go-jet/jet/v2/postgres"
)

type UserRepository struct {
	SqlConnection *sql.DB
	Config        *utils.Config
}

func NewUserRepository(dbConn *sql.DB, config *utils.Config) repository.UserRepository {
	// Initialize any dependencies and return a new UserRepository instance.
	return &UserRepository{
		SqlConnection: dbConn,
		Config:        config,
	}
}

func (repo *UserRepository) LoginUser(input model.LoginUserInput) (string, string, error) {
	stmt := jet.SELECT(table.Users.ID, table.Users.Name, table.Users.Password, table.Users.Role).FROM(table.Users).WHERE(table.Users.Username.EQ(jet.String(input.Username))).LIMIT(1)
	user := jetModel.Users{}
	err := stmt.Query(repo.SqlConnection, &user)
	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			return "", "", errors.New("invalid password")
		}
		return "", "", err
	}

	isMatch, err := argon2id.ComparePasswordAndHash(input.Password, user.Password)
	if err != nil {
		log.Fatal(err)
	}

	if !isMatch {
		return "", "", errors.New("invalid password")
	}

	payload, err := tokenizer.NewPayload(user.ID, *user.Name, user.Role, time.Hour*6)
	if err != nil {
		return "", "", err
	}
	token := tokenizer.CreateToken(payload, []byte(repo.Config.JwtSecret))

	refreshTokenPayload, err := tokenizer.NewPayload(user.ID, *user.Name, user.Role, time.Hour*12)
	if err != nil {
		return "", "", err
	}
	refreshToken := tokenizer.CreateToken(refreshTokenPayload, []byte(repo.Config.JwtSecret))

	return token, refreshToken, err
}

func (repo *UserRepository) Create(input model.CreateUserInput) (*jetModel.Users, error) {
	//encrypt the password
	hashedPassword, err := argon2id.CreateHash(input.Password, argon2id.DefaultParams)
	if err != nil {
		log.Fatal(err)
	}

	stmt := table.Users.INSERT(table.Users.Username, table.Users.Password, table.Users.Name).
		VALUES(input.Username, hashedPassword, input.Name).RETURNING(table.Users.AllColumns)

	user := jetModel.Users{}
	err = stmt.Query(repo.SqlConnection, &user)

	return &user, err
}

func (repo *UserRepository) FindAll() ([]jetModel.Users, error) {
	stmt := jet.SELECT(table.Users.AllColumns).FROM(table.Users)

	users := []jetModel.Users{}
	err := stmt.Query(repo.SqlConnection, &users)
	return users, err
}
