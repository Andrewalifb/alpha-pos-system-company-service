package controller

import (
	"net/http"
	"strconv"

	pb "github.com/Andrewalifb/alpha-pos-system-company-service/api/proto"
	"github.com/Andrewalifb/alpha-pos-system-company-service/dto"
	"github.com/Andrewalifb/alpha-pos-system-company-service/utils"

	"github.com/gin-gonic/gin"
)

type PosStoreController interface {
	HandleCreatePosStoreRequest(c *gin.Context)
	HandleReadPosStoreRequest(c *gin.Context)
	HandleUpdatePosStoreRequest(c *gin.Context)
	HandleDeletePosStoreRequest(c *gin.Context)
	HandleReadAllPosStoresRequest(c *gin.Context)
}

type posStoreController struct {
	service pb.PosStoreServiceClient
}

func NewPosStoreController(service pb.PosStoreServiceClient) PosStoreController {
	return &posStoreController{
		service: service,
	}
}

func (c *posStoreController) HandleCreatePosStoreRequest(ctx *gin.Context) {
	// Declare req body Pos Store
	var req pb.CreatePosStoreRequest

	// First, binding store data
	if err := ctx.ShouldBindJSON(&req.PosStore); err != nil {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_STORE, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	// Get JWT Payload data from middleware
	getJwtPayload, exist := ctx.Get("user")
	if !exist {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_STORE, "Jwt Payload is Empty", nil)
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	// Add JWT payload from middleware into req body
	req.JwtPayload = getJwtPayload.(*pb.JWTPayload)

	// Service call
	res, err := c.service.CreatePosStore(ctx, &req)
	if err != nil {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_STORE, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	// Success response
	successResponse := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATE_STORE, res)
	ctx.JSON(http.StatusOK, successResponse)
}

func (c *posStoreController) HandleReadPosStoreRequest(ctx *gin.Context) {
	var req pb.ReadPosStoreRequest

	// Get store ID from URL
	storeID := ctx.Param("id")
	req.StoreId = storeID

	// Get JWT Payload data from middleware
	getJwtPayload, exist := ctx.Get("user")
	if !exist {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_STORE, "Jwt Payload is Empty", nil)
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	// Add JWT payload from middleware into req body
	req.JwtPayload = getJwtPayload.(*pb.JWTPayload)

	res, err := c.service.ReadPosStore(ctx, &req)
	if err != nil {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_STORE, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	successResponse := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_STORE, res)
	ctx.JSON(http.StatusOK, successResponse)
}

func (c *posStoreController) HandleUpdatePosStoreRequest(ctx *gin.Context) {
	var req pb.UpdatePosStoreRequest
	if err := ctx.ShouldBindJSON(&req.PosStore); err != nil {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_STORE, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	// Get JWT Payload data from middleware
	getJwtPayload, exist := ctx.Get("user")
	if !exist {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_STORE, "Jwt Payload is Empty", nil)
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	// Add JWT payload from middleware into req body
	req.JwtPayload = getJwtPayload.(*pb.JWTPayload)

	res, err := c.service.UpdatePosStore(ctx, &req)
	if err != nil {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_STORE, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	successResponse := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_STORE, res)
	ctx.JSON(http.StatusOK, successResponse)
}
func (c *posStoreController) HandleDeletePosStoreRequest(ctx *gin.Context) {
	var req pb.DeletePosStoreRequest

	// Get store ID from URL
	storeID := ctx.Param("id")
	req.StoreId = storeID

	// Get JWT Payload data from middleware
	getJwtPayload, exist := ctx.Get("user")
	if !exist {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_STORE, "Jwt Payload is Empty", nil)
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	// Add JWT payload from middleware into req body
	req.JwtPayload = getJwtPayload.(*pb.JWTPayload)

	res, err := c.service.DeletePosStore(ctx, &req)
	if err != nil {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_STORE, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	successResponse := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETE_STORE, res)
	ctx.JSON(http.StatusOK, successResponse)
}

func (c *posStoreController) HandleReadAllPosStoresRequest(ctx *gin.Context) {
	limitQuery := ctx.Query("limit")
	pageQuery := ctx.Query("page")

	if (limitQuery == "" && pageQuery != "") || (limitQuery != "" && pageQuery == "") {
		errorResponse := utils.BuildResponseFailed("Both limit and page must be provided", "Value Is Empty", pb.CreatePosStoreResponse{})
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	var req pb.ReadAllPosStoresRequest

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
			errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_STORE, "Jwt Payload is Empty", nil)
			ctx.JSON(http.StatusBadRequest, errorResponse)
			return
		}

		req = pb.ReadAllPosStoresRequest{
			Limit:      int32(limit),
			Page:       int32(page),
			JwtPayload: getJwtPayload.(*pb.JWTPayload),
		}
	}

	res, err := c.service.ReadAllPosStores(ctx, &req)
	if err != nil {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_STORE, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	successResponse := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_STORE, res)
	ctx.JSON(http.StatusOK, successResponse)
}
