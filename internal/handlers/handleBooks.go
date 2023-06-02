package handlers

import (
	"fmt"
	"github.com/duramash/go-halyklife-tt-2/internal/types"
	"github.com/duramash/go-halyklife-tt-2/internal/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetBooks(ctx *gin.Context) {
	books := make([]types.Book, 0)
	DB.Find(&books)
	ctx.JSON(http.StatusOK, books)
}

func CreateBook(ctx *gin.Context) {
	book := &types.Book{}
	if err := ctx.BindJSON(&book); err != nil {
		util.HandleRequestError(ctx, http.StatusBadRequest, err)
		return
	}
	if err := DB.Create(book).Error; err != nil {
		util.HandleRequestError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusCreated, book)
}

func UpdateBook(ctx *gin.Context) {
	id := ctx.Param("id")
	book, updates := types.Book{}, types.BookInterface{}
	if err := DB.First(&book, "id = ?", id).Error; err != nil {
		util.HandleRequestError(ctx, http.StatusBadRequest, err)
		return
	}
	if err := ctx.BindJSON(&updates); err != nil {
		util.HandleRequestError(ctx, http.StatusBadRequest, err)
		return
	}
	if err := DB.Model(&book).Updates(updates).Error; err != nil {
		util.HandleRequestError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, updates)
}

func DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := DB.Unscoped().Where("id = ?", id).Delete(&types.Book{}).Error; err != nil {
		util.HandleRequestError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, fmt.Sprintf("Deleted book with id: %s", id))
}
