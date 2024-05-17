package controller

import (
	"fmt"
	"net/http"
	"strconv"

	pb "github.com/Andrewalifb/alpha-pos-system-company-service/api/proto"
	"github.com/Andrewalifb/alpha-pos-system-company-service/dto"
	"github.com/Andrewalifb/alpha-pos-system-company-service/utils"

	"github.com/gin-gonic/gin"
)

type PosUserController interface {
	HandleCreatePosUserRequest(c *gin.Context)
	HandleReadPosUserRequest(c *gin.Context)
	HandleUpdatePosUserRequest(c *gin.Context)
	HandleDeletePosUserRequest(c *gin.Context)
	HandleReadAllPosUsersRequest(c *gin.Context)
	HandleLoginRequest(c *gin.Context)
}

type posUserController struct {
	service pb.PosUserServiceClient
}

func NewPosUserController(service pb.PosUserServiceClient) PosUserController {
	return &posUserController{
		service: service,
	}
}

func (c *posUserController) HandleCreatePosUserRequest(ctx *gin.Context) {
	// Declare req body Pos User
	var req pb.CreatePosUserRequest

	// First, binding user data
	if err := ctx.ShouldBindJSON(&req.PosUser); err != nil {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_USER, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	// Get JWT Payload data from middleware
	getJwtPayload, exist := ctx.Get("user")
	if !exist {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_USER, "Jwt Payload is Empty", nil)
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	// Add JWT payload from middleware into req body
	req.JwtPayload = getJwtPayload.(*pb.JWTPayload)

	res, err := c.service.CreatePosUser(ctx, &req)
	if err != nil {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_USER, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	successResponse := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATE_USER, res)
	ctx.JSON(http.StatusOK, successResponse)
}

func (c *posUserController) HandleLoginRequest(ctx *gin.Context) {
	var req pb.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		errorResponse := utils.BuildResponseFailed("Failed to login", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	if req.Username == "" || req.Password == "" {
		errorResponse := utils.BuildResponseFailed("Failed to login", "Username and password must not be empty", nil)
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	res, err := c.service.Login(ctx, &req)
	if err != nil {
		errorResponse := utils.BuildResponseFailed("Failed to login", "Invalid username or password", nil)
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	successResponse := utils.BuildResponseSuccess("Login successful", res)
	ctx.JSON(http.StatusOK, successResponse)
}

func (c *posUserController) HandleReadAllPosUsersRequest(ctx *gin.Context) {
	limitQuery := ctx.Query("limit")
	pageQuery := ctx.Query("page")

	if (limitQuery == "" && pageQuery != "") || (limitQuery != "" && pageQuery == "") {
		errorResponse := utils.BuildResponseFailed("Both limit and page must be provided", "Value Is Empty", pb.CreatePosUserResponse{})
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	var req pb.ReadAllPosUsersRequest

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
			errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_USER, "Jwt Payload is Empty", nil)
			ctx.JSON(http.StatusBadRequest, errorResponse)
			return
		}

		req = pb.ReadAllPosUsersRequest{
			Limit:      int32(limit),
			Page:       int32(page),
			JwtPayload: getJwtPayload.(*pb.JWTPayload),
		}
	}

	res, err := c.service.ReadAllPosUsers(ctx, &req)
	if err != nil {
		errorResponse := utils.BuildResponseFailed("Failed to get users", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	successResponse := utils.BuildResponseSuccess("Successfully fetched users", res)
	ctx.JSON(http.StatusOK, successResponse)
}
func (c *posUserController) HandleReadPosUserRequest(ctx *gin.Context) {
	var req pb.ReadPosUserRequest

	// Get user ID from URL
	userID := ctx.Param("id")
	req.UserId = userID

	// Get JWT Payload data from middleware
	getJwtPayload, exist := ctx.Get("user")
	if !exist {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_USER, "Jwt Payload is Empty", nil)
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	// Add JWT payload from middleware into req body
	req.JwtPayload = getJwtPayload.(*pb.JWTPayload)

	res, err := c.service.ReadPosUser(ctx, &req)
	if err != nil {
		errorResponse := utils.BuildResponseFailed("Failed to get user", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	successResponse := utils.BuildResponseSuccess("Successfully fetched user", res)
	ctx.JSON(http.StatusOK, successResponse)
}

func (c *posUserController) HandleUpdatePosUserRequest(ctx *gin.Context) {
	var req pb.UpdatePosUserRequest

	// Get user ID from URL
	userID := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&req.PosUser); err != nil {
		errorResponse := utils.BuildResponseFailed("Failed to update user", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	req.PosUser.UserId = userID

	// Get JWT Payload data from middleware
	getJwtPayload, exist := ctx.Get("user")
	if !exist {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_USER, "Jwt Payload is Empty", nil)
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	// Add JWT payload from middleware into req body
	req.JwtPayload = getJwtPayload.(*pb.JWTPayload)

	res, err := c.service.UpdatePosUser(ctx, &req)
	if err != nil {
		errorResponse := utils.BuildResponseFailed("Failed to update user", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	successResponse := utils.BuildResponseSuccess("Successfully updated user", res)
	ctx.JSON(http.StatusOK, successResponse)
}

func (c *posUserController) HandleDeletePosUserRequest(ctx *gin.Context) {
	var req pb.DeletePosUserRequest

	// Get user ID from URL
	userID := ctx.Param("id")
	req.UserId = userID

	// Get JWT Payload data from middleware
	getJwtPayload, exist := ctx.Get("user")
	if !exist {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_USER, "Jwt Payload is Empty", nil)
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	// Add JWT payload from middleware into req body
	req.JwtPayload = getJwtPayload.(*pb.JWTPayload)
	fmt.Println("JWT payloads :", req.JwtPayload)

	res, err := c.service.DeletePosUser(ctx, &req)
	if err != nil {
		errorResponse := utils.BuildResponseFailed("Failed to delete user", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	successResponse := utils.BuildResponseSuccess("Successfully deleted user", res)
	ctx.JSON(http.StatusOK, successResponse)
}
