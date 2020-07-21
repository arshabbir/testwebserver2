package main

import (
	"log"
	"os"

	"github.com/arshabbir/docker1/src/app"
)

//import "github.com/arshabbir/docker1/src/app"

func main() {

	dbhost := os.Getenv("DBHOST")
	logpath := os.Getenv("LOGPATH")

	log.Println(dbhost, logpath)

	app.StartApplication()

}
