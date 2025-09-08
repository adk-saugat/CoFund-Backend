package models

import "time"

type Wallets struct{
	ID				int64
	Balance 		float64
	InterestEarned 	float64
	CreatedAt		time.Time
	userId			int64
}

func (wallet *Wallets) SaveDefault(){
	// query := `
	// 	INSERT INTO wallets ()
	// `
}