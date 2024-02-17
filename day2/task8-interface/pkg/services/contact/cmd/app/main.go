package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"

    "go/pkg/services/contact/internal"
    "go/pkg/services/contact/internal/delivery"
    "go/pkg/services/contact/internal/repository"
    "go/pkg/services/contact/internal/usecase"

    _ "github.com/joho/godotenv/autoload"
)

func main() {
    contactRepo := internal.NewContactRepository()
    groupRepo := internal.NewGroupRepository()

    contactUseCase := internal.NewContactUseCase(contactRepo)
    groupUseCase := internal.NewGroupUseCase(groupRepo)

    contactHandler := internal.NewContactHandler(contactUseCase)
    groupHandler := internal.NewGroupHandler(groupUseCase)

    http.HandleFunc("/contacts", contactHandler.HandleHTTP)
    http.HandleFunc("/groups", groupHandler.HandleHTTP)

    go func() {
        if err := http.ListenAndServe(":8080", nil); err != nil {
            log.Fatal("HTTP server error: ", err)
        }
    }()

    fmt.Println("Server started on port 8080")

    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit

    fmt.Println("Server shutting down...")
}
