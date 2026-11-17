package main

import (
	"fmt"

	"github.com/conceptoHH/gator/internal/config"
)

func main() {
	conf := config.Read()

	fmt.Println("Current user: ",conf.CurrentUserName)
	conf.SetUser("user2")

	conf2 := config.Read()
	fmt.Println(conf2.DbURL, conf2.CurrentUserName)
}
