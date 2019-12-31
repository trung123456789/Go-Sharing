package api

import (
	"net/http"
	"messagecustom"
	"responsecustom"
	"fmt"
	"database/sql"
	"encoding/json"
	"bytes"
	"utils"
	"log"
)


// PostGetUser function
func PostGetUser(w http.ResponseWriter, r *http.Request) {
	var userList []User
	var userRequest UserRequest
	database, body := utils.InitialConnect(w, r)
	defer database.Close()
	err := json.Unmarshal(body, &userRequest)
	if err != nil {
		log.Println("Err: ", err)
		messageUtil := messagecustom.MessageUtil(
			messagecustom.GetMessage().Msg.JSONErr,
			messagecustom.GetMessage().ErrMsg.JSONErr)
		responsecustom.ResponseCustom(w, http.StatusNotFound, messageUtil)
		return
	}
	// Check user exist
	var query bytes.Buffer
	if countUser(query, userRequest, database) == 0 {
		messageUtil := messagecustom.MessageUtil(
			messagecustom.GetMessage().Msg.SearchErr,
			"User does not exist")
		responsecustom.ResponseCustom(w, http.StatusNotFound, messageUtil)
		return
	}
	rows := selectUser(w, query, userRequest, database)
	defer rows.Close()

	for rows.Next() {
		user := User{}
		if err := rows.Scan(&user.ID, &user.Name, &user.Age); err != nil {
			log.Println(err)
			messageUtil := messagecustom.MessageUtil(
				messagecustom.GetMessage().Msg.ServerErr,
				messagecustom.GetMessage().ErrMsg.ServerErr)
			responsecustom.ResponseCustom(w, http.StatusNotFound, messageUtil)
			return
		}

		userList = append(userList, user)
	}

	userRs, errJSON := json.Marshal(userList)
	if errJSON != nil {
		messageUtil := messagecustom.MessageUtil(
			messagecustom.GetMessage().Msg.JSONErr,
			messagecustom.GetMessage().ErrMsg.JSONErr)
		responsecustom.ResponseCustom(w, http.StatusInternalServerError, messageUtil)
		return
	}
	responsecustom.ResponseCustom(w, http.StatusOK, userRs)
}

func queryUtil(query bytes.Buffer, request UserRequest) bytes.Buffer {
	query.WriteString("WHERE 1 = 1")
	var queryTemp string
	if request.ID != 0 {
		queryTemp = fmt.Sprintf(" AND user_id = %d", request.ID)
		query.WriteString(queryTemp)
	}
	if request.Name != "" {
		queryTemp = fmt.Sprintf(" AND name = '%s'", request.Name)
		query.WriteString(queryTemp)
	}
	if request.Age != 0 {
		queryTemp = fmt.Sprintf(" AND age = %d", request.Age)
		query.WriteString(queryTemp)
	}
	return query
}


func countUser(query bytes.Buffer, userRequest UserRequest, database *sql.DB) int {
	var count int
	queryCount := `
		SELECT COUNT(*)
		FROM userinfo
	`
	query.WriteString(queryCount)
	query = queryUtil(query, userRequest)
	err := database.QueryRow(query.String()).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	return count
}

func selectUser(w http.ResponseWriter, query bytes.Buffer, userRequest UserRequest, database *sql.DB) *sql.Rows {
	querySelect := `
		SELECT *
		FROM userinfo
	`
	query.WriteString(querySelect)

	query = queryUtil(query, userRequest)
	rows, errSelect := database.Query(query.String())
	if errSelect != nil {
		log.Println(errSelect)
		messageUtil := messagecustom.MessageUtil(
			messagecustom.GetMessage().Msg.NotFound,
			messagecustom.GetMessage().ErrMsg.NotFound)
		responsecustom.ResponseCustom(w, http.StatusNotFound, messageUtil)
		return nil
	}
	return rows
}