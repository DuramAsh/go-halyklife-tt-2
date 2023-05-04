package main

import (
	"github.com/duramash/go-halyklife-tt-2/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.Use(handlers.CheckMethodType) // check for 405

	router.NoRoute(func(ctx *gin.Context) { // check for 404
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Page not found",
		})
	})

	authors := router.Group("/authors")
	{
		authors.GET("/", handlers.GetAuthors)
		authors.POST("/", handlers.CreateAuthor)
		authors.PUT("/:id", handlers.UpdateAuthor)
		authors.DELETE("/:id", handlers.DeleteAuthor)
		authors.GET("/:id/books", handlers.GetBooksByAuthorID) // книги, у которых автор совпадает с нашим по айдишке, а также не забран никем
	}
	books := router.Group("/books")
	{
		books.GET("/", handlers.GetBooks)
		books.POST("/", handlers.CreateBook)
		books.PUT("/:id", handlers.UpdateBook)
		books.DELETE("/:id", handlers.DeleteBook)
	}
	members := router.Group("/members")
	{
		members.GET("/", handlers.GetMembers)
		members.POST("/", handlers.CreateMember)
		members.PUT("/:id", handlers.UpdateMember)
		members.DELETE("/:id", handlers.DeleteMember)
		members.GET("/:id/books", handlers.GetBooksByMemberID) // книги, у которых еще нет borrower-а
	}
	_ = router.Run()
}
