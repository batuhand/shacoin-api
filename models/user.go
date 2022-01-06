package models

type User struct {
	ID           int    `json:"id" sql:"userid"`
	Uname        string `json:"uname" sql:"username"`
	Password     string `json:"pass" sql:"password"`
	WalletAdress string `json:"wallet_adress" sql:"wallet_id"`
}

