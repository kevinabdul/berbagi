package routes

import(
	handler "berbagi/controllers"
)

func productRoutes() {
	e.GET("/products", handler.GetProductsController)
}

