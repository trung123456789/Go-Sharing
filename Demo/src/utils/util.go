package utils

import (
	"database-demo/connection"
	"database/sql"
	"io/ioutil"
	"messagecustom"
	"net/http"
	"responsecustom"
	"structdemo"
)

// UserInfo type
type UserInfo structdemo.UserInfo

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

// AutoMigrateDb func
func AutoMigrateDb() {
	var w http.ResponseWriter
	database, errConnectDb := connection.GormConnection()
	if errConnectDb != nil {
		messageUtil := messagecustom.MessageUtil(
			messagecustom.GetMessage().Msg.DatabseErr,
			messagecustom.GetMessage().ErrMsg.DatabseErr)
		responsecustom.ResponseCustom(w, http.StatusInternalServerError, messageUtil)
	}
	// Migrate the schema
	database.AutoMigrate(&UserInfo{})
}
