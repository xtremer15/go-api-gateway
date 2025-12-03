package common_utils

import (
	common_types "api-gateway/pkg/types"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func ReadFile(path string) string {
	fileContent, err := os.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(string(fileContent))

	return string(fileContent)
}

func ReadJsonFile(path string) (*common_types.DBCreds, error) {
	var config *common_types.DBCreds
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil, err
	}
	fmt.Println(config)
	return config, nil
}
