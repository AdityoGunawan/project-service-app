package main

import (
	"fmt"
	"log"
	"test-project/entities"
	"test-project/controllers/user"
	// "test-project/config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


func main() {
	db, err := sql.Open("mysql", "root:@tcp(192.168.193.9:3306)/service_app")

	if err != nil {
		log.Fatal("error", err.Error())
	} else {
		fmt.Println("success connect to DB")
	}

	defer db.Close()

	fmt.Print("MENU:\n1. Buat Akun\n2. Login\n3. Read Data\n4. Update Data\n5. Delete Data\n")
	fmt.Println("Masukkan piihan anda:")
	var pilihan int
	fmt.Scanln(&pilihan)
	


switch pilihan {
case 1:
	{
		// fitur register -Farhan
		newUsers := entities.Users{}
		fmt.Println("Masukkan Nomor Telepon")
		fmt.Scanln(&newUsers.No_telepon)
		fmt.Println("Masukkan Nama Anda")
		fmt.Scanln(&newUsers.Nama)
		fmt.Println("Masukkan Password")
		fmt.Scanln(&newUsers.Password)
		fmt.Println("Jenis Kelamin")
		fmt.Scanln(&newUsers.Gender)
		fmt.Println("Alamat")
		fmt.Scanln(&newUsers.Addres)

		rowAffect, err := user.InsertDataUsers(db, newUsers)
		if err != nil {
			fmt.Println("Gagal Memasukkan Data", err)
		} else {
			if rowAffect == 0 {
				fmt.Println("Gagal Memasukkan Data")
			} else {
				fmt.Println("Register Berhasil")
			}
		}
	}
case 2:
	{
		// fitur login -Adit
		loginUsers := entities.Users{}
		fmt.Println("Masukkan Nomor Telepon Anda:")
		fmt.Scanln(&loginUsers.No_telepon)
		fmt.Println("Masukkan Password:")
		fmt.Scanln(&loginUsers.Password)

		rowAffect, err := user.LoginUsers(db, loginUsers)
		if err != nil {
			fmt.Println("Login Gagal", err)
		} else {
			if rowAffect < 1 {
				fmt.Println("Nomor Telepon dan Password Anda Salah")
			} else {
			fmt.Println("Login Berhasil")
			}
		}

		// 	statement := "SELECT * "
		// 	loginUsers := entities.Users{}
		// 	fmt.Scanf(&login.No_telepon)
		// 	fmt.Scanf(&login.Password)
		// 	if no_telepon == "" {

		// }
	}
case 3:
	{
		//fitur read account -Luz
		readUsers := entities.Users{}
		fmt.Println("Masukkan Nomor Telepon:")
		fmt.Scanln(&readUsers.No_telepon)

		result, err := user.GetAllUsers(db)
			if err != nil {
				fmt.Println("Data Tidak Ditemukan", err)
			} else {
				for _, v := range result {
					fmt.Println("No_telepon:", v.No_telepon, "Nama:", v.Nama, "Password:", v.Password, "Gender:", v.Gender)
				}
			}
			// results, errselect := db.Query("SELECT no_rekening, no_telepon, nama, password, saldo, gender, addres FROM users")
			// if errselect != nil {
			// 	fmt.Println("error select", errselect.Error())
			// }

			// var dataAllUsers []entities.Users //penampung semua data user
			// for results.Next() {            // membaca per baris
			// 	var rowusers entities.Users // penampung tiap baris
			// 	errScan := results.Scan(&rowusers.No_rekening, &rowusers.No_telepon, &rowusers.Nama, &rowusers.Password, &rowusers.Saldo, &rowusers.Gender, &rowusers.addres)
			// 	if errScan != nil {
			// 		fmt.Println("error scan", errScan.Error())
			// 	}
			// 	dataAllUsers = append(dataAllUsers, rowusers) //menambahkan ke slice
			// }
		
	}

case 4:
	{
		//update account -farhan
	}
case 5:
	{
		//delete account -Adit
		deleteUsers := entities.Users{}
		fmt.Println("Masukkan Nomor Telepon:")
		fmt.Scanln(&deleteUsers.No_telepon)

		rowAffect, err := user.DeleteUsers(db, deleteUsers)
		if err != nil {
			fmt.Println("Proses Gagal", err)
			return
		} else {
			if rowAffect == 0 {
				fmt.Println("Data Tidak Ditemukan!")
			} else {
				fmt.Println("Delete Berhasil")
			}
		}
	}
case 6:
	{
		//topup -Luz
	}
case 7:
	{
		//transfer -Farhan
	}
case 8:
	{
		//history topup -Adit
	}
case 9:
	{
		//history trasnfer -Luz
	}
case 10:
	{
		//melihat profil user lain -Farhan
	}
case 11:
	{
		//keluar dari sistem dan menampilkan tulisa "terimakasih telah bertransaksi"
	}
}

}


// type Users struct {
// 	No_rekening string
// 	No_telepon string
// 	Nama  string
// 	Password string
// 	Saldo  string
// 	Gender  string
// 	Addres  string
// }

//  func QueryUser(no_rekening string) Users {
// 	var user = Users{}
// 	err = db.QueryRow(`
// 	SELECT 	no_rekening,
// 	no_telepon,
// 	nama,
// 	password,
// 	saldo,
// 	gender,
// 	addres
// 	FROM  user  WHERE no_telepon=?`, no_telepon)
// 	Scan(
// 		&user.No_rekening,
// 		&user.No_telepon,
// 		&user.Nama,
// 		&user.Password,
// 		&user.Saldo,
// 		&user.Gender,
// 		&user.Addres,
// 	)
// 	return user
//  }

//  func register(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != "POST" {
// 		http.ServeFile(w, r, "views/register.html")
// 		return
// 	}
// 	no_telepon := r.FormValue("no_telepon")
// 	nama := r.FormValue("nama")
// 	password := r.FormValue("password")
// 	saldo := r.FormValue("saldo")
// 	gender := r.FormValue("gender")
// 	addres := r.FormValue("addres")

// 	user := QueryUser(no_telepon)

// 	if (Users{}) == user {
// 		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

// 		if len(hashedPassword) != 0 && checkErr(r, w, err) {
// 			stmt, err :=
// 		}
// 	}

//  }