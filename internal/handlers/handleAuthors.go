package handlers

import (
	"fmt"
	"github.com/duramash/go-halyklife-tt-2/internal/db"
	"github.com/duramash/go-halyklife-tt-2/internal/types"
	"github.com/duramash/go-halyklife-tt-2/internal/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

var DB = db.GetDB()

func GetAuthors(ctx *gin.Context) {
	authors := make([]types.Author, 0)
	DB.Find(&authors)
	ctx.JSON(http.StatusOK, authors)
}

func CreateAuthor(ctx *gin.Context) {
	author := &types.Author{}
	if err := ctx.BindJSON(&author); err != nil {
		util.HandleRequestError(ctx, http.StatusBadRequest, err)
		return
	}
	if err := DB.Create(author).Error; err != nil {
		util.HandleRequestError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusCreated, author)
}

func UpdateAuthor(ctx *gin.Context) {
	id := ctx.Param("id")
	author, updates := types.Author{}, types.AuthorInterface{}
	if err := DB.First(&author, "id = ?", id).Error; err != nil {
		util.HandleRequestError(ctx, http.StatusBadRequest, err)
		return
	}
	if err := ctx.BindJSON(&updates); err != nil {
		util.HandleRequestError(ctx, http.StatusBadRequest, err)
		return
	}
	if err := DB.Model(&author).Updates(updates).Error; err != nil {
		util.HandleRequestError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, updates)
}

func DeleteAuthor(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := DB.Unscoped().Where("id = ?", id).Delete(&types.Author{}).Error; err != nil {
		util.HandleRequestError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, fmt.Sprintf("Deleted author with id: %s", id))
}

func GetBooksByAuthorID(ctx *gin.Context) {
	id := ctx.Param("id")
	filteredBooks := make([]types.Book, 0)
	if err := DB.Where("member_id IS NULL AND author_id = ?", id).Find(&filteredBooks).Error; err != nil {
		util.HandleRequestError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, filteredBooks)
}
