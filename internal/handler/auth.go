package handler

import (
	"github.com/ZhanserikKalmukhambet/Trello/internal/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param newUser body entity.RegisterRequest true "account info"
// @Success 201 {integer} integer 1
// @Failure 400,404 {object} Error
// @Failure 500 {object} Error
// @Failure default {object} Error
// @Router /auth/sign-up [post]
func (h *Handler) signUp(ctx *gin.Context) {
	var newUser entity.RegisterRequest

	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user := entity.User{
		Username: newUser.Username,
		Email:    newUser.Email,
		Password: newUser.Password,
	}
	id, err := h.srvs.Register(user)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"user_id": id})
}

// @Summary SignIn
// @Tags auth
// @Description login
// @ID login
// @Accept  json
// @Produce  json
// @Param input body entity.LoginRequest true "credentials"
// @Success 200 {string} string "token"
// @Failure 400,404 {object} Error
// @Failure 500 {object} Error
// @Failure default {object} Error
// @Router /auth/sign-in [post]
func (h *Handler) signIn(ctx *gin.Context) {
	var input entity.LoginRequest

	if err := ctx.ShouldBindJSON(&input); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.srvs.GenerateToken(input.Username, input.Password)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"token": token})
}
