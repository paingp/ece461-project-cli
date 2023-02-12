package ratom

import (
	"fmt"
	"log"
	"os"

	//"github.com/joho/godotenv"
	"github.com/paingp/ece461-project-cli/ratom/metrics"
)

func LoggerVerbOne(output []Module) {

	// error := godotenv.Load(".env")
	// if error != nil {
	// 	panic("Error loading .env file")
	// }

	logPath := os.Getenv("LOG_FILE")

	f, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("GoLang could not find logging file!")
		panic(err)
	}

	defer f.Close()

	// logger := log.New(f, "", log.LstdFlags)
	log.SetOutput(f)
	log.SetPrefix("")
	log.SetFlags(0)
	//logger.Println("")

	log.Printf("\n")
	log.Printf("\nOutput:\n")
	for a := 0; a < len(output); a++ {
		log.Println(output[a])
	}
	log.Printf("\n")
	log.Printf("\n")
}

func LoggerVerbTwo(output []Module) {
	// error := godotenv.Load(".env")
	// if error != nil {
	// 	panic("Error loading .env file")
	// }

	logPath := os.Getenv("LOG_FILE")

	f, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("GoLang could not find logging file!")
		panic(err)
	}

	defer f.Close()

	// logger := log.New(f, "", log.LstdFlags)
	log.SetOutput(f)
	log.SetPrefix("")
	log.SetFlags(0)
	//logger.Println("")

	//log.Printf("\n")

	for a := 0; a < len(metrics.Functions); a++ {
		log.Println(metrics.Functions[a])
	}

	log.Printf("\nOutput:\n")
	for a := 0; a < len(output); a++ {
		log.Println(output[a])
	}
	log.Printf("\n")
	log.Printf("\n")
}
