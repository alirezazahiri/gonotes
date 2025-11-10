package jsonutils

import (
	"encoding/json"
	"os"
)

func WriteJSONFile(data interface{}, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(data)
}

func ReadJSONFile(filename string) (interface{}, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data interface{}
	err = json.NewDecoder(file).Decode(&data)

	return data, err
}
