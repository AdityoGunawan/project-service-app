package user

import (
	"database/sql"
	"fmt"
	"test-project/entities"
)

//Register User
func InsertDataUser(db *sql.DB, newUser entities.User) (int, error) {

	var query = "INSERT INTO users (no_rekening, nomor_telepon, nama, password, gender, addres) VALUES (?,?,?,?,?,?)"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		return -1, errPrepare
	}

	result, errExec := statement.Exec(&newUser.No_rekening, &newUser.No_telepon, &newUser.Nama, &newUser.Password, &newUser.Gender, &newUser.Addres)
	if errExec != nil {
		return -1, errExec
	} else {
		row, errRow := result.RowsAffected()
		if errRow != nil {
			return 0, errRow
		}
		return int(row), nil
	}

}

//Update User
func UpdateDataUser(db *sql.DB, newUser entities.User) (int, error) {
	oldUser := entities.User{}

	err := db.QueryRow("SELECT * FROM users WHERE no_rekening = ?", newUser.No_rekening).Scan(&oldUser.No_rekening, &oldUser.No_telepon, &oldUser.Nama, &oldUser.Password, &oldUser.Saldo, &oldUser.Gender, &oldUser.Addres)
	if err != nil {
		return -1, err
	}
	if oldUser.No_rekening == "" {
		newUser.No_rekening = oldUser.No_rekening
	}
	if newUser.No_telepon == "" {
		newUser.No_telepon = oldUser.No_telepon
	}
	if newUser.Nama == "" {
		newUser.Nama = oldUser.Nama
	}
	if newUser.Password == "" {
		newUser.Password = oldUser.Password
	}
	if newUser.Gender == "" {
		newUser.Gender = oldUser.Gender
	}
	if newUser.Addres == "" {
		newUser.Addres = oldUser.Addres
	}

	var query = "UPDATE users SET nomor_telepon = ?, nama = ?, password = ?, gender = ?, addres = ? WHERE no_rekening = ?"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		return -1, errPrepare
	}

	fmt.Println(newUser)
	result, errStatement := statement.Exec(newUser.No_telepon, newUser.Nama, newUser.Password, newUser.Gender, newUser.Addres, newUser.No_rekening)
	if errStatement != nil {
		return -1, errStatement
	} else {
		row, errAffected := result.RowsAffected()
		if errAffected != nil {
			return 0, errAffected
		}
		return int(row), nil
	}

}

//melihat profil user lain
func GetUser(db *sql.DB, nomortelepon string) (entities.User, error) {
	newUser := entities.User{}
	err := db.QueryRow("SELECT no_rekening, nomor_telepon, nama, gender, addres FROM users WHERE nomor_telepon = ?", nomortelepon).Scan(&newUser.No_rekening, &newUser.No_telepon, &newUser.Nama, &newUser.Gender, &newUser.Addres)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}
