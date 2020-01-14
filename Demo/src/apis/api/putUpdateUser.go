package api

import (
	"encoding/json"
	"log"
	"messagecustom"
	"net/http"
	"responsecustom"
	"utils"
)

// PutUpdateUser function
func PutUpdateUser(w http.ResponseWriter, r *http.Request) {
	var userRequest UserRequest
	var userResponse UserResponse
	database, body := utils.InitialConnect(w, r)
	err := json.Unmarshal(body, &userRequest)
	if err != nil {
		messageUtil := messagecustom.MessageUtil(
			messagecustom.GetMessage().Msg.JSONErr,
			messagecustom.GetMessage().ErrMsg.JSONErr)
		responsecustom.ResponseCustom(w, http.StatusNotFound, messageUtil)
		return
	}
	// Check user exist
	querySelect := `
		SELECT *
		FROM user_infos
		WHERE user_id = $1
	`
	row, errSelect := database.Query(querySelect, userRequest.ID)
	if errSelect != nil {
		messageUtil := messagecustom.MessageUtil(
			messagecustom.GetMessage().Msg.SearchErr,
			messagecustom.GetMessage().ErrMsg.SearchErr)
		responsecustom.ResponseCustom(w, http.StatusInternalServerError, messageUtil)
		return
	}
	defer row.Close()

	// Update user
	queryUpdate := `
		UPDATE userinfo
		SET 
			name = $1,
			age = $2
		WHERE user_id = $3 returning user_id`
	var lastInsertID int
	errQuery := database.QueryRow(queryUpdate, userRequest.Name, userRequest.Age, userRequest.ID).Scan(&userResponse.ID)
	log.Println(lastInsertID)

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
