package main

import (
	// "fmt"
	// "log"
	"test-project/config"
	// "database/sql"
	_ "github.com/go-sql-driver/mysql"
)
func main() {
	db := config.GetConnection()
	defer db.Close()
	
}

// type user struct {
// 	no_rekening        string
// 	nama  string
// 	password string
// 	saldo  string
// 	gender  string
// 	addres  string
// }


switch pilihan {
case 1:
	{
		// fitur register -Farhan
	}
case 2:
	{
		// fitur login -Adit
	}
case 3:
	{
		//fitur read account -Luz
	}

case 4:
	{
		//update account -farhan
	}
case 5:
	{
		//delete account -Adit
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


