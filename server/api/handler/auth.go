package handler

import (
	"net/http"
	"simple-dashboard-server/api/request"
	"simple-dashboard-server/service"
	w "simple-dashboard-server/wrapper"

	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
	NotifForgotPassword(c *gin.Context)
	VerifyEmail(c *gin.Context)
	ValidateToken(c *gin.Context)
	ResetPassword(c *gin.Context)

	EditUserLogin(c *gin.Context)
	ProfileUserLogin(c *gin.Context)
}

type authHandler struct {
	authService service.AuthService
	userService service.UserService
}

func NewAuthHandler(authSvc service.AuthService, userService service.UserService) AuthHandler {
	return &authHandler{
		authService: authSvc,
		userService: userService,
	}
}

func (h *authHandler) Login(c *gin.Context) {
	var request request.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		w.ResponseError(c, http.StatusBadRequest, err)
		return
	}

	resp, err := h.authService.Login(c, request)
	if err != nil {
		w.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	w.ResponseJSON(c, http.StatusOK, resp)
}

func (h *authHandler) Register(c *gin.Context) {
	var request request.RegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		w.ResponseError(c, http.StatusBadRequest, err)
		return
	}

	err := h.authService.Register(c, request)
	if err != nil {
		w.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	w.ResponseJSONWithMessage(c, http.StatusOK, "Register success")
}

func (h *authHandler) VerifyEmail(c *gin.Context) {
	var request request.VerifyEmailRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		w.ResponseError(c, http.StatusBadRequest, err)
		return
	}

	resp, err := h.authService.VerifyEmail(c, request)
	if err != nil {
		w.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	w.ResponseJSON(c, http.StatusOK, resp)
}

func (h *authHandler) ValidateToken(c *gin.Context) {
	w.ResponseJSON(c, http.StatusOK, "Token is valid")
}

func (h *authHandler) NotifForgotPassword(c *gin.Context) {
	var request request.ResetPasswordRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		w.ResponseError(c, http.StatusBadRequest, err)
		return
	}

	err := h.authService.NotifForgotPassword(c, request)
	if err != nil {
		w.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	w.ResponseJSONWithMessage(c, http.StatusOK, "Reset password success")
}

func (h *authHandler) ResetPassword(c *gin.Context) {
	var request request.ResetPasswordConfirmRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		w.ResponseError(c, http.StatusBadRequest, err)
		return
	}

	err := h.authService.ResetPassword(c, request)
	if err != nil {
		w.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	w.ResponseJSONWithMessage(c, http.StatusOK, "Reset password success")
}

func (h *authHandler) EditUserLogin(c *gin.Context) {
	var request request.EditUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		w.ResponseError(c, http.StatusBadRequest, err)
		return
	}

	userId, ok := c.Get("user_id")
	if !ok {
		w.ResponseJSONWithMessage(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	// get param id
	request.ID = userId.(string)
	if request.ID == "" {
		w.ResponseJSONWithMessage(c, http.StatusBadRequest, "ID is required")
		return
	}

	err := h.userService.EditUserLogin(c, request)
	if err != nil {
		w.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	w.ResponseJSONWithMessage(c, http.StatusOK, "Edit user success")
}

func (h *authHandler) ProfileUserLogin(c *gin.Context) {
	userId, ok := c.Get("user_id")
	if !ok {
		w.ResponseJSONWithMessage(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	resp, err := h.userService.GetProfileUserLogin(c, userId.(string))
	if err != nil {
		w.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	w.ResponseJSON(c, http.StatusOK, resp)
}
