package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"batuhand.com/api/models"
	"batuhand.com/api/utils"
)

func UploadHash(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	CheckError(err)
	minedHash := models.MinedHash{}
	json.Unmarshal(body, &minedHash)
	fmt.Println(minedHash)
	_, err = utils.Db.Exec(`INSERT INTO hashtable (generatedstr, generatedhash, senderwallet)
    VALUES ($1, $2, $3)`, minedHash.GeneratedString, minedHash.GeneratedHash, minedHash.SenderWalletID)
	CheckError(err)
	RewardWallet(minedHash.SenderWalletID)
	fmt.Fprint(w, "{'status': 'ok', 'received' : '1.0'}")

}

func RewardWallet(wallet_id string) {
	tx := models.Transaction{
		SenderWalletAdress:   "0xbatuhanthecreator",
		RecieverWalletAdress: wallet_id,
		Amount:               1,
	}
	txData, _ := json.Marshal(tx)
	AppendTempFile(txData)
	CheckTempFileSize()
}
