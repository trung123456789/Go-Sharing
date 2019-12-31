package utils

import (
	"database-demo/connection"
	"messagecustom"
	"responsecustom"
	"database/sql"
	"io/ioutil"
	"net/http"
)

// InitialConnect function
func InitialConnect(w http.ResponseWriter, r *http.Request) (*sql.DB, []byte) {
	database, errConnectDb := connection.CreateConnection()
	if errConnectDb != nil {
		messageUtil := messagecustom.MessageUtil(
			messagecustom.GetMessage().Msg.DatabseErr,
			messagecustom.GetMessage().ErrMsg.DatabseErr)
		responsecustom.ResponseCustom(w, http.StatusInternalServerError, messageUtil)
		return nil, nil
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return nil, nil
	}
	return database, body
}
