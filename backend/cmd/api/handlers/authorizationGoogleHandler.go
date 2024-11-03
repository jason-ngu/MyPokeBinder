package handlers

import (
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

type App struct {
	config *oauth2.Config
}

func (h handler) NewGoogleAuth() App {
	clientid := h.Config.Oauth.Google.ClientID
	clientSecret := h.Config.Oauth.Google.ClientSecret

	conf := &oauth2.Config{
		ClientID:     clientid,
		ClientSecret: clientSecret,
		RedirectURL:  "http://localhost:8080/auth/callback/google",
		Scopes:       []string{"profile"},
		Endpoint:     google.Endpoint,
	}

	app := App{config: conf}
	return app
}

func (a *App) GoogleLoginHandler(w http.ResponseWriter, r *http.Request) {
	url := a.config.AuthCodeURL("state", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (a *App) GoogleCallbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	t, err := a.config.Exchange(context.Background(), code)
	if err != nil {
		log.Fatal(err)
	}

	client := a.config.Client(context.Background(), t)

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
