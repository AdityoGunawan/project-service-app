package main

import (
	"fmt"
	"test-project/config"
	"test-project/controllers/transfer"
	"test-project/controllers/user"
	"test-project/entities"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := config.GetConnection()
	defer db.Close()

	var loop bool = true
	var nohp string
	var norekening string
	var pilihan int
	for loop {

		if nohp == "" {
			fmt.Println("====================")
			fmt.Println("Selamat Datang di AccountServiceApp")
			fmt.Println("1. Register\n2. Login")
		} else {
			fmt.Println("===== Beranda =====\n3. Read Account\n4. Update Account\n5. Delete Account\n6. Top-up\n7. Transfer\n8. History Top-up\n9. History Tranfer\n10. Lihat Profil\n0. Keluar")
		}
		fmt.Println("====================")
		fmt.Print("Masukan Menu Pilihan Anda = ")
		fmt.Scanln(&pilihan)

		if nohp == "" {
			switch pilihan {
			case 1: // Register
				{
					newUser := entities.User{}
					fmt.Println("Input Nomor Rekening :")
					fmt.Scanln(&newUser.No_rekening)
					fmt.Println("Input Nomor Telepon :")
					fmt.Scanln(&newUser.No_telepon)
					fmt.Println("Input Nama :")
					fmt.Scanln(&newUser.Nama)
					fmt.Println("Input Password :")
					fmt.Scanln(&newUser.Password)
					fmt.Println("Input Gender :")
					fmt.Scanln(&newUser.Gender)
					fmt.Println("Input Addres")
					fmt.Scanln(&newUser.Addres)

					rowAffect, err := user.InsertDataUser(db, newUser)
					if err != nil {
						fmt.Println("error insert data", err)
					} else {
						if rowAffect == 0 {
							fmt.Println("====================")
							fmt.Println("gagal insert data row affectred = 0")
						} else {
							fmt.Println("====================")
							fmt.Println("insert sukses, row affected =", rowAffect)
						}
					}
				}

			case 2:
				{
					fmt.Println("Masukan No Telepon")
					fmt.Scanln(&nohp)
					// norekening
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

				}
			}
		} else {
			switch pilihan {
			case 3:
				{
					fmt.Println("masuk menu 3")
				}
			case 4: //update
				{
					newUser := entities.User{}
					fmt.Println(newUser)
					// fmt.Println("Input Nomor Rekening :")
					// fmt.Scanln(&newUser.No_rekening)
					newUser.No_rekening = norekening
					fmt.Println("Input Nomor Telepon :")
					fmt.Scanln(&newUser.No_telepon)
					fmt.Println("Input Nama :")
					fmt.Scanln(&newUser.Nama)
					fmt.Println("Input Password :")
					fmt.Scanln(&newUser.Password)
					fmt.Println("Input Gender :")
					fmt.Scanln(&newUser.Gender)
					fmt.Println("Input Addres")
					fmt.Scanln(&newUser.Addres)
					fmt.Println(newUser)

					rowAffect, err := user.UpdateDataUser(db, newUser)
					if err != nil {
						fmt.Println("error update data", err)
					} else {
						if rowAffect == 0 {
							fmt.Println("Tidak ada data yang di update")
						} else {
							fmt.Println("Update data sukses, row affected =", rowAffect)
						}
					}
				}
			case 5:
				{
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

				}
			case 7:
				{
					newUser := entities.User{}
					var nominal float64
					fmt.Println("Masukan Nomor Telepon Tujuan :")
					fmt.Scanln(&newUser.No_telepon)
					fmt.Println("Masukan Nominal Transfer :")
					fmt.Scanln(&nominal)

					affectedRow, err := transfer.Transfer(db, newUser, nominal, norekening)
					if err != nil {
						fmt.Println(err.Error())
						break
					}
					if affectedRow < 1 {
						fmt.Println("Transfer Gagal")
						break
					}

					fmt.Println("Transfer Berhasil")

				}
			case 8:
				{

				}
			case 9:
				{

				}
			case 10:
				{
					newUser := entities.User{}
					fmt.Println("Masukkan No Telepon :")
					fmt.Scanln(&newUser.No_telepon)
					if newUser.No_telepon == "" {
						fmt.Println("wajib memasukkan ID User.")
						break
					}

					affected, err := user.GetUser(db, newUser.No_telepon)
					if err != nil {
						fmt.Println(err.Error())
					}

					if affected.No_telepon != "" {
						fmt.Println("====================")
						fmt.Printf("No Rekening: %s\nNo Telepon: %s\nNama: %s\nJenis Kelamin: %s\nAlamat: %s\n", affected.No_rekening, affected.No_telepon, affected.Nama, affected.Gender, affected.Addres)
						break
					}
					fmt.Println("Data tidak ditemukan.")
				}
			case 0:
				{
					nohp = ""
					fmt.Println("====================")
					fmt.Println("Terimakasih telah bertransaksi")
				}
			}

		}

		var yn string
		fmt.Println("====================")
		fmt.Println("Apakah Ingin Melanjutkan Transkasi? y/n")
		fmt.Scanln(&yn)
		if yn != "y" {
			loop = false
		}
	}
}
