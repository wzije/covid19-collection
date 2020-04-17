package main

import (
	"fmt"
	"github.com/wzije/covid19-collection/jobs"
	"github.com/wzije/covid19-collection/routes"
	"log"
	"os"
	"time"
)

func Init() {
	//set default timezone to jakarta
	os.Setenv("TZ", "Asia/Jakarta")
	fmt.Print("-------------------- \n")
	fmt.Printf("Tanggal : %q \n", time.Now().String())
	fmt.Print("-------------------- \n\n")

	//run job
	jobs.RunJob()

}

func main() {
	Init()

	env := os.Getenv("APP_ENV")
	if env == "production" {
		log.Println("Running api server in production mode")
	} else {
		log.Println("Running api server in dev mode")
	}

	//run service
	log.Fatal(routes.Route().Run(":8000"))
}
