package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/Thiti-Dev/scrumerization-core-service/graph/model"
	"github.com/Thiti-Dev/scrumerization-core-service/pkg/validation"
	_ "github.com/alexedwards/argon2id"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	verr, valid := validation.ValidateJsonSchemaFromStruct(input)

	if !valid {
		graphql.AddError(ctx, &gqlerror.Error{Message: "validation Error", Extensions: verr.FieldErrorsIntoArbitaryMapInterface()})
		return nil, nil // no need to return error as the graphql will infer the errors added by .AddError by itself
	}

	createdUser, err := r.UserRepository.Create(input)
	if err != nil {
		if strings.Contains(err.Error(), "violates unique constraint") {
			return nil, fmt.Errorf("username is already taken")
		}
		return nil, fmt.Errorf("unknown error has occured")
	}

	user := &model.User{
		ID:        createdUser.ID,
		Username:  createdUser.Username,
		CreatedAt: createdUser.CreatedAt,
		UpdatedAt: createdUser.UpdatedAt,
	}

	return user, err
}

// LoginUser is the resolver for the loginUser field.
func (r *mutationResolver) LoginUser(ctx context.Context, input model.LoginUserInput) (*model.LoginUserResponse, error) {
	if verr, valid := validation.ValidateJsonSchemaFromStruct(input); !valid {
		graphql.AddError(ctx, &gqlerror.Error{
			Message:    "validation Error",
			Extensions: verr.FieldErrorsIntoArbitaryMapInterface(),
		})
		return nil, nil // no need to return error as the graphql will infer the errors added by .AddError by itself
	}

	token, refreshToken, err := r.UserRepository.LoginUser(input)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return &model.LoginUserResponse{
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	users, err := r.UserRepository.FindAll()
	if err != nil {
		log.Fatal(err)
	}

	var res []*model.User
	for _, user := range users {
		res = append(res, &model.User{
			ID:        user.ID,
			Username:  user.Name,
			Name:      user.Name,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}
	return res, nil
}
