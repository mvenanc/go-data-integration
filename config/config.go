package config


import (
	"fmt"
	"encoding/json"
	"os"
)

type Config struct {
	Server        string        `bson:"server json:"server"`
	Database      string        `bson:"database json:"database"`
}

func (c *Config) Read() {
	file, _ := os.Open("config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Config{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
}
