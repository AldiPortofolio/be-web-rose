package dbmodels

import "time"

type Owner struct {
	Id  			int64 		`json:"id" gorm:"column:id"`
	OwnerFirstName  string 		`json:"owner_first_name" gorm:"column:owner_first_name"`
	OwnerLastName   string      `json:"owner_last_name" gorm:"column:owner_last_name"`
	OwnerAddress	string  	`json:"owner_address" gorm:"column:owner_address"`
	OwnerEmail      string 		`json:"owner_email" gorm:"column:owner_email"`
	OwnerJenisKelamin string 	`json:"owner_jenis_kelamin" gorm:"column:owner_jenis_kelamin"`
	OwnerKabupaten   string     `json:"owner_kabupaten" gorm:"column:owner_kabupaten"`
	OwnerKecamatan   string     `json:"owner_kecamatan" gorm:"column:owner_kecamatan"`
	OwnerKelurahan   string     `json:"owner_kelurahan" gorm:"column:owner_kelurahan"`
	OwnerKodePos     string   	`json:"owner_kode_pos" gorm:"column:owner_kode_pos"`
	OwnerNamaIbuKandung string  `json:"owner_nama_ibu_kandung" gorm:"column:owner_nama_ibu_kandung"`
	OwnerNoId        string    	`json:"owner_no_id" gorm:"column:owner_no_id"`
	OwnerNoTelp      string     `json:"owner_no_telp" gorm:"column:owner_no_telp"`
	OwnerPekerjaan   string    	`json:"owner_pekerjaan" gorm:"column:owner_pekerjaan"`
	OwnerProvinsi    string     `json:"owner_provinsi" gorm:"column:owner_provinsi"`
	OwnerRt  		 string      `json:"owner_rt" gorm:"column:owner_rt"`
	OwnerRw  		 string     `json:"owner_rw" gorm:"column:owner_rw"`
	OwnerTanggalExpiredId time.Time `json:"owner_tanggal_expired_id" gorm:"column:owner_tanggal_expired_id"`
	OwnerTanggalLahir   time.Time 	`json:"owner_tanggal_lahir" gorm:"column:owner_tanggal_lahir"`
	OwnerTelpLain     string    `json:"owner_telp_lain" gorm:"column:owner_telp_lain"`
	OwnerTempatLahir  string   `json:"owner_tempat_lahir" gorm:"column:owner_tempat_lahir"`
	OwnerTipeId       string   	`json:"owner_tipe_id" gorm:"column:owner_tipe_id"`
	OwnerTitle        string    `json:"owner_title" gorm:"column:owner_title"`
	OwnerNpwp   	  string    `json:"owner_npwp" gorm:"column:owner_npwp"`
}

// TableName ..
func (q *Owner) TableName() string {
	return "public.owner"
}
