package handlers

import (
	"fmt"
	"github.com/duramash/go-halyklife-tt-2/types"
	"github.com/duramash/go-halyklife-tt-2/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMembers(ctx *gin.Context) {
	members := make([]types.Member, 0)
	DB.Find(&members)
	ctx.JSON(http.StatusOK, members)
}

func CreateMember(ctx *gin.Context) {
	member := &types.Member{}
	if err := ctx.BindJSON(&member); err != nil {
		util.HandleRequestError(ctx, http.StatusBadRequest, err)
		return
	}
	if err := DB.Create(member).Error; err != nil {
		util.HandleRequestError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusCreated, member)
}

func UpdateMember(ctx *gin.Context) {
	id := ctx.Param("id")
	member, updates := types.Member{}, types.MemberInterface{}
	if err := DB.First(&member, "id = ?", id).Error; err != nil {
		util.HandleRequestError(ctx, http.StatusBadRequest, err)
		return
	}
	if err := ctx.BindJSON(&updates); err != nil {
		util.HandleRequestError(ctx, http.StatusBadRequest, err)
		return
	}
	if err := DB.Model(&member).Updates(updates).Error; err != nil {
		util.HandleRequestError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, updates)
}

func DeleteMember(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := DB.Unscoped().Where("id = ?", id).Delete(&types.Member{}).Error; err != nil {
		util.HandleRequestError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, fmt.Sprintf("Deleted member with id: %s", id))
}

func GetBooksByMemberID(ctx *gin.Context) {
	memberId := ctx.Param("id")
	filteredBooks := make([]types.Book, 0)
	if err := DB.Where("member_id = ?", memberId).Find(&filteredBooks).Error; err != nil {
		util.HandleRequestError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, filteredBooks)
}
