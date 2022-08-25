package entities

import "time"

type Transfer struct {
	Transfer_id          int
	No_rekening_pengirim string
	No_rekening_penerima string
	Nominal_transfer     float64
	History              time.Time
}
