package controller

import (
    "diary_api/helper"
    "diary_api/model"
    "github.com/gin-gonic/gin"
    "net/http"
    "fmt"
    "diary_api/database"
    // "encoding/json"
)

func AddEntry(context *gin.Context) {
    var input model.Entry
    if err := context.ShouldBindJSON(&input); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := helper.CurrentUser(context)

    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    input.UserID = user.ID

    savedEntry, err := input.Save()

    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    context.JSON(http.StatusCreated, gin.H{"data": savedEntry})
}

func GetAllEntries(context *gin.Context) {
    user, err := helper.CurrentUser(context)

    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    fmt.Println("user: ", user)
    fmt.Println("user.Entries: ", user.Entries)
    context.JSON(http.StatusOK, gin.H{"data": user.Entries})
}

func GetEntry(context *gin.Context) {
    user, err := helper.CurrentUser(context)

    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var input model.Entry
    result := database.Database.Where("user_id = ?", user.ID).First(&input, context.Param("entry_id"))
    if result.Error != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": "Not Found"})
        return
    }
    context.JSON(http.StatusOK, gin.H{"data": input})
}
