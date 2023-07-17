package authentication

import (
	"fmt"
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

func (j *Auth) GenerateTokenPair(user *JWTUser) (JWTTokenPair, error) {
	// create access token
	accessToken := jwt.New(jwt.SigningMethodHS256)

	accessTokenClaims := accessToken.Claims.(jwt.MapClaims)
	accessTokenClaims["name"] = fmt.Sprintf("%s %s", user.FirstName, user.LastName)
	accessTokenClaims["sub"] = fmt.Sprint(user.ID)
	accessTokenClaims["aud"] = j.Audience
	accessTokenClaims["iss"] = j.Issuer
	accessTokenClaims["iat"] = time.Now().UTC().Unix()
	accessTokenClaims["typ"] = "JWT"
	accessTokenClaims["exp"] = time.Now().UTC().Add(j.AccessTokenExpiry).Unix()

	signedAccessToken, err := accessToken.SignedString([]byte(j.Secret))
	if err != nil {
		return JWTTokenPair{}, fmt.Errorf("error signing access token: %w", err)
	}

	// create refresh token
	refreshToken := jwt.New(jwt.SigningMethodHS256)

	refreshTokenClaims := refreshToken.Claims.(jwt.MapClaims)
	refreshTokenClaims["sub"] = fmt.Sprint(user.ID)
	refreshTokenClaims["iat"] = time.Now().UTC().Unix()
	refreshTokenClaims["exp"] = time.Now().UTC().Add(j.RefreshTokenExpiry).Unix()

	signedRefreshToken, err := refreshToken.SignedString([]byte(j.Secret))
	if err != nil {
		return JWTTokenPair{}, fmt.Errorf("error signing refresh token: %w", err)
	}

	return JWTTokenPair{
		AccessToken:  signedAccessToken,
		RefreshToken: signedRefreshToken,
	}, nil
}
