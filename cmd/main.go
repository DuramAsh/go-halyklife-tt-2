package main

import (
	handlers2 "github.com/duramash/go-halyklife-tt-2/internal/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.Use(handlers2.CheckMethodType) // check for 405

	router.NoRoute(func(ctx *gin.Context) { // check for 404
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Page not found",
		})
	})

	authors := router.Group("/authors")
	{
		authors.GET("/", handlers2.GetAuthors)
		authors.POST("/", handlers2.CreateAuthor)
		authors.PUT("/:id", handlers2.UpdateAuthor)
		authors.DELETE("/:id", handlers2.DeleteAuthor)
		authors.GET("/:id/books", handlers2.GetBooksByAuthorID) // книги, у которых автор совпадает с нашим по айдишке, а также не забран никем
	}
	books := router.Group("/books")
	{
		books.GET("/", handlers2.GetBooks)
		books.POST("/", handlers2.CreateBook)
		books.PUT("/:id", handlers2.UpdateBook)
		books.DELETE("/:id", handlers2.DeleteBook)
	}
	members := router.Group("/members")
	{
		members.GET("/", handlers2.GetMembers)
		members.POST("/", handlers2.CreateMember)
		members.PUT("/:id", handlers2.UpdateMember)
		members.DELETE("/:id", handlers2.DeleteMember)
		members.GET("/:id/books", handlers2.GetBooksByMemberID) // книги, у которых еще нет borrower-а
	}
	_ = router.Run()
}
