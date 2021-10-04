package routes

import(
	handler "berbagi/controllers"
	middleware "berbagi/middlewares"
)

func productCartRoutes() {
	e.GET("/product-carts", handler.GetProductCartByUserIdController, middleware.AuthenticateUser, middleware.AuthorizeUser)

	e.PUT("/product-carts", handler.UpdateProductCartByUserIdController, middleware.AuthenticateUser, middleware.AuthorizeUser)

	e.DELETE("/product-carts", handler.DeleteProductCartByUserIdController, middleware.AuthenticateUser, middleware.AuthorizeUser)
}

