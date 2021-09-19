package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"weekend.side/SocialMedia/internal/daos"
	"weekend.side/SocialMedia/internal/services"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {

	var account daos.Account
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, err)
	}

	err = json.Unmarshal(bodyBytes, &account)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, err)
	}

	resp, respErr := services.CreateAccount(account)
	if respErr != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(respErr)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(resp)
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {

}
