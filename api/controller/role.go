package controller

import (
	"net/http"
	"strconv"

	pb "github.com/Andrewalifb/alpha-pos-system-company-service/api/proto"
	"github.com/Andrewalifb/alpha-pos-system-company-service/dto"
	"github.com/Andrewalifb/alpha-pos-system-company-service/utils"

	"github.com/gin-gonic/gin"
)

type PosRoleController interface {
	HandleCreatePosRoleRequest(c *gin.Context)
	HandleReadPosRoleRequest(c *gin.Context)
	HandleUpdatePosRoleRequest(c *gin.Context)
	HandleDeletePosRoleRequest(c *gin.Context)
	HandleReadAllPosRolesRequest(c *gin.Context)
}

type posRoleController struct {
	service pb.PosRoleServiceClient
}

func NewPosRoleController(service pb.PosRoleServiceClient) PosRoleController {
	return &posRoleController{
		service: service,
	}
}

func (c *posRoleController) HandleCreatePosRoleRequest(ctx *gin.Context) {
	// Declare req body Pos Role
	var req pb.CreatePosRoleRequest

	// First, binding role data
	if err := ctx.ShouldBindJSON(&req.PosRole); err != nil {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_ROLE, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	// Get JWT Payload data from middleware
	getJwtPayload, exist := ctx.Get("user")
	if !exist {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_ROLE, "Jwt Payload is Empty", nil)
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	// Add JWT payload from middleware into req body
	req.JwtPayload = getJwtPayload.(*pb.JWTPayload)

	// Service call
	res, err := c.service.CreatePosRole(ctx, &req)
	if err != nil {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_ROLE, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	// Success response
	successResponse := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATE_ROLE, res)
	ctx.JSON(http.StatusOK, successResponse)
}

func (c *posRoleController) HandleReadPosRoleRequest(ctx *gin.Context) {
	var req pb.ReadPosRoleRequest

	// Get role ID from URL
	roleID := ctx.Param("id")
	req.RoleId = roleID

	// Get JWT Payload data from middleware
	getJwtPayload, exist := ctx.Get("user")
	if !exist {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_ROLE, "Jwt Payload is Empty", nil)
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	// Add JWT payload from middleware into req body
	req.JwtPayload = getJwtPayload.(*pb.JWTPayload)

	res, err := c.service.ReadPosRole(ctx, &req)
	if err != nil {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_ROLE, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	successResponse := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_ROLE, res)
	ctx.JSON(http.StatusOK, successResponse)
}

func (c *posRoleController) HandleUpdatePosRoleRequest(ctx *gin.Context) {
	var req pb.UpdatePosRoleRequest
	if err := ctx.ShouldBindJSON(&req.PosRole); err != nil {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_ROLE, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	// Get JWT Payload data from middleware
	getJwtPayload, exist := ctx.Get("user")
	if !exist {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_ROLE, "Jwt Payload is Empty", nil)
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	// Add JWT payload from middleware into req body
	req.JwtPayload = getJwtPayload.(*pb.JWTPayload)

	res, err := c.service.UpdatePosRole(ctx, &req)
	if err != nil {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_ROLE, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	successResponse := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_ROLE, res)
	ctx.JSON(http.StatusOK, successResponse)
}

func (c *posRoleController) HandleDeletePosRoleRequest(ctx *gin.Context) {
	var req pb.DeletePosRoleRequest

	// Get role ID from URL
	roleID := ctx.Param("id")
	req.RoleId = roleID

	// Get JWT Payload data from middleware
	getJwtPayload, exist := ctx.Get("user")
	if !exist {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_ROLE, "Jwt Payload is Empty", nil)
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	// Add JWT payload from middleware into req body
	req.JwtPayload = getJwtPayload.(*pb.JWTPayload)

	res, err := c.service.DeletePosRole(ctx, &req)
	if err != nil {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_ROLE, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	successResponse := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETE_ROLE, res)
	ctx.JSON(http.StatusOK, successResponse)
}

func (c *posRoleController) HandleReadAllPosRolesRequest(ctx *gin.Context) {
	limitQuery := ctx.Query("limit")
	pageQuery := ctx.Query("page")

	if (limitQuery == "" && pageQuery != "") || (limitQuery != "" && pageQuery == "") {
		errorResponse := utils.BuildResponseFailed("Both limit and page must be provided", "Value Is Empty", pb.CreatePosRoleResponse{})
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	var req pb.ReadAllPosRolesRequest

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
			errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_ROLE, "Jwt Payload is Empty", nil)
			ctx.JSON(http.StatusBadRequest, errorResponse)
			return
		}

		req = pb.ReadAllPosRolesRequest{
			Limit:      int32(limit),
			Page:       int32(page),
			JwtPayload: getJwtPayload.(*pb.JWTPayload),
		}
	}

	res, err := c.service.ReadAllPosRoles(ctx, &req)
	if err != nil {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_ROLE, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	successResponse := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_ROLE, res)
	ctx.JSON(http.StatusOK, successResponse)
}
