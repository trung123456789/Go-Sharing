package api

import (
	"database-demo/connection"
	"encoding/json"
	"fmt"
	"messagecustom"
	"net/http"
	"responsecustom"
	"strconv"

	"github.com/gorilla/mux"
)

// DeleteUser function
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	var userResponse UserResponse

	userID := -1
	var err error
	if val, ok := pathParams["userID"]; ok {
		userID, err = strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "need a number"}`))
			return
		}
	}
	database, errConnectDb := connection.CreateConnection()
	if errConnectDb != nil {
		fmt.Println(errConnectDb)
		messageUtil := messagecustom.MessageUtil(
			messagecustom.GetMessage().Msg.ConnectErr,
			messagecustom.GetMessage().ErrMsg.ConnectErr)
		responsecustom.ResponseCustom(w, http.StatusInternalServerError, messageUtil)
		return
	}

	// Check user exist
	querySelect := `
		SELECT user_id
		FROM user_infos
		WHERE user_id = $1
	`
	var id int
	errSelect := database.QueryRow(querySelect, userID).Scan(&id)
	if errSelect != nil {
		messageUtil := messagecustom.MessageUtil(
			messagecustom.GetMessage().Msg.SearchErr,
			messagecustom.GetMessage().ErrMsg.SearchErr)
		responsecustom.ResponseCustom(w, http.StatusInternalServerError, messageUtil)
		return
	}

	// Update user
	queryUpdate := `
		DELETE
		FROM user_infos 
		WHERE user_id = $1 returning user_id`
	errQuery := database.QueryRow(queryUpdate, userID).Scan(&userResponse.ID)

	if errQuery != nil {
		messageUtil := messagecustom.MessageUtil(
			messagecustom.GetMessage().Msg.DatabseErr,
			messagecustom.GetMessage().ErrMsg.DatabseErr)
		responsecustom.ResponseCustom(w, http.StatusInternalServerError, messageUtil)
		return
	}

	userRs, errJSON := json.Marshal(userResponse)
	if errJSON != nil {
		messageUtil := messagecustom.MessageUtil(
			messagecustom.GetMessage().Msg.JSONErr,
			messagecustom.GetMessage().ErrMsg.JSONErr)
		responsecustom.ResponseCustom(w, http.StatusInternalServerError, messageUtil)
		return
	}
	responsecustom.ResponseCustom(w, http.StatusOK, userRs)
}
