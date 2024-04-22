package project

import (
	"context"
	"project-testing/internal/entity/project"
	"project-testing/pkg/errors"
)

func (d Data) GetAllUser(ctx context.Context) ([]project.User, error) {
	var (
		results []project.User
		result  project.User
		err     error
	)

	rows, err := d.stmt[getAllUser].QueryxContext(ctx)
	if err != nil {
		return results, errors.Wrap(err, "[DATA][GetAllUser]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.StructScan(&result); err != nil {
			return results, errors.Wrap(err, "[DATA][GetAllUser]")
		}
		results = append(results, result)
	}

	return results, err

}

func (d Data) GetKeluargaByID(ctx context.Context, id int) (map[int][]project.Keluarga, error) {
	var (
		results = make(map[int][]project.Keluarga)
		result  project.Keluarga
		err     error
	)

	rows, err := d.stmt[getKeluargaByID].QueryxContext(ctx, id)
	if err != nil {
		return results, errors.Wrap(err, "[DATA][GetKeluargaByID]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.StructScan(&result); err != nil {
			return results, errors.Wrap(err, "[DATA][GetKeluargaByID]")
		}
		results[result.UserID] = append(results[result.UserID], result)
	}

	return results, err

}
