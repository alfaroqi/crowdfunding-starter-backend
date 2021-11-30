package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	// capture input from user
	// map input from user to struct RegisterUserInput
	// struct above we pass it as a service parameter

	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValitationError(err)
		errorsMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorsMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.ApiResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// token, err := h.jwtService.GenerateToken(newUser)

	formatter := user.FormatUser(newUser, "tokentokentoken")

	response := helper.ApiResponse("Account has been registered", http.StatusOK, "sucess", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) LoginUser(c *gin.Context) {

	/* 	workflow :

	1. user memasukan input(email, password)
	2. input ditangkap handler
	3. mapping dari input user ke input struct
	5. input struct passing service
	6. di service mencari dengan bantuan repository user dan email x
	7. mencocokan password
	*/

	var input user.LoginUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValitationError(err)
		errorsMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Login failed", http.StatusUnprocessableEntity, "error", errorsMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedInUser, err := h.userService.LoginUser(input)

	if err != nil {
		errorsMessage := gin.H{"errors": err.Error()}

		response := helper.ApiResponse("Login failed", http.StatusUnprocessableEntity, "error", errorsMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := user.FormatUser(loggedInUser, "tokentokentoken")

	response := helper.ApiResponse("Logged successfully", http.StatusOK, "sucess", formatter)

	c.JSON(http.StatusOK, response)
}
