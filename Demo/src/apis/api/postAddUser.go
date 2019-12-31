package api

import (
	"net/http"
	"messagecustom"
	"responsecustom"
	"utils"
	"encoding/json"
)

// PostAddUser function
func PostAddUser(w http.ResponseWriter, r *http.Request) {
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
	query := `
		INSERT INTO
		userinfo(name, age)
		VALUES($1, $2) returning user_id`
	errQuery := database.QueryRow(query, userRequest.Name, userRequest.Age).Scan(&userResponse.ID)
	// rs, errQuery := database.Exec(query, userReuqest.Name, userReuqest.Age)

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
