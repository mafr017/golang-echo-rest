package models

import (
	"net/http"

	"github.com/mafr017/rest_echo/db"
)

type Pegawai struct {
	Id int			`json:"id"`
	Nama string		`json:"nama"`
	Alamat string	`json:"alamat"`
	Telepon string	`json:"telepon"`
}

func FetchAllPegawai() (Response, error) {
	var obj Pegawai
	var arrobj []Pegawai
	var res Response

	con := db.CreateConn()

	sqlStatement := "SELECT * FROM pegawai"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Nama, &obj.Alamat, &obj.Telepon)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	return res, nil
}

func StorePegawai(nama, alamat, telepon string) (Response, error) {
	var res Response

	con := db.CreateConn()
	
	sqlStatement := "INSERT pegawai (nama, alamat, telepon) VALUES (?, ?, ?)"

	sqlpre, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := sqlpre.Exec(nama, alamat, telepon)
	if err != nil {
		return res, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}
	
	res.Status = http.StatusOK
	res.Message = "Succes"
	res.Data = map[string]int64{
		"last_inserted_id" : lastInsertId,
	}

	return res, nil
}