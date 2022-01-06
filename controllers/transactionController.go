package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"batuhand.com/api/models"
	"github.com/gorilla/mux"
)

// Returns all the transactions
func GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	transactions, err := json.Marshal(models.DummyTransactions)
	CheckError(err)
	fmt.Println(string(transactions))
	fmt.Fprint(w, string(transactions))

}

// Creates a transaction if there are enough currency in the sender wallet adress
func SendCoin(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	CheckError(err)
	transaction := models.Transaction{}
	json.Unmarshal(body, &transaction)
	senderCurrency := GetWalletAmount(transaction.SenderWalletAdress)
	if senderCurrency > transaction.Amount {
		txData, _ := json.Marshal(transaction)
		AppendTempFile(txData)
		CheckTempFileSize()
		fmt.Fprint(w, "Transaction created")
	} else {
		fmt.Fprintf(w, "Transaction failed, not enough currency")
	}

}

// Returns the currency of the given wallet adress
func GetCurrency(w http.ResponseWriter, r *http.Request) {
	var wallet_adress string = mux.Vars(r)["wallet_id"]
	currency := GetWalletAmount(wallet_adress)
	fmt.Fprintf(w, "%f", currency)

}
