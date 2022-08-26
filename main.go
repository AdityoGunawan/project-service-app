package main

import (
	"fmt"
	"test-project/config"
	"test-project/controllers/topup"
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
	var validasi string
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
					newUser := entities.User{}
					fmt.Println("Masukan No Telepon")
					fmt.Scanln(&newUser.No_telepon)
					fmt.Println("Masukan Password")
					fmt.Scanln(&newUser.Password)

					_, login, norek, notelepon, _ := user.LoginUser(db, newUser)
					nohp = login
					norekening = norek
					validasi = notelepon
				}
			}
		} else if nohp == "login sukses" {
			switch pilihan {
			case 3:
				{
					newUser := entities.User{}
					// fmt.Println("Masukkan No Telepon :")
					newUser.No_rekening = norekening
					affected, err := user.ReadUser(db, newUser)
					if err != nil {
						fmt.Println(err.Error())
					}

					if affected.No_telepon != "" {
						fmt.Println("====================")
						fmt.Printf("No Rekening: %s\nNo Telepon: %s\nNama: %s\nPassword: %s\nSaldo: %v\nJenis Kelamin: %s\nAlamat: %s\n", affected.No_rekening, affected.No_telepon, affected.Nama, affected.Password, affected.Saldo, affected.Gender, affected.Addres)
						break
					}
					fmt.Println("Data tidak ditemukan.")
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
					newUser := entities.User{}
					// fmt.Println("Masukkan Guru ID")
					// fmt.Scanln(&newGuru.ID)
					newUser.No_rekening = norekening

					affected, err := user.DeleteDataUser(db, newUser)
					if err != nil {
						fmt.Println(err.Error())
					}

					if affected > 0 {
						fmt.Println("Data berhasil di hapus.")
						break
					} else {
						fmt.Println("Tidak ada data yang dihapus.")
						break
					}
				}
			case 6:
				{
					newUser := entities.User{}
					var nominal float64
					fmt.Println("Masukan Nomor Telepon Anda :")
					fmt.Scanln(&newUser.No_telepon)
					if newUser.No_telepon != validasi {
						fmt.Println("Nomor Telepon Anda Salah")
						break
					}
					fmt.Println("Masukan Nominal topup :")
					fmt.Scanln(&nominal)

					_, err := topup.Topup(db, newUser, nominal, norekening)
					if err != nil {
						fmt.Println(err.Error())
						break
					}
					fmt.Println("Topup Berhasil")
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
					result, err := topup.TopupHistory(db)
					if err != nil {
						fmt.Println(err.Error())
					} else {
						for _, v := range result {
							fmt.Println("Nominal:", v.Nominal, "Date:", v.History)
						}
					}
				}
			case 9:
				{
					result, err := transfer.TransferHistory(db)
					if err != nil {
						fmt.Println(err.Error())
					} else {
						for _, v := range result {
							fmt.Println("Penerima:", v.No_rekening_penerima, "Nominal:", v.Nominal_transfer, "Date:", v.History)
						}
					}
				}
			case 10:
				{
					newUser := entities.User{}
					fmt.Println("Masukkan No Telepon :")
					fmt.Scanln(&newUser.No_telepon)
					if newUser.No_telepon == "" {
						fmt.Println("wajib memasukkan No Telepon.")
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