package project

import (
	"context"
	"log"
	"project-testing/internal/entity/project"
)

func (s Service) GetAllUser(ctx context.Context) ([]project.User, error) {
	result, err := s.data.GetAllUser(ctx)
	if err != nil {
		log.Println("ERROR GetAllUser", err.Error())
		return result, err
	}
	for k, v := range result {
		keluargaDet, err := s.data.GetKeluargaByID(ctx, v.UserID)
		if err != nil {
			log.Println("ERROR GetKeluargaByID", err.Error())
			return result, err
		}
		result[k].Keluarga = keluargaDet[v.UserID]
	}
	return result, err
}
