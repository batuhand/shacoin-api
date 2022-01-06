package models

type Transaction struct {
	SenderWalletAdress   string  `json:"sender_wallet"`
	RecieverWalletAdress string  `json:"reciever_wallet"`
	Amount               float64 `json:"amount"`
}

var DummyTransactions = []Transaction{
	{SenderWalletAdress: "0x123dsadsdasd", RecieverWalletAdress: "0x91230sf90sd9s", Amount: 258.23},
	{SenderWalletAdress: "0x91230sf90sd9s", RecieverWalletAdress: "0x321skadlkfna", Amount: 23.23},
	{SenderWalletAdress: "0xs8df7sd9f87sd", RecieverWalletAdress: "0x321skadlkfna", Amount: 543.23},
	{SenderWalletAdress: "0x67sd8f6ds78f6sd", RecieverWalletAdress: "0x91230sf90sd9s", Amount: 123.23},
}
