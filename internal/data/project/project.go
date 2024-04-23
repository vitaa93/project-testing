package project

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
)

type (
	Data struct {
		db   *sqlx.DB
		stmt map[string]*sqlx.Stmt
	}

	statement struct {
		key   string
		query string
	}
)

// PRODUK
const (
	getAllUser  = "GetAllUser"
	qGetAllUser = `SELECT UserName, TanggalLahir, Kewarganegaraan, UserEmail, UserID, Telepon
	FROM user`

	getKeluargaByID  = "GetKeluargaByID"
	qGetKeluargaByID = `SELECT UserID, Hubungan, Nama, TanggalLahir
	FROM keluarga
	WHERE UserID = ?`

	searchUserByNameAndttl  = "SearchUserByNameAndttl"
	qSearchUserByNameAndttl = `SELECT UserName, TanggalLahir, Kewarganegaraan, UserEmail, UserID, Telepon
	FROM user
	WHERE UserName = ? AND TanggalLahir = ?`

	getAllUserPagination  = "GetAllUserPagination"
	qGetAllUserPagination = `SELECT UserName, TanggalLahir, Kewarganegaraan, UserEmail, UserID, Telepon
	FROM user
	LIMIT ? OFFSET ?`

	getTotalAllUser  = "GetTotalAllUser"
	qGetTotalAllUser = `SELECT COUNT(*) AS Total
	FROM user`

	getUserByKwn  = "GetUserByKwn"
	qGetUserByKwn = `SELECT UserName, TanggalLahir, Kewarganegaraan, UserEmail, UserID, Telepon
	FROM user
	WHERE Kewarganegaraan = ?`

	searchUserDataByName  = "SearchUserDataByName"
	qSearchUserDataByName = `SELECT UserName, TanggalLahir, Kewarganegaraan, UserEmail, UserID, Telepon
	FROM user
	WHERE UserName LIKE '%?%`

	searchUserDataByKwn  = "SearchUserDataByKwn"
	qSearchUserDataByKwn = `SELECT UserName, TanggalLahir, Kewarganegaraan, UserEmail, UserID, Telepon
	FROM user
	WHERE Kewarganegaraan LIKE '%?%`

	searchUserDataByKwnOrName  = "SearchUserDataByKwnOrName"
	qSearchUserDataByKwnOrName = `SELECT UserName, TanggalLahir, Kewarganegaraan, UserEmail, UserID, Telepon
	FROM user
	WHERE Kewarganegaraan LIKE '%?% OR UserName LIKE '%?%'`

	insertUser  = "InsertUser"
	qInsertUser = `INSERT INTO user(UserName, TanggalLahir, Kewarganegaraan, UserEmail, UserID, Telepon)
	VALUES(?, ?, ?, ?, ?, ?)`

	insertKeluarga  = "InsertKeluarga"
	qInsertKeluarga = `INSERT INTO user(UserID, Hubungan, Nama, TanggalLahir)
	VALUES(?, ?, ?, ?)`
)

var (
	readStmt = []statement{
		{getAllUser, qGetAllUser},
		{getKeluargaByID, qGetKeluargaByID},
		{searchUserByNameAndttl, qSearchUserByNameAndttl},
		{getAllUserPagination, qGetAllUserPagination},
		{getTotalAllUser, qGetTotalAllUser},
		{getUserByKwn, qGetUserByKwn},
		{searchUserDataByName, qSearchUserDataByName},
		{searchUserDataByKwn, qSearchUserDataByKwn},
		{searchUserDataByKwnOrName, qSearchUserDataByKwnOrName},
	}
	insertStmt = []statement{
		{insertUser, qInsertUser},
		{insertKeluarga, qInsertKeluarga},
	}
	updateStmt = []statement{}
	deleteStmt = []statement{}
)

func New(db *sqlx.DB) *Data {
	var (
		stmts = make(map[string]*sqlx.Stmt)
	)

	d := &Data{
		db:   db,
		stmt: stmts,
	}

	d.InitStmt()
	return d
}

func (d *Data) InitStmt() {
	var (
		err   error
		stmts = make(map[string]*sqlx.Stmt)
	)

	for _, v := range readStmt {
		stmts[v.key], err = d.db.PreparexContext(context.Background(), v.query)
		if err != nil {
			log.Fatalf("[DB] Failed to initialize select statement key %v, err : %v", v.key, err)
		}
	}

	for _, v := range insertStmt {
		stmts[v.key], err = d.db.PreparexContext(context.Background(), v.query)
		if err != nil {
			log.Fatalf("[DB] Failed to initialize insert statement key %v, err : %v", v.key, err)
		}
	}

	for _, v := range updateStmt {
		stmts[v.key], err = d.db.PreparexContext(context.Background(), v.query)
		if err != nil {
			log.Fatalf("[DB] Failed to initialize update statement key %v, err : %v", v.key, err)
		}
	}

	for _, v := range deleteStmt {
		stmts[v.key], err = d.db.PreparexContext(context.Background(), v.query)
		if err != nil {
			log.Fatalf("[DB] Failed to initialize delete statement key %v, err : %v", v.key, err)
		}
	}

	d.stmt = stmts
}
