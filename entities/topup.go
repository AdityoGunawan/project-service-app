package entities

import "time"

type Topup struct {
	Topup_id    int
	No_rekening string
	Nominal     float64
	History     time.Time
}
