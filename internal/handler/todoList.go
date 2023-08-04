package handler

import (
	"github.com/ZhanserikKalmukhambet/Trello/internal/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
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
