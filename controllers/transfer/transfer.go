package transfer

import (
	"database/sql"
	"fmt"
	"test-project/entities"
)

func Transfer(db *sql.DB, newUser entities.User, nominal float64, norekening string) (int, error) {

	oldUserPengirim := entities.User{}
	// oldUserPenerima := entities.User{}
	err := db.QueryRow("SELECT saldo FROM users WHERE no_rekening = ?", norekening).Scan(&oldUserPengirim.Saldo)
	if err != nil {
		return -1, err
	}
	err1 := db.QueryRow("SELECT saldo, no_rekening FROM users WHERE nomor_telepon = ?", newUser.No_telepon).Scan(&newUser.Saldo, &newUser.No_rekening)
	if err1 != nil {
		return -1, err
	}

	if oldUserPengirim.Saldo >= nominal {

		oldUserPengirim.Saldo -= nominal

		var querypengirim = "UPDATE users SET saldo = ? WHERE no_rekening = ?"
		statement, errPrepare := db.Prepare(querypengirim)
		if errPrepare != nil {
			return -1, errPrepare
		}
		result, errStatement := statement.Exec(oldUserPengirim.Saldo, norekening)
		if errStatement != nil {
			return -1, errStatement
		} else {
			_, errAffected := result.RowsAffected()
			if errAffected != nil {
				return 0, errAffected
			}

		}

		newUser.Saldo += nominal
		var querypenerima = "UPDATE users SET saldo = ? WHERE nomor_telepon = ?"
		statement1, errPrepare1 := db.Prepare(querypenerima)
		if errPrepare1 != nil {
			return -1, errPrepare
		}
		result1, errStatement1 := statement1.Exec(newUser.Saldo, newUser.No_telepon)
		if errStatement1 != nil {
			return -1, errStatement
		} else {
			_, errAffected := result1.RowsAffected()
			if errAffected != nil {
				return 0, errAffected
			}
		}
	} else {
		return 0, fmt.Errorf("saldo tidak mencukupi")
	}

	var queryinsert = "INSERT INTO transfer (no_rekening_pengirim, no_rekening_penerima, nominal_transfer) VALUES (?,?,?)"
	statementInsert, errPrepare := db.Prepare(queryinsert)
	if errPrepare != nil {
		return -1, errPrepare
	}
	resultinsert, errExecinsert := statementInsert.Exec(norekening, newUser.No_rekening, nominal)
	if errExecinsert != nil {
		return -1, errExecinsert
	} else {
		row, errRow := resultinsert.RowsAffected()
		if errRow != nil {
			return 0, errRow
		}
		return int(row), nil
	}

}
