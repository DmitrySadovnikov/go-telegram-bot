package main

import (
	"wat-r-u-doing-bot/config"
	"fmt"
	log "github.com/Sirupsen/logrus"
	_ "github.com/lib/pq"
	"os"
)

func main() {
	db, err := config.OpenSQLConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var commandArray [2]string

	if os.Args[1] == "drop" {
		commandArray = [2]string{"DROP", "deleted"}
	} else {
		commandArray = [2]string{"CREATE", "created"}
	}

	_, err = db.Exec(commandArray[0] + " DATABASE " + os.Getenv("POSTGRES_DNAME"))

	if err != nil {
		panic(err)
	}
	log.Info(fmt.Sprintf("База %v %v", os.Getenv("POSTGRES_DNAME"), commandArray[1]))

	db.Close()
}
