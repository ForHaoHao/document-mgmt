package main

import (
	"DocumentMgmtSystem/library"
	"DocumentMgmtSystem/routers"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	router := routers.NewRouter()

	godotenv.Load()

	port := os.Getenv("HTTP_PORT")
	domain := os.Getenv("HTTP_DOMAIN")
	portocol := os.Getenv("HTTP_PORTOCOL")

	library.OpenBrowser(fmt.Sprintf("%s://%s:%s", portocol, domain, port))

	router.Run(fmt.Sprintf("%s:%s", domain, port))
}
