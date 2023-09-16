package init_process

import (
	"encoding/json"
	"os"
)

var Myconfig Config

func init() {
	//config
	var jsonFile []byte
	if jsonByte, err := os.ReadFile("./config.json"); err != nil {
		panic(err)
	} else {
		jsonFile = jsonByte
	}
	if err := json.Unmarshal([]byte(jsonFile), &Myconfig); err != nil {
		panic(err)
	}
	if os.Getenv("MODE") == "DEV" {
		Myconfig.Debug = true
	} else {
		Myconfig.Debug = false
	}
}
