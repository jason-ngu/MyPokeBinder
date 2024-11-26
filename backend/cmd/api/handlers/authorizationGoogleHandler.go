package handlers

import (
	"backend/config"
	services "backend/internal"
	"backend/internal/models"
	usersService "backend/internal/services/users"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type googleHandler struct {
	oauthconfig *oauth2.Config
}

func (h *handler) NewGoogleHandler(config *config.Config) *googleHandler {
	clientid := config.Oauth.Google.ClientID
	clientSecret := config.Oauth.Google.ClientSecret

	conf := &oauth2.Config{
		ClientID:     clientid,
		ClientSecret: clientSecret,
		RedirectURL:  "http://localhost:8080/auth/callback/google",
		Scopes:       []string{"profile"},
		Endpoint:     google.Endpoint,
	}

	return &googleHandler{oauthconfig: conf}
}

func (h *googleHandler) GoogleLoginHandler(w http.ResponseWriter, r *http.Request) {
	url := h.oauthconfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (h *googleHandler) GoogleCallbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	t, err := h.oauthconfig.Exchange(context.Background(), code)
	if err != nil {
		log.Fatal(err)
	}

	client := h.oauthconfig.Client(context.Background(), t)

	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	var decodedResp map[string]string

	// Reading the JSON body using JSON decoder
	err = json.NewDecoder(resp.Body).Decode(&decodedResp)
	if err != nil {
		log.Fatal(err)
	}

	foundUser := models.UserModel{
		Name:         decodedResp["name"],
		ProviderKey:  decodedResp["id"],
		ProviderType: "Google",
	}

	_, err = usersService.EnsureUser(&services.Env{}, foundUser)
	if err != nil {
		log.Fatal(err)
	}

	// Redirect to user's collections page
	http.Redirect(w, r, "/collections", http.StatusTemporaryRedirect)
}
