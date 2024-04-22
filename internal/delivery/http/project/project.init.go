package project

import (
	"context"
	"project-testing/internal/entity/project"
)

type ProjectSvc interface {
	GetAllUser(ctx context.Context) ([]project.User, error)
}

type Handler struct {
	projectSvc ProjectSvc
}

func New(is ProjectSvc) *Handler {
	return &Handler{
		projectSvc: is,
	}
}
