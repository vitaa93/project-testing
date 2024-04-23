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

func (d Data) SearchUserByNameAndttl(ctx context.Context, nama string, ttl string) ([]project.User, error) {
	var (
		results []project.User
		result  project.User
		err     error
	)

	rows, err := d.stmt[searchUserByNameAndttl].QueryxContext(ctx, nama, ttl)
	if err != nil {
		return results, errors.Wrap(err, "[DATA][SearchUserByNameAndttl]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.StructScan(&result); err != nil {
			return results, errors.Wrap(err, "[DATA][SearchUserByNameAndttl]")
		}
		results = append(results, result)
	}

	return results, err

}

func (d Data) GetAllUserPagination(ctx context.Context, limit int, offset int) ([]project.User, error) {
	var (
		results []project.User
		result  project.User
		err     error
	)

	rows, err := d.stmt[getAllUserPagination].QueryxContext(ctx, limit, offset)
	if err != nil {
		return results, errors.Wrap(err, "[DATA][GetAllUserPagination]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.StructScan(&result); err != nil {
			return results, errors.Wrap(err, "[DATA][GetAllUserPagination]")
		}
		results = append(results, result)
	}

	return results, err
}

func (d Data) GetTotalAllUser(ctx context.Context) (int, error) {
	var (
		result int
		err    error
	)
	if err := d.stmt[getTotalAllUser].QueryRowxContext(ctx).Scan(&result); err != nil {
		return result, errors.Wrap(err, "[DATA][GetTotalAllUser]")
	}
	return result, err
}

func (d Data) GetUserByKwn(ctx context.Context, kewarganegaraan string) ([]project.User, error) {
	var (
		results []project.User
		result  project.User
		err     error
	)

	rows, err := d.stmt[getUserByKwn].QueryxContext(ctx, kewarganegaraan)
	if err != nil {
		return results, errors.Wrap(err, "[DATA][GetUserByKwn]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.StructScan(&result); err != nil {
			return results, errors.Wrap(err, "[DATA][GetUserByKwn]")
		}
		results = append(results, result)
	}

	return results, err
}

func (d Data) SearchUserDataByName(ctx context.Context, keyword string) ([]project.User, error) {
	var (
		results []project.User
		result  project.User
		err     error
	)

	rows, err := d.stmt[searchUserDataByName].QueryxContext(ctx, keyword)
	if err != nil {
		return results, errors.Wrap(err, "[DATA][SearchUserDataByName]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.StructScan(&result); err != nil {
			return results, errors.Wrap(err, "[DATA][SearchUserDataByName]")
		}
		results = append(results, result)
	}

	return results, err
}

func (d Data) SearchUserDataByKwn(ctx context.Context, keyword string) ([]project.User, error) {
	var (
		results []project.User
		result  project.User
		err     error
	)

	rows, err := d.stmt[searchUserDataByKwn].QueryxContext(ctx, keyword)
	if err != nil {
		return results, errors.Wrap(err, "[DATA][SearchUserDataByKwn]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.StructScan(&result); err != nil {
			return results, errors.Wrap(err, "[DATA][SearchUserDataByKwn]")
		}
		results = append(results, result)
	}

	return results, err
}

func (d Data) SearchUserDataByKwnOrName(ctx context.Context, keyword string) ([]project.User, error) {
	var (
		results []project.User
		result  project.User
		err     error
	)

	rows, err := d.stmt[searchUserDataByKwnOrName].QueryxContext(ctx, keyword, keyword)
	if err != nil {
		return results, errors.Wrap(err, "[DATA][SearchUserDataByKwnOrName]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.StructScan(&result); err != nil {
			return results, errors.Wrap(err, "[DATA][SearchUserDataByKwnOrName]")
		}
		results = append(results, result)
	}

	return results, err
}

func (d Data) InsertDatauser(ctx context.Context, datas project.User) error {
	tx, err := d.db.Begin()
	if err != nil {
		return errors.Wrap(err, "[DATA][InsertDatauser][Begin]")
	}

	{
		if _, err := tx.ExecContext(ctx, qInsertUser,
			datas.UserName,
			datas.TanggalLahir,
			datas.Kewarganegaraan,
			datas.UserEmail,
			datas.UserID,
			datas.Telepon); err != nil {
			tx.Rollback() // return an error too, we may want to wrap them
			return errors.Wrap(err, "[DATA][InsertDatauser][qInsertUser]")
		}
	}
	{
		for _, kel := range datas.Keluarga {
			if _, err := tx.ExecContext(ctx, qInsertKeluarga,
				datas.UserID,
				kel.Hubungan,
				kel.Nama,
				kel.TanggalLahir); err != nil {
				tx.Rollback() // return an error too, we may want to wrap them
				return errors.Wrap(err, "[DATA][InsertDatauser][qInsertKeluarga]")
			}
		}
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "[DATA][InsertDatauser][Commit]")
	}

	return err
}
