package handler

import "github.com/gin-gonic/gin"

func (h *Handler) InitRouter() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	api.Use(h.authMiddleware())
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getLists)
			lists.GET("/:id", h.getListByID)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.removeList)

			items := lists.Group("/:id/items")
			{
				{
					items.POST("/", h.createListItem)
					items.GET("/", h.getListItems)
					items.PUT("/:item_id", h.updateListItem)
					items.DELETE("/:item_id", h.removeListItem)
				}
			}
		}
	}

	return router
}
