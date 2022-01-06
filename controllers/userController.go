package controllers

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"batuhand.com/api/models"
	"batuhand.com/api/utils"
)

// Returns all users
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := utils.Db.Query("SELECT * FROM shacoin_users;")
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No Records Found")
			return
		}
	}
	defer rows.Close()
	var users []*models.User
	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(&user.ID, &user.Uname, &user.Password, &user.WalletAdress)
		if err != nil {
			fmt.Println("saaaa-------")
			log.Fatal(err)

		}
		users = append(users, user)
		fmt.Println(users)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	for _, usr := range users {
		fmt.Printf("%d - %s, %s, %s\n", usr.ID, usr.Uname, usr.Password, usr.WalletAdress)
	}
	output, _ := json.Marshal(users)
	fmt.Fprint(w, string(output))

}

// Returns spesific user by id
func GetUser(w http.ResponseWriter, r *http.Request) {
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	CheckError(err)
	fmt.Println(id)
	rows, err := utils.Db.Query(fmt.Sprintf("SELECT * FROM shacoin_users where userid= %d", id))
	CheckError(err)
	var users []*models.User
	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(&user.ID, &user.Uname, &user.Password, &user.WalletAdress)
		if err != nil {
			fmt.Println("saaaa-------")
			log.Fatal(err)

		}
		users = append(users, user)
		fmt.Println(users)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	for _, usr := range users {
		fmt.Printf("%d - %s, %s, %s\n", usr.ID, usr.Uname, usr.Password, usr.WalletAdress)
	}
	output, _ := json.Marshal(users)
	fmt.Fprint(w, string(output))
}

// Updates user by id
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	url_list := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(url_list[4])
	fmt.Println(id)

	CheckError(err)
	body, err := ioutil.ReadAll(r.Body)
	CheckError(err)
	newUser := models.User{}
	json.Unmarshal(body, &newUser)

}

// For removing object from the given index
func RemoveIndex(users []models.User, index int) []models.User {
	return append(users[:index], users[index+1:]...)
}

// Deletes user by id
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	CheckError(err)
	fmt.Println(id)

	_, err2 := utils.Db.Query(fmt.Sprintf("DELETE FROM shacoin_users where userid= %d", id))
	if err2 != nil {
		fmt.Fprintf(w, "Error")
	} else {
		fmt.Fprintf(w, "User deleted")

	}

}

// Creates user 
func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	CheckError(err)
	user := models.User{}
	json.Unmarshal(body, &user)
	_, err = utils.Db.Exec(`INSERT INTO shacoin_users (username, password, wallet_id)
    VALUES ($1, $2, $3)
    RETURNING userid`, user.Uname, user.Password, user.WalletAdress)
	CheckError(err)
	fmt.Fprint(w, "User created")

}
