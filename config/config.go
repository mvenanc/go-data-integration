package config


import (
	"fmt"
	//"log"
	"encoding/json"
	//"github.com/BurntSushi/toml"
	"os"
)

type Config struct {
	Server        string        `bson:"server json:"server"`
	Database      string        `bson:"database json:"database"`
}

func (c *Config) Read() {
	//if _, err := toml.DecodeFile("config.toml", &c); err != nil {
	//	log.Fatal(err)
	//}
	file, _ := os.Open("config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Config{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
}
