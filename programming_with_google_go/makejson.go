package main

import (
	"encoding/json"
	"fmt"
)

func mainss() {
	mp := make(map[string]string)

	var name string
	var address string

	fmt.Println("Enter name:")
	fmt.Scanln(&name)

	fmt.Println("Enter address:")
	fmt.Scanln(&address)

	mp["name"] = name
	mp["address"] = address

	jsonData, err := json.Marshal(mp)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	fmt.Println("JSON data:", string(jsonData))
	fmt.Println("Unmarshalling JSON data...")
	var unmarshalledData map[string]string
	err = json.Unmarshal(jsonData, &unmarshalledData)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	fmt.Println("Unmarshalled data:", unmarshalledData)
}
