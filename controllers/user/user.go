package user

import (
	"test-project/entities"
	"database/sql"
	"fmt"
) 

func GetAllUsers(db *sql.DB) ([]entities.Users, error) {

	results, errselect := db.Query("SELECT * FROM users WHERE nomor_telepon=?")
	if errselect != nil {
		// fmt.Println("error select", errselect.Error())
		return nil, errselect
	}

	var dataAllUsers []entities.Users //penampung semua data user
	for results.Next() {            // membaca per baris
		var rowusers entities.Users // penampung tiap baris
		errScan := results.Scan(&rowusers.No_telepon, &rowusers.Nama, &rowusers.Password, &rowusers.Gender, &rowusers.Addres)
		if errScan != nil {
			// fmt.Println("error scan", errScan.Error())
			return nil, errScan
		}
		// fmt.Println("row", rowusers)
		dataAllUsers = append(dataAllUsers, rowusers) //menambahkan ke slice
	}
	// fmt.Println(dataAllUsers)
	// for _, v := range dataAllUsers {
		// fmt.Println("No_rekening:", v.No_rekening, "No_telepon:", v.No_telepon, "Nama:", v.Nama, "Password:", v.Password, "Gender:", v.Gender)
	// }
	return dataAllUsers, nil
	
	}



func InsertDataUsers(db *sql.DB, newUsers entities.Users) (int, error) {
	var query = "insert into users (nomor_telepon, nama, password, gender, addres) values (?, ?, ?, ?, ?)"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		return -1, errPrepare
	}

	result, errExec := statement.Exec(newUsers.No_telepon, newUsers.Nama, newUsers.Password, newUsers.Gender, newUsers.Addres)
	if errExec != nil {
		return -1, errExec

	} else {
		row, errRow := result.RowsAffected()
		if errRow != nil {
			return 0, nil
		}
		return int(row), nil
	}

}


func LoginUsers(db *sql.DB, loginUsers entities.Users) (int, error) {
	var login = "select * from users WHERE nomor_telepon=? and password=?"
	_, errPrepare := db.Prepare(login)
	if errPrepare != nil {
		fmt.Println("Nomor Telepon atau Password yang Anda Masukkan Salah")
		return 0, errPrepare
	}
	return 1, nil

}

func DeleteUsers(db *sql.DB, deleteUsers entities.Users) (int, error) {
	var delete = "DELETE FROM users WHERE nomor_telepon = ?"
	_, errPrepare := db.Prepare(delete)
	if errPrepare != nil {
		fmt.Println("Hapus Akun Gagal")
		return 0, errPrepare
	}
	return 0, nil
}
