
package topup

import (
	"database/sql"
	"fmt"
	"test-project/entities"
)

func Topup(db *sql.DB, newUser entities.User, nominal float64, norekening string) (int, error) {

	oldUser := entities.User{}
	// oldUserPenerima := entities.User{}
	err := db.QueryRow("SELECT saldo FROM users WHERE no_rekening = ?", norekening).Scan(&oldUser.Saldo)
	if err != nil {
		return -1, err
	}

	oldUser.Saldo += nominal
	var querypengirim = "UPDATE users SET saldo = ? WHERE no_rekening = ?"
	statement, errPrepare := db.Prepare(querypengirim)
	if errPrepare != nil {
		return -1, errPrepare
	}
	result, errStatement := statement.Exec(oldUser.Saldo, norekening)
	if errStatement != nil {
		return -1, errStatement
	} else {
		_, errAffected := result.RowsAffected()
		if errAffected != nil {
		}
	}

	var queryinsert = "INSERT INTO topup (no_rekening, nominal_topup) VALUES (?,?)"
	statementInsert, errPrepare := db.Prepare(queryinsert)
	if errPrepare != nil {
		return -1, errPrepare
	}
	resultinsert, errExecinsert := statementInsert.Exec(norekening, nominal)
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

func TopupHistory(db *sql.DB) ([]entities.Topup, error) {
	result, errselect := db.Query("SELECT nominal_topup, history_topup FROM topup")
	if errselect != nil {
		fmt.Println("error select", errselect.Error())
		return nil, errselect
	}
	var dataTopup []entities.Topup
	for result.Next() {
		var rowTopup entities.Topup
		errScan := result.Scan(&rowTopup.Nominal, &rowTopup.History)
		if errScan != nil {
			return nil, errScan
		}
		dataTopup = append(dataTopup, rowTopup)
	}
	return dataTopup, nil
}
=======
package topup

import (
	"database/sql"
	"fmt"
	"test-project/entities"
)

func Topup(db *sql.DB, newUser entities.User, nominal float64, norekening string) (int, error) {

	oldUser := entities.User{}
	// oldUserPenerima := entities.User{}
	err := db.QueryRow("SELECT saldo FROM users WHERE no_rekening = ?", norekening).Scan(&oldUser.Saldo)
	if err != nil {
		return -1, err
	}

	oldUser.Saldo += nominal
	var querypengirim = "UPDATE users SET saldo = ? WHERE no_rekening = ?"
	statement, errPrepare := db.Prepare(querypengirim)
	if errPrepare != nil {
		return -1, errPrepare
	}
	result, errStatement := statement.Exec(oldUser.Saldo, norekening)
	if errStatement != nil {
		return -1, errStatement
	} else {
		_, errAffected := result.RowsAffected()
		if errAffected != nil {
		}
	}

	var queryinsert = "INSERT INTO topup (no_rekening, nominal_topup) VALUES (?,?)"
	statementInsert, errPrepare := db.Prepare(queryinsert)
	if errPrepare != nil {
		return -1, errPrepare
	}
	resultinsert, errExecinsert := statementInsert.Exec(norekening, nominal)
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

func TopupHistory(db *sql.DB) ([]entities.Topup, error) {
	result, errselect := db.Query("SELECT nominal_topup, history_topup FROM topup")
	if errselect != nil {
		fmt.Println("error select", errselect.Error())
		return nil, errselect
	}
	var dataTopup []entities.Topup
	for result.Next() {
		var rowTopup entities.Topup
		errScan := result.Scan(&rowTopup.Nominal, &rowTopup.History)
		if errScan != nil {
			return nil, errScan
		}
		dataTopup = append(dataTopup, rowTopup)
	}
	return dataTopup, nil
}
