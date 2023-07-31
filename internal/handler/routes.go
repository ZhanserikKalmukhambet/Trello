package handler

import "github.com/gin-gonic/gin"

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getLists)
			lists.GET("/:id", h.getListByID)
			lists.PATCH("/:id", h.updateList)
			lists.DELETE("/:id", h.removeList)

			items := lists.Group("/:id/items")
			{
				{
					items.POST("/", h.createItem)
					items.GET("/", h.getItems)
				}
			}
		}

		items := api.Group("items")
		{
			items.GET("/:id", h.getItemByID)
			items.PUT("/:id", h.updateItem)
			items.DELETE("/:id", h.removeItem)
		}
	}

	return router
}
