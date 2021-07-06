package utils

import (
	"encoding/json"
	"log"
)

func Marshal(v interface{}) string {
	bytes, err := json.Marshal(v)
	if err != nil {
		log.Fatal("Error while marshalling: ", err)
	}

	return string(bytes)
}

func Unmarshal(data []byte, v interface{}) {
	err := json.Unmarshal(data, v)
	if err != nil {
		log.Fatal("Error while unmarshalling: ", err)
	}
}
