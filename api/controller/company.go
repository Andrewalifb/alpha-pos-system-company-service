package controller

import (
	"net/http"
	"strconv"

	pb "github.com/Andrewalifb/alpha-pos-system-company-service/api/proto"
	"github.com/Andrewalifb/alpha-pos-system-company-service/dto"
	"github.com/Andrewalifb/alpha-pos-system-company-service/utils"

	"github.com/gin-gonic/gin"
)

type PosCompanyController interface {
	HandleCreatePosCompanyRequest(c *gin.Context)
	HandleReadPosCompanyRequest(c *gin.Context)
	HandleUpdatePosCompanyRequest(c *gin.Context)
	HandleDeletePosCompanyRequest(c *gin.Context)
	HandleReadAllPosCompaniesRequest(c *gin.Context)
}

type posCompanyController struct {
	service pb.PosCompanyServiceClient
}

func NewPosCompanyController(service pb.PosCompanyServiceClient) PosCompanyController {
	return &posCompanyController{
		service: service,
	}
}

func (c *posCompanyController) HandleCreatePosCompanyRequest(ctx *gin.Context) {
	// Declare req body Pos Company
	var req pb.CreatePosCompanyRequest

	// First, binding company data
	if err := ctx.ShouldBindJSON(&req.PosCompany); err != nil {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_COMPANY, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	// Get JWT Payload data from middleware
	getJwtPayload, exist := ctx.Get("user")
	if !exist {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_COMPANY, "Jwt Payload is Empty", nil)
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	// Add JWT payload from middleware into req body
	req.JwtPayload = getJwtPayload.(*pb.JWTPayload)

	// Service call
	res, err := c.service.CreatePosCompany(ctx, &req)
	if err != nil {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_COMPANY, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	// Success response
	successResponse := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATE_COMPANY, res)
	ctx.JSON(http.StatusOK, successResponse)
}

func (c *posCompanyController) HandleReadPosCompanyRequest(ctx *gin.Context) {
	var req pb.ReadPosCompanyRequest

	// Get user ID from URL
	companyID := ctx.Param("id")
	req.CompanyId = companyID

	// Get JWT Payload data from middleware
	getJwtPayload, exist := ctx.Get("user")
	if !exist {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_COMPANY, "Jwt Payload is Empty", nil)
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	// Add JWT payload from middleware into req body
	req.JwtPayload = getJwtPayload.(*pb.JWTPayload)

	res, err := c.service.ReadPosCompany(ctx, &req)
	if err != nil {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_COMPANY, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	successResponse := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_COMPANY, res)
	ctx.JSON(http.StatusOK, successResponse)
}

func (c *posCompanyController) HandleUpdatePosCompanyRequest(ctx *gin.Context) {
	var req pb.UpdatePosCompanyRequest

	// Get user ID from URL
	companyID := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&req.PosCompany); err != nil {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_COMPANY, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	req.PosCompany.CompanyId = companyID
	// Get JWT Payload data from middleware
	getJwtPayload, exist := ctx.Get("user")
	if !exist {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_COMPANY, "Jwt Payload is Empty", nil)
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	// Add JWT payload from middleware into req body
	req.JwtPayload = getJwtPayload.(*pb.JWTPayload)

	res, err := c.service.UpdatePosCompany(ctx, &req)
	if err != nil {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_COMPANY, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	successResponse := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_COMPANY, res)
	ctx.JSON(http.StatusOK, successResponse)
}

func (c *posCompanyController) HandleDeletePosCompanyRequest(ctx *gin.Context) {
	var req pb.DeletePosCompanyRequest

	// Get user ID from URL
	companyID := ctx.Param("id")
	req.CompanyId = companyID
	// Get JWT Payload data from middleware
	getJwtPayload, exist := ctx.Get("user")
	if !exist {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_COMPANY, "Jwt Payload is Empty", nil)
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	// Add JWT payload from middleware into req body
	req.JwtPayload = getJwtPayload.(*pb.JWTPayload)

	res, err := c.service.DeletePosCompany(ctx, &req)
	if err != nil {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_COMPANY, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	successResponse := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETE_COMPANY, res)
	ctx.JSON(http.StatusOK, successResponse)
}

func (c *posCompanyController) HandleReadAllPosCompaniesRequest(ctx *gin.Context) {
	limitQuery := ctx.Query("limit")
	pageQuery := ctx.Query("page")

	if (limitQuery == "" && pageQuery != "") || (limitQuery != "" && pageQuery == "") {
		errorResponse := utils.BuildResponseFailed("Both limit and page must be provided", "Value Is Empty", pb.CreatePosCompanyResponse{})
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	var req pb.ReadAllPosCompaniesRequest

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
			errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_COMPANY, "Jwt Payload is Empty", nil)
			ctx.JSON(http.StatusBadRequest, errorResponse)
			return
		}

		req = pb.ReadAllPosCompaniesRequest{
			Limit:      int32(limit),
			Page:       int32(page),
			JwtPayload: getJwtPayload.(*pb.JWTPayload),
		}
	}

	res, err := c.service.ReadAllPosCompanies(ctx, &req)
	if err != nil {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_COMPANY, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	successResponse := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_COMPANY, res)
	ctx.JSON(http.StatusOK, successResponse)
}
