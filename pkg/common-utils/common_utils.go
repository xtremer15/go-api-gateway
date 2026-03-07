package common_utils

import (
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

func WriteToFile(path string, dataToWrite any) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	json, err := json.MarshalIndent(dataToWrite, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString(string(json))
}

func WriteToFileBase64(path string, data any) error {
	schema, _ := json.Marshal(data)
	return os.WriteFile(path, schema, 0644)
}

func ReadJsonFile(path string) (any, error) {
	var config any
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
