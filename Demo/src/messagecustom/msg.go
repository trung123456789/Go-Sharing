package messagecustom

import (
	"constants"
	"encoding/json"
	"os"
	"structdemo"

	yaml "gopkg.in/yaml.v2"
)

// MessageCustom type
type MessageCustom structdemo.MessageCustom

// Message type
type Message structdemo.Message

// MessageUtil function
func MessageUtil(msg string, errMsg string) []byte {
	var messageCustom MessageCustom
	messageCustom.Msg = msg
	messageCustom.ErrMsg = errMsg

	rs, errJSON := json.Marshal(messageCustom)
	if errJSON != nil {
		panic(errJSON)
	}
	return rs
}

// GetMessage function
func GetMessage() (message *Message) {
	f, err := os.Open(constants.MsgDic)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&message)
	if err != nil {
		panic(err)
	}
	return
}
