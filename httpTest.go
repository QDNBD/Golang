package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	var data = []byte(`{"status": 200,
			"Err": "invalid iam req arg",
            "Reqid": "DQoAACDgAhK37ogW",
            "Details": ["PILI-LINA/400"],
            "Code": 400}`)
	var result map[string]interface{}

	if err := json.Unmarshal(data, &result); err != nil {
		log.Fatalln(err)
	}

	var status = uint64(result["Code"].(float64))
	fmt.Println("Status value: ", status)
}
