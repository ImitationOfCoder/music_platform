package auth_handler

import (
	"errors"
	"net/http"
	"time"

	"auth_microservice/internal/domain"
	"auth_microservice/internal/handler/auth/dto"

	"github.com/ImitationOfCoder/music_platform/pkg/http_server"
	"github.com/ImitationOfCoder/music_platform/pkg/logger"
)

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := http_server.NewHTTPResponseHandler(w, log)
	request := http_server.NewHttpRequest(r)

	var body dto.LoginUserRequest
	if err := request.DecodeAndValidate(&body, body.GetCustomMessagesForValidator()); err != nil {
		responseHandler.InvalidRequestData(err)
		return
	}

	token, err := h.authService.Login(
		ctx,
		body.Email,
		body.Password,
	)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidCredentials) {
			responseHandler.JSONResponse(http_server.HTTPResponse{
				Success: false,
				Code:    http.StatusUnauthorized,
				Message: "Invalid credentials.",
			})
			return
		}

		responseHandler.InternalErrorResponse(err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   int(24 * time.Hour * 30), // 30 дней в секундах
	})

	responseHandler.SuccessResponse(true, http.StatusOK)
}
