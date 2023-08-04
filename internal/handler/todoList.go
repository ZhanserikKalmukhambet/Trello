package handler

import (
	"github.com/ZhanserikKalmukhambet/Trello/internal/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary Create todo list
// @Security ApiKeyAuth
// @Tags lists
// @Description create todo list
// @ID create-list
// @Accept  json
// @Produce  json
// @Param input body entity.TodoList true "list info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} Error
// @Failure 500 {object} Error
// @Failure default {object} Error
// @Router /api/lists [post]
func (h *Handler) createList(ctx *gin.Context) {
	userID, err := getAuthUserID(ctx)
	if err != nil {
		return
	}

	var input entity.TodoList
	if err := ctx.ShouldBindJSON(&input); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.srvs.CreateTodoList(userID, input)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"todo_list_id": id})
}

// @Summary Get All Lists
// @Security ApiKeyAuth
// @Tags lists
// @Description get all lists
// @ID get-all-lists
// @Accept  json
// @Produce  json
// @Success 200 {object} StatusResponse
// @Failure 400,404 {object} Error
// @Failure 500 {object} Error
// @Failure default {object} Error
// @Router /api/lists [get]
func (h *Handler) getLists(ctx *gin.Context) {
	userID, err := getAuthUserID(ctx)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	lists, err := h.srvs.TodoList.GetTodoLists(userID)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, lists)
}

// @Summary Get List By ID
// @Security ApiKeyAuth
// @Tags lists
// @Description get list by id
// @ID get-list-by-id
// @Accept  json
// @Produce  json
// @Success 200 {object} entity.TodoList
// @Failure 400,404 {object} Error
// @Failure 500 {object} Error
// @Failure default {object} Error
// @Router /api/lists/:id [get]
func (h *Handler) getListByID(ctx *gin.Context) {
	userID, err := getAuthUserID(ctx)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "invalid id param")
		return
	}

	list, err := h.srvs.TodoList.GetTodoListByID(userID, id)
	if err != nil || list.ID == 0 {
		NewErrorResponse(ctx, http.StatusInternalServerError, "You can't get other's lists")
		return
	}

	ctx.JSON(http.StatusOK, list)
}

func (h *Handler) updateList(ctx *gin.Context) {
	userID, err := getAuthUserID(ctx)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "invalid id param")
		return
	}

	var input entity.UpdateListInput
	if err := ctx.BindJSON(&input); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.srvs.TodoList.UpdateTodoList(userID, id, input); err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func (h *Handler) removeList(ctx *gin.Context) {
	userID, err := getAuthUserID(ctx)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.srvs.TodoList.DeleteTodoList(userID, id)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
