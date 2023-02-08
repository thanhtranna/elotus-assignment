package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"github.com/thanhtranna/elotus-assignment/hackathon/auth"
	"github.com/thanhtranna/elotus-assignment/hackathon/database"
	"github.com/thanhtranna/elotus-assignment/hackathon/models"
	"github.com/thanhtranna/elotus-assignment/hackathon/utils"
)

var (
	acceptImageMimes = map[string]struct{}{
		"image/png":  struct{}{},
		"image/jpeg": struct{}{},
		"image/gif":  struct{}{},
		"image/webp": struct{}{},
	}

	PathSaveImage = "./hackathon/tmp"
)

func RegisterUser(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	if err := user.HashPassword(user.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	record := database.Instance.Create(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusCreated, gin.H{"userId": user.ID, "email": user.Email, "username": user.Username, "created_at": user.CreatedAt})
}

func GenerateToken(context *gin.Context) {
	var request TokenRequest
	var user models.User
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	// check if email exists and password is correct
	record := database.Instance.Where("email = ?", request.Email).First(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		context.Abort()
		return
	}
	tokenString, err := auth.GenerateJWT(user.Email, user.Username)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func Ping(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func UploadFile(context *gin.Context) {
	// Source
	file, err := context.FormFile("data")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("get form err: %s", err.Error())})
		return
	}

	filename := filepath.Base(file.Filename)
	fileType, err := utils.GetFileType(file)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("open file err: %s", err.Error())})
		return
	}

	if _, ok := acceptImageMimes[fileType]; !ok {
		context.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("unsupported file type")})
		return
	}

	userEmail := context.GetString("user_email")
	var user models.User
	// check if email exists
	record := database.Instance.Where("email = ?", userEmail).First(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	newFileName := utils.AddPrefixFileName(filename)
	if err := context.SaveUploadedFile(file, fmt.Sprintf("%s/%s", PathSaveImage, newFileName)); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("upload file err: %s", err.Error())})
		return
	}

	image := models.Image{
		UserID:   uint64(user.ID),
		Name:     newFileName,
		Size:     uint64(file.Size),
		Metadata: getMetaHttpRequest(context),
	}

	record = database.Instance.Create(&image)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Successfully!", "filename": newFileName})
}

func getMetaHttpRequest(ctx *gin.Context) string {
	contentType := ctx.GetHeader("Content-Type")
	headers, _ := json.Marshal(ctx.Request.Header)

	meta := models.HttpRequestInfo{
		Method:      ctx.Request.Method,
		URL:         ctx.Request.URL.String(),
		ContentType: contentType,
		Proto:       ctx.Request.Proto,
		Header:      string(headers),
	}

	return meta.ToString()
}
