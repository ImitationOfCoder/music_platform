package auth_handler

import (
	"errors"
	"net/http"

	"auth_microservice/internal/domain"
	"auth_microservice/internal/handler/auth/dto"

	"github.com/ImitationOfCoder/music_platform/pkg/http_server"
	"github.com/ImitationOfCoder/music_platform/pkg/logger"
)

func (h *AuthHandler) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := http_server.NewHTTPResponseHandler(w, log)

	sessionTokenCookie, err := r.Cookie("session_token")
	if err != nil {
		responseHandler.UnauthenticatedResponse()
		return
	}

	user, err := h.authService.GetCurrentUser(ctx, sessionTokenCookie.Value)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			responseHandler.JSONResponse(http_server.HTTPResponse{
				Success: false,
				Code:    http.StatusUnauthorized,
				Message: "Invalid credentials.",
				Error:   err,
			})
			return
		}

		responseHandler.InternalErrorResponse(err)
		return
	}

	responseHandler.JSONResponse(http_server.HTTPResponse{
		Success: true,
		Code:    http.StatusOK,
		Data: map[string]any{
			"user": dto.UserResponse{
				Id:    user.Id,
				Name:  user.Name,
				Email: user.Email,
			},
		},
	})
}
