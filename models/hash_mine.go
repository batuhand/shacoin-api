package models

type MinedHash struct {
	RecID           int    `json:"rec_id" sql:"recid"`
	GeneratedString string `json:"generated_string" sql:"generatedstr"`
	GeneratedHash   string `json:"generated_hash" sql:"generatedhash"`
	SenderWalletID  string `json:"sender_wallet_id" sql:"senderwallet"`
}
