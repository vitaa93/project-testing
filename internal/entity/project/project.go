package project

type Keluarga struct {
	UserID       int    `db:"UserID" json:"-"`
	Hubungan     string `db:"Hubungan" json:"hubungan"`
	Nama         string `db:"Nama" json:"nama"`
	TanggalLahir string `db:"TanggalLahir" json:"tanggal_lahir"`
}

type User struct {
	Keluarga        []Keluarga `json:"keluarga"`
	UserID          int        `db:"UserID" json:"-"`
	UserName        string     `db:"UserName" json:"nama"`
	TanggalLahir    string     `db:"TanggalLahir" json:"tanggal_lahir"`
	Telepon         string     `db:"Telepon" json:"telepon"`
	Kewarganegaraan string     `db:"Kewarganegaraan" json:"kewarganegaraan"`
	UserEmail       string     `db:"UserEmail" json:"email"`
}
