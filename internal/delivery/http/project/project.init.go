package project

import (
	"context"
	"project-testing/internal/entity/project"
)

type ProjectSvc interface {
	GetAllUser(ctx context.Context) ([]project.User, error)
	SearchUserByNameAndttl(ctx context.Context, name string, ttl string) ([]project.User, string, error)
	GetAllUserPagination(ctx context.Context, limit, offset int) ([]project.User, interface{}, error)
	GetUserByKwn(ctx context.Context, kewarganegaraan string) ([]project.User, error)
	SearchUserDataByName(ctx context.Context, keyword string) ([]project.User, error)
	SearchUserDataByKwn(ctx context.Context, keyword string) ([]project.User, error)
	SearchUserDataByKwnOrName(ctx context.Context, keyword string) ([]project.User, error)
	InsertDatauser(ctx context.Context, datas project.User) error
}

type Handler struct {
	projectSvc ProjectSvc
}

func New(is ProjectSvc) *Handler {
	return &Handler{
		projectSvc: is,
	}
}
