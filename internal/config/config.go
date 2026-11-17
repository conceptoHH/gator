package config

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

type Config struct {
	DbURL 			string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() *Config {
	cfg := Config{}
	home,err := os.UserHomeDir() 
	if err != nil {
		fmt.Printf("ERROR GETTING HOME DIR")
	}
	file,err := os.Open(path.Join(home,"REPOS/bootdev/gator/.gatorconfig.json"))
	if err != nil {
		fmt.Printf("ERROR OPENING THE %v FILE", "gatorconfing.json")
	}
	data,err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("ERROR CONVERTING TO BYTES THE FILE: %v", err)
	}
	if err := json.Unmarshal(data,&cfg); err != nil {
		fmt.Printf("ERROR UNMARSHALLING THE DATA: %v", err)
	}
	return &cfg
}

func (c *Config) SetUser(usr string) {
	c.CurrentUserName = usr

	fmt.Println(c)

	data, err := json.Marshal(c)
	if err != nil {
		fmt.Printf("ERROR MARSHALLING THE STRUCT: %v", err)
	}
	home,err := os.UserHomeDir() 
	if err != nil {
		fmt.Printf("ERROR GETTING HOME DIR")
	}
	file,err := os.OpenFile(path.Join(home,"REPOS/bootdev/gator/.gatorconfig.json"),os.O_RDWR,0644)
	if err != nil {
		fmt.Printf("ERROR OPENING THE %v FILE", "gatorconfing.json")
	}
	a,err := file.Write(data)
	if err != nil {
		fmt.Printf("num lines: %v",a)
		log.Fatal(err)
	}
}
