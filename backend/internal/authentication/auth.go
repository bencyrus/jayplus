package authentication

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Auth struct {
	Issuer             string        `json:"issuer"`
	Audience           string        `json:"audience"`
	Secret             string        `json:"secret"`
	AccessTokenExpiry  time.Duration `json:"access_token_expiry"`
	RefreshTokenExpiry time.Duration `json:"refresh_token_expiry"`
	CookieDomain       string        `json:"cookie_domain"`
	CookiePath         string        `json:"cookie_path"`
	CookieName         string        `json:"cookie_name"`
}

type JWTUser struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type JWTTokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type JWTClaims struct {
	jwt.RegisteredClaims
}

func NewAuth(a Auth) *Auth {
	return &Auth{
		Issuer:             a.Issuer,
		Audience:           a.Audience,
		Secret:             a.Secret,
		AccessTokenExpiry:  a.AccessTokenExpiry,
		RefreshTokenExpiry: a.RefreshTokenExpiry,
		CookieDomain:       a.CookieDomain,
		CookiePath:         a.CookiePath,
		CookieName:         a.CookieName,
	}
}

func (a *Auth) GenerateTokenPair(user *JWTUser) (JWTTokenPair, error) {
	if user == nil {
		return JWTTokenPair{}, fmt.Errorf("error generating token pair: user is nil")
	}

	// create access token
	accessToken := jwt.New(jwt.SigningMethodHS256)

	accessTokenClaims := accessToken.Claims.(jwt.MapClaims)
	accessTokenClaims["name"] = fmt.Sprintf("%s %s", user.FirstName, user.LastName)
	accessTokenClaims["sub"] = fmt.Sprint(user.ID)
	accessTokenClaims["aud"] = a.Audience
	accessTokenClaims["iss"] = a.Issuer
	accessTokenClaims["iat"] = time.Now().UTC().Unix()
	accessTokenClaims["typ"] = "JWT"
	accessTokenClaims["exp"] = time.Now().UTC().Add(a.AccessTokenExpiry).Unix()

	signedAccessToken, err := accessToken.SignedString([]byte(a.Secret))
	if err != nil {
		return JWTTokenPair{}, fmt.Errorf("error signing access token: %w", err)
	}

	// create refresh token
	refreshToken := jwt.New(jwt.SigningMethodHS256)

	refreshTokenClaims := refreshToken.Claims.(jwt.MapClaims)
	refreshTokenClaims["sub"] = fmt.Sprint(user.ID)
	refreshTokenClaims["iat"] = time.Now().UTC().Unix()
	refreshTokenClaims["exp"] = time.Now().UTC().Add(a.RefreshTokenExpiry).Unix()

	signedRefreshToken, err := refreshToken.SignedString([]byte(a.Secret))
	if err != nil {
		return JWTTokenPair{}, fmt.Errorf("error signing refresh token: %w", err)
	}

	return JWTTokenPair{
		AccessToken:  signedAccessToken,
		RefreshToken: signedRefreshToken,
	}, nil
}

func (a *Auth) GetRefreshCookie(refreshToken string) *http.Cookie {
	return &http.Cookie{
		Name:     a.CookieName,
		Value:    refreshToken,
		Path:     a.CookiePath,
		Expires:  time.Now().UTC().Add(a.RefreshTokenExpiry),
		MaxAge:   int(a.RefreshTokenExpiry.Seconds()),
		SameSite: http.SameSiteStrictMode,
		Domain:   a.CookieDomain,
		HttpOnly: true,
		Secure:   true,
	}
}

func (a *Auth) GetExpiredRefreshCookie() *http.Cookie {
	return &http.Cookie{
		Name:     a.CookieName,
		Value:    "",
		Path:     a.CookiePath,
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		SameSite: http.SameSiteStrictMode,
		Domain:   a.CookieDomain,
		HttpOnly: true,
		Secure:   true,
	}

}
