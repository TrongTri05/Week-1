package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"ShortLink/controller"
	"ShortLink/db"
	"ShortLink/repository"
	"ShortLink/service"
)

func main() {
	database, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer database.Close()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter URL: ")
	url, _ := reader.ReadString('\n')
	url = strings.TrimSpace(url)

	repo := repository.NewLinkRepository(database)
	svc := service.NewLinkService(repo)
	ctrl := controller.NewLinkController(svc)

	shortCode, err := svc.CreateShortLink(url)
	if err != nil {
		panic(err)
	}

	fmt.Println("Short URL:")
	fmt.Println("http://localhost:8080/" + shortCode)
	http.HandleFunc("/", ctrl.Redirect)
	http.ListenAndServe(":8080", nil)
}
