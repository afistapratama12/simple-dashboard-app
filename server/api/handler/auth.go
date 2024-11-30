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

// @Summary Login
// @Description Login User
// @ID login
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param request body request.LoginRequest true "Login Request"
// @Success 200 {object} wrapper.ResponseAPI{data=response.LoginResponse}
// @Failure 400 {object} wrapper.ResponseAPI{data=string}
// @Failure 500 {object} wrapper.ResponseAPI{data=string}
// @Router /v1/login [post]
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

// @Summary Register
// @Description Register User
// @ID register
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param request body request.RegisterRequest true "Register Request"
// @Success 200 {object} wrapper.ResponseAPI{data=string}
// @Failure 400 {object} wrapper.ResponseAPI{data=string}
// @Failure 500 {object} wrapper.ResponseAPI{data=string}
// @Router /v1/register [post]
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

// @Summary Verify Email
// @Description Verify Email
// @ID verify-email
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param request body request.VerifyEmailRequest true "Verify Email Request"
// @Success 200 {object} wrapper.ResponseAPI{data=string}
// @Failure 400 {object} wrapper.ResponseAPI{data=string}
// @Failure 500 {object} wrapper.ResponseAPI{data=string}
// @Router /v1/verify-email [post]
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

// @Summary Validate Token
// @Description Validate Token
// @ID validate-token
// @Tags Auth
// @Accept  json
// @Produce  json
// @Success 200 {object} wrapper.ResponseAPI{data=string}
// @Failure 401 {object} wrapper.ResponseAPI{data=string}
// @Security ApiKeyAuth
// @Router /v1/validate-token [get]
func (h *authHandler) ValidateToken(c *gin.Context) {
	w.ResponseJSON(c, http.StatusOK, "Token is valid")
}

// @Summary Notif Forgot Password
// @Description Notif Forgot Password
// @ID notif-forgot-password
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param request body request.ResetPasswordRequest true "Reset Password Request"
// @Success 200 {object} wrapper.ResponseAPI{data=string}
// @Failure 400 {object} wrapper.ResponseAPI{data=string}
// @Failure 500 {object} wrapper.ResponseAPI{data=string}
// @Router /v1/notif-forgot-password [post]
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

// @Summary Reset Password
// @Description Reset Password
// @ID reset-password
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param request body request.ResetPasswordConfirmRequest true "Reset Password Request"
// @Success 200 {object} wrapper.ResponseAPI{data=string}
// @Failure 400 {object} wrapper.ResponseAPI{data=string}
// @Failure 500 {object} wrapper.ResponseAPI{data=string}
// @Router /v1/reset-password [post]
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

// @Summary Edit User Login
// @Description Edit User Login
// @ID edit-user-login
// @Tags Users
// @Accept  json
// @Produce  json
// @Param request body request.EditUserRequest true "Edit User Request"
// @Success 200 {object} wrapper.ResponseAPI{data=string}
// @Failure 400 {object} wrapper.ResponseAPI{data=string}
// @Failure 401 {object} wrapper.ResponseAPI{data=string}
// @Failure 500 {object} wrapper.ResponseAPI{data=string}
// @Security ApiKeyAuth
// @Router /v1/users/edit [put]
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

// @Summary Profile User Login
// @Description Profile User Login
// @ID profile-user-login
// @Tags Users
// @Accept  json
// @Produce  json
// @Success 200 {object} wrapper.ResponseAPI{data=object}
// @Failure 401 {object} wrapper.ResponseAPI{data=string}
// @Failure 500 {object} wrapper.ResponseAPI{data=string}
// @Security ApiKeyAuth
// @Router /v1/users/profile [get]
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
