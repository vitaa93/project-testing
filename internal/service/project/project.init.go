package project

import (
	"context"
	"project-testing/internal/entity/project"
)

type Data interface {
	GetAllUser(ctx context.Context) ([]project.User, error)
	GetKeluargaByID(ctx context.Context, id int) (map[int][]project.Keluarga, error)
	SearchUserByNameAndttl(ctx context.Context, nama string, ttl string) ([]project.User, error)
	GetAllUserPagination(ctx context.Context, limit int, offset int) ([]project.User, error)
	GetTotalAllUser(ctx context.Context) (int, error)
	SearchUserDataByName(ctx context.Context, keyword string) ([]project.User, error)
	SearchUserDataByKwn(ctx context.Context, keyword string) ([]project.User, error)
	SearchUserDataByKwnOrName(ctx context.Context, keyword string) ([]project.User, error)
	InsertDatauser(ctx context.Context, datas project.User) error
}

type Service struct {
	data Data
}

func New(data Data) Service {
	return Service{
		data: data,
	}
}
