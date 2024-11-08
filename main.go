package main

import (
    "diary_api/controller"
    "diary_api/database"
    "diary_api/middleware"
    "diary_api/model"
    "fmt"
    "log"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
)

func main() {
    loadEnv()
    loadDatabase()
    serveApplication()
}

func loadDatabase() {
    database.Connect()
    database.Database.AutoMigrate(&model.User{})
    database.Database.AutoMigrate(&model.Entry{})
}

func loadEnv() {
    err := godotenv.Load(".env.local")
    if err != nil {
        log.Fatal("Error loading .env file")
    }
}

func serveApplication() {
    router := gin.Default()

    publicRoutes := router.Group("/auth")
    publicRoutes.POST("/register", controller.Register)
    publicRoutes.POST("/login", controller.Login)

    protectedRoutes := router.Group("/api")
    protectedRoutes.Use(middleware.JWTAuthMiddleware())
    protectedRoutes.POST("/entry", controller.AddEntry)
    protectedRoutes.GET("/entries", controller.GetAllEntries)
    protectedRoutes.GET("/entry/:entry_id", controller.GetEntry)
    // protectedRoutes.PUT("/entry/:entry_id", controller.UpdateEntry)

    router.Run(":8000")
    fmt.Println("Server running on port 8000")
}