package router

import (
	"server/internal/handler"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(
	userHandler *handler.UserHandler,
	authHandler *handler.AuthHandler,
	productHandler *handler.ProductHandler,
	categoryHandler *handler.CategoryHandler,
	linkHandler *handler.CategoryProductHandler,
	wsHandler *handler.WSHandler,
) {
	r = gin.Default()

	// Users
	r.GET("/users", userHandler.GetAllUsers)
	r.GET("cpf/:cpf", userHandler.GetUserByCpf)

	// Auth
	r.POST("signup", userHandler.CreateUser)
	r.POST("login", authHandler.Login)
	r.POST("logout", authHandler.Logout)

	// Chat
	r.POST("ws/create-room", wsHandler.CreateRoom)
	r.GET("ws/joinRoom/:roomID", wsHandler.JoinRoom)
	r.GET("ws/getRooms", wsHandler.GetRooms)
	r.GET("ws/getClientsInRoom/:roomID", wsHandler.GetClientsInRoom)

	// Products
	r.POST("products", productHandler.CreateProduct)
	r.GET("products", productHandler.GetAllProducts)
	r.GET("products/:id", productHandler.GetProduct)
	r.PUT("products/:id", productHandler.UpdateProduct)
	r.DELETE("products/:id", productHandler.DeleteProduct)
	r.POST("products/:id/categories", linkHandler.Attach)
	r.GET("products/:id/categories", linkHandler.ListCategoriesByProduct)
	r.DELETE("products/:id/categories/:categoryId", linkHandler.Detach)

	// Categories
	r.POST("categories", categoryHandler.CreateCategory)
	r.GET("categories", categoryHandler.GetAllCategorys)
	r.GET("categories/:id", categoryHandler.GetCategoryById)
	r.GET("categories/:id/products", linkHandler.ListProductsByCategory)

	// Users (por id — após rotas nomeadas para não conflitar com /:id)
	r.GET("/users/:id", userHandler.GetUserById)
	r.DELETE("/users/:id", userHandler.DeleteUserById)
	r.PUT("/users/:id", userHandler.UpdateUserById)
}

func Start(addr string) error {
	return r.Run(addr)
}
