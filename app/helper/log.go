package helper

import (
	"encoding/json"
	"log"
)

func LogStruct(prefix string, obj interface{}) {
	res, err := json.MarshalIndent(obj, "", "\t")
	var p string
	if err != nil {
		p = "nothing"
	}
	p = string(res)
	log.Print(prefix, "=======")
	log.Print(p)
}
