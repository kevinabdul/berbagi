package routes

import(
	handler "berbagi/controllers"
	middleware "berbagi/middlewares"
)

func productCartRoutes() {
	e.GET("/product-cart", handler.GetProductCartByUserIdController, middleware.AuthenticateUser)

	e.PUT("/product-cart", handler.UpdateProductCartByUserIdController, middleware.AuthenticateUser)

	e.DELETE("/product-cart", handler.DeleteProductCartByUserIdController, middleware.AuthenticateUser)
}

