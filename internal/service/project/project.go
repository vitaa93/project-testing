package project

import (
	"context"
	"log"
	"math"
	"project-testing/internal/entity/project"
	"project-testing/pkg/errors"
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

func (s Service) SearchUserByNameAndttl(ctx context.Context, name string, ttl string) ([]project.User, string, error) {
	var message string
	var result []project.User
	var err error
	if name == "" {
		message = "Harap mengisi nama anda"
		return result, message, err
	} else if ttl == "" {
		message = "Harap mengisi tanggal lahir anda"
		return result, message, err
	}

	result, err = s.data.SearchUserByNameAndttl(ctx, name, ttl)
	if err != nil {
		log.Println("ERROR SearchUserByNameAndttl", err.Error())
		return result, message, err
	}
	for k, v := range result {
		keluargaDet, err := s.data.GetKeluargaByID(ctx, v.UserID)
		if err != nil {
			log.Println("ERROR SearchUserByNameAndttl-GetKeluargaByID", err.Error())
			return result, message, err
		}
		result[k].Keluarga = keluargaDet[v.UserID]
	}
	return result, message, err
}

func (s Service) GetAllUserPagination(ctx context.Context, limit, offset int) ([]project.User, interface{}, error) {
	var metadata = make(map[string]interface{})

	metadata["maxpage"] = 1
	metadata["totaldata"] = 0
	result, err := s.data.GetAllUserPagination(ctx, limit, offset)
	if err != nil {
		log.Println("ERROR GetAllUserPagination-GetAllUserPagination", err.Error())
		return result, metadata, err
	}
	for k, v := range result {
		keluargaDet, err := s.data.GetKeluargaByID(ctx, v.UserID)
		if err != nil {
			log.Println("ERROR GetAllUserPagination-GetKeluargaByID", err.Error())
			return result, metadata, err
		}
		result[k].Keluarga = keluargaDet[v.UserID]
	}
	total, err := s.data.GetTotalAllUser(ctx)
	if err != nil {
		return result, metadata, err
	}
	page := math.Ceil(float64(total) / float64(limit))
	if page < 1 {
		page = 1
	}
	metadata["maxpage"] = page
	metadata["totaldata"] = total

	return result, metadata, err
}

func (s Service) GetUserByKwn(ctx context.Context, kewarganegaraan string) ([]project.User, error) {
	var result []project.User
	var err error

	result, err = s.data.GetUserByKwn(ctx, kewarganegaraan)
	if err != nil {
		log.Println("ERROR GetUserByKwn", err.Error())
		return result, err
	}
	for k, v := range result {
		keluargaDet, err := s.data.GetKeluargaByID(ctx, v.UserID)
		if err != nil {
			log.Println("ERROR GetUserByKwn-GetKeluargaByID", err.Error())
			return result, err
		}
		result[k].Keluarga = keluargaDet[v.UserID]
	}
	return result, err
}

func (s Service) SearchUserDataByName(ctx context.Context, keyword string) ([]project.User, error) {
	var result []project.User
	var err error

	result, err = s.data.SearchUserDataByName(ctx, keyword)
	if err != nil {
		log.Println("ERROR SearchUserDataByName", err.Error())
		return result, err
	}
	for k, v := range result {
		keluargaDet, err := s.data.GetKeluargaByID(ctx, v.UserID)
		if err != nil {
			log.Println("ERROR SearchUserDataByName-GetKeluargaByID", err.Error())
			return result, err
		}
		result[k].Keluarga = keluargaDet[v.UserID]
	}
	return result, err
}

func (s Service) SearchUserDataByKwn(ctx context.Context, keyword string) ([]project.User, error) {
	var result []project.User
	var err error

	result, err = s.data.SearchUserDataByKwn(ctx, keyword)
	if err != nil {
		log.Println("ERROR SearchUserDataByKwn", err.Error())
		return result, err
	}
	for k, v := range result {
		keluargaDet, err := s.data.GetKeluargaByID(ctx, v.UserID)
		if err != nil {
			log.Println("ERROR SearchUserDataByKwn-GetKeluargaByID", err.Error())
			return result, err
		}
		result[k].Keluarga = keluargaDet[v.UserID]
	}
	return result, err
}

func (s Service) SearchUserDataByKwnOrName(ctx context.Context, keyword string) ([]project.User, error) {
	var result []project.User
	var err error

	result, err = s.data.SearchUserDataByKwnOrName(ctx, keyword)
	if err != nil {
		log.Println("ERROR SearchUserDataByKwnOrName", err.Error())
		return result, err
	}
	for k, v := range result {
		keluargaDet, err := s.data.GetKeluargaByID(ctx, v.UserID)
		if err != nil {
			log.Println("ERROR SearchUserDataByKwnOrName-GetKeluargaByID", err.Error())
			return result, err
		}
		result[k].Keluarga = keluargaDet[v.UserID]
	}
	return result, err
}

func (s Service) InsertDatauser(ctx context.Context, datas project.User) error {
	err := s.data.InsertDatauser(ctx, datas)
	if err != nil {
		log.Println("ERROR InsertDataUser", err.Error())
		return errors.Wrap(err, "ERROR InsertDataUser")
	}
	return err
}
