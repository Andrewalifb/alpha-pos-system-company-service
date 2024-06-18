package controller

import (
	"net/http"
	"strconv"

	pb "github.com/Andrewalifb/alpha-pos-system-company-service/api/proto"
	"github.com/Andrewalifb/alpha-pos-system-company-service/dto"
	"github.com/Andrewalifb/alpha-pos-system-company-service/utils"

	"github.com/gin-gonic/gin"
)

type PosStoreBranchController interface {
	HandleCreatePosStoreBranchRequest(c *gin.Context)
	HandleReadPosStoreBranchRequest(c *gin.Context)
	HandleUpdatePosStoreBranchRequest(c *gin.Context)
	HandleDeletePosStoreBranchRequest(c *gin.Context)
	HandleReadAllPosStoreBranchesRequest(c *gin.Context)
}

type posStoreBranchController struct {
	service pb.PosStoreBranchServiceClient
}

func NewPosStoreBranchController(service pb.PosStoreBranchServiceClient) PosStoreBranchController {
	return &posStoreBranchController{
		service: service,
	}
}

func (c *posStoreBranchController) HandleCreatePosStoreBranchRequest(ctx *gin.Context) {
	// Declare req body Pos Store Branch
	var req pb.CreatePosStoreBranchRequest

	// First, binding store branch data
	if err := ctx.ShouldBindJSON(&req.PosStoreBranch); err != nil {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_BRANCH, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	// Get JWT Payload data from middleware
	getJwtPayload, exist := ctx.Get("user")
	if !exist {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_BRANCH, "Jwt Payload is Empty", nil)
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	// Add JWT payload from middleware into req body
	req.JwtPayload = getJwtPayload.(*pb.JWTPayload)

	// Service call
	res, err := c.service.CreatePosStoreBranch(ctx, &req)
	if err != nil {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_BRANCH, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	// Success response
	successResponse := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATE_BRANCH, res)
	ctx.JSON(http.StatusOK, successResponse)
}

func (c *posStoreBranchController) HandleReadPosStoreBranchRequest(ctx *gin.Context) {
	var req pb.ReadPosStoreBranchRequest

	// Get branch ID from URL
	branchID := ctx.Param("id")
	req.BranchId = branchID

	// Get JWT Payload data from middleware
	getJwtPayload, exist := ctx.Get("user")
	if !exist {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_BRANCH, "Jwt Payload is Empty", nil)
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	// Add JWT payload from middleware into req body
	req.JwtPayload = getJwtPayload.(*pb.JWTPayload)

	res, err := c.service.ReadPosStoreBranch(ctx, &req)
	if err != nil {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_BRANCH, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	successResponse := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_BRANCH, res)
	ctx.JSON(http.StatusOK, successResponse)
}

func (c *posStoreBranchController) HandleUpdatePosStoreBranchRequest(ctx *gin.Context) {
	var req pb.UpdatePosStoreBranchRequest

	// Get branch ID from URL
	branchID := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&req.PosStoreBranch); err != nil {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_BRANCH, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	// Get JWT Payload data from middleware
	getJwtPayload, exist := ctx.Get("user")
	if !exist {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_BRANCH, "Jwt Payload is Empty", nil)
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	// Set Branh ID from req
	req.PosStoreBranch.BranchId = branchID
	// Add JWT payload from middleware into req body
	req.JwtPayload = getJwtPayload.(*pb.JWTPayload)

	res, err := c.service.UpdatePosStoreBranch(ctx, &req)
	if err != nil {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_BRANCH, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	successResponse := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_BRANCH, res)
	ctx.JSON(http.StatusOK, successResponse)
}

func (c *posStoreBranchController) HandleDeletePosStoreBranchRequest(ctx *gin.Context) {
	var req pb.DeletePosStoreBranchRequest

	// Get branch ID from URL
	branchID := ctx.Param("id")
	req.BranchId = branchID

	// Get JWT Payload data from middleware
	getJwtPayload, exist := ctx.Get("user")
	if !exist {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_BRANCH, "Jwt Payload is Empty", nil)
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	// Add JWT payload from middleware into req body
	req.JwtPayload = getJwtPayload.(*pb.JWTPayload)

	res, err := c.service.DeletePosStoreBranch(ctx, &req)
	if err != nil {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_BRANCH, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	successResponse := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETE_BRANCH, res)
	ctx.JSON(http.StatusOK, successResponse)
}

func (c *posStoreBranchController) HandleReadAllPosStoreBranchesRequest(ctx *gin.Context) {
	limitQuery := ctx.Query("limit")
	pageQuery := ctx.Query("page")

	if (limitQuery == "" && pageQuery != "") || (limitQuery != "" && pageQuery == "") {
		errorResponse := utils.BuildResponseFailed("Both limit and page must be provided", "Value Is Empty", pb.CreatePosStoreBranchResponse{})
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	var req pb.ReadAllPosStoreBranchesRequest

	if limitQuery != "" && pageQuery != "" {
		limit, err := strconv.Atoi(limitQuery)
		if err != nil {
			errorResponse := utils.BuildResponseFailed("Invalid limit value", err.Error(), nil)
			ctx.JSON(http.StatusBadRequest, errorResponse)
			return
		}

		page, err := strconv.Atoi(pageQuery)
		if err != nil {
			errorResponse := utils.BuildResponseFailed("Invalid page value", err.Error(), nil)
			ctx.JSON(http.StatusBadRequest, errorResponse)
			return
		}

		// Get JWT Payload data from middleware
		getJwtPayload, exist := ctx.Get("user")
		if !exist {
			errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_BRANCH, "Jwt Payload is Empty", nil)
			ctx.JSON(http.StatusBadRequest, errorResponse)
			return
		}

		req = pb.ReadAllPosStoreBranchesRequest{
			Limit:      int32(limit),
			Page:       int32(page),
			JwtPayload: getJwtPayload.(*pb.JWTPayload),
		}
	}

	res, err := c.service.ReadAllPosStoreBranches(ctx, &req)
	if err != nil {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_BRANCH, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	successResponse := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_BRANCH, res)
	ctx.JSON(http.StatusOK, successResponse)
}
