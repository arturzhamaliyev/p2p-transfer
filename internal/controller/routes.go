package controller

func (h *Handler) InitRoutes() {
	transfer := h.Router.Group("/transfer")
	{
		transfer.POST("/", h.MakeTransfer)
	}

	user := h.Router.Group("/user")
	{
		user.POST("/", h.CreateUser)
	}
}
