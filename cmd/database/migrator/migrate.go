package main

import (
	"fmt"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
)

func init() {
	godotenv.Load()
}

func main() {
	var command = []string{os.Args[1]}

	commandSlice := []string{
		"-database",
		fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v",
			os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DNAME"), os.Getenv("POSTGRES_SSLMODE")),
		"-path",
		os.Getenv("GOPATH") + "/src/go-telegram-bot/db/migrate",
	}

	out, err := exec.Command(os.Getenv("GOPATH")+"/bin/migrate", append(commandSlice, command...)...).Output()

	if err != nil {
		log.Fatal(err)
	}
	log.Info(string(out))
}
