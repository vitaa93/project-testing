package project

type Keluarga struct {
	UserID       int    `db:"UserID" json:"userid"`
	Hubungan     string `db:"Hubungan" json:"hubungan"`
	Nama         string `db:"Nama" json:"nama"`
	TanggalLahir string `db:"TanggalLahir" json:"tanggal_lahir"`
}

type User struct {
	UserID          int        `db:"UserID" json:"userid"`
	UserName        string     `db:"UserName" json:"nama"`
	TanggalLahir    string     `db:"TanggalLahir" json:"tanggal_lahir"`
	Telepon         string     `db:"Telepon" json:"telepon"`
	Kewarganegaraan string     `db:"Kewarganegaraan" json:"kewarganegaraan"`
	UserEmail       string     `db:"UserEmail" json:"email"`
	Keluarga        []Keluarga `json:"keluarga"`
}
