package auth_handler

import (
	"net/http"

	"github.com/ImitationOfCoder/music_platform/pkg/http_server"
	"github.com/ImitationOfCoder/music_platform/pkg/logger"
)

func (h *AuthHandler) LogOutUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := http_server.NewHTTPResponseHandler(w, log)

	sessionTokenCookie, err := r.Cookie("session_token")
	if err != nil {
		responseHandler.UnauthenticatedResponse()
		return
	}

	err = h.authService.LogOutUser(ctx, sessionTokenCookie.Value)
	if err != nil {
		responseHandler.UnauthenticatedResponse()
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   -1,
	})
	responseHandler.SuccessResponse(true, http.StatusOK)
}
