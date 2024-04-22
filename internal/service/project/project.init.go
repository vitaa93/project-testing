package project

import (
	"context"
	"project-testing/internal/entity/project"
)

type Data interface {
	GetAllUser(ctx context.Context) ([]project.User, error)
	GetKeluargaByID(ctx context.Context, id int) (map[int][]project.Keluarga, error)
}

type Service struct {
	data Data
}

func New(data Data) Service {
	return Service{
		data: data,
	}
}
