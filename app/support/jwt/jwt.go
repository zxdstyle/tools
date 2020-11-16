package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"tools/app/support/h"
)

func (mw *AuthMiddleware) GenerateToken(data map[string]interface{}) (string, error) {
	token := jwt.New(jwt.GetSigningMethod(mw.SigningAlgorithm))
	claims := token.Claims.(jwt.MapClaims)
	if mw.PayloadFunc != nil {
		for key, value := range mw.PayloadFunc(data) {
			claims[key] = value
		}
	}

	expire := mw.TimeFunc().Add(mw.Timeout)
	claims["exp"] = expire.Unix()
	claims["orig_iat"] = mw.TimeFunc().Unix()
	tokenString, err := mw.signedString(token)
	if err != nil {
		return "", ErrFailedTokenCreation
	}

	return tokenString, nil
}

func (mw *AuthMiddleware) signedString(token *jwt.Token) (string, error) {
	var tokenString string
	var err error
	if mw.usingPublicKeyAlgo() {
		tokenString, err = token.SignedString(mw.privateKey)
	} else {
		tokenString, err = token.SignedString(mw.Key)
	}
	return tokenString, err
}

// ExtractClaims help to extract the JWT claims
func ExtractClaims(r *ghttp.Request) MapClaims {
	claims := r.GetParam("JWT_PAYLOAD")
	return claims.(MapClaims)
}

// GetToken help to get the JWT token string
func GetToken(r *ghttp.Request) string {
	token := r.GetString("JWT_TOKEN")
	if len(token) == 0 {
		return ""
	}

	return token
}

// New for check error with GfJWTMiddleware
func New(m *AuthMiddleware) (*AuthMiddleware, error) {
	if err := m.MiddlewareInit(); err != nil {
		return nil, err
	}

	return m, nil
}

// MiddlewareInit initialize jwt configs.
func (mw *AuthMiddleware) MiddlewareInit() error {

	if mw.TokenLookup == "" {
		mw.TokenLookup = "header:Authorization"
	}

	if mw.SigningAlgorithm == "" {
		mw.SigningAlgorithm = "HS256"
	}

	if mw.Timeout == 0 {
		mw.Timeout = time.Hour
	}

	if mw.TimeFunc == nil {
		mw.TimeFunc = time.Now
	}

	mw.TokenHeadName = strings.TrimSpace(mw.TokenHeadName)
	if len(mw.TokenHeadName) == 0 {
		mw.TokenHeadName = "Bearer"
	}

	if mw.Authorization == nil {
		mw.Authorization = func(data interface{}, r *ghttp.Request) bool {
			return true
		}
	}

	if mw.LoginResponse == nil {
		mw.LoginResponse = func(r *ghttp.Request, code int, token string, expire time.Time) {
			r.Response.WriteJson(g.Map{
				"code":   http.StatusOK,
				"token":  token,
				"expire": expire.Format(time.RFC3339),
			})
		}
	}

	if mw.RefreshResponse == nil {
		mw.RefreshResponse = func(r *ghttp.Request, code int, token string, expire time.Time) {
			r.Response.WriteJson(g.Map{
				"code":   http.StatusOK,
				"token":  token,
				"expire": expire.Format(time.RFC3339),
			})
		}
	}

	if mw.IdentityKey == "" {
		mw.IdentityKey = IdentityKey
	}

	if mw.IdentityHandler == nil {
		mw.IdentityHandler = func(r *ghttp.Request) interface{} {
			claims := ExtractClaims(r)
			return claims[mw.IdentityKey]
		}
	}

	if mw.HTTPStatusMessageFunc == nil {
		mw.HTTPStatusMessageFunc = func(e error, r *ghttp.Request) string {
			return e.Error()
		}
	}

	if mw.Realm == "" {
		mw.Realm = "gf jwt"
	}

	if mw.CookieName == "" {
		mw.CookieName = "jwt"
	}

	if mw.usingPublicKeyAlgo() {
		return mw.readKeys()
	}

	if mw.Key == nil {
		return ErrMissingSecretKey
	}
	return nil
}

func (mw *AuthMiddleware) readKeys() error {
	err := mw.setPrivateKey()
	if err != nil {
		return err
	}
	err = mw.publicKey()
	if err != nil {
		return err
	}
	return nil
}

func (mw *AuthMiddleware) setPrivateKey() error {
	keyData, err := ioutil.ReadFile(mw.PrivateKeyFile)
	if err != nil {
		return ErrNoPrivateKeyFile
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if err != nil {
		return ErrInvalidPrivateKey
	}
	mw.privateKey = key
	return nil
}

func (mw *AuthMiddleware) publicKey() error {
	keyData, err := ioutil.ReadFile(mw.PubKeyFile)
	if err != nil {
		return ErrNoPubKeyFile
	}
	key, err := jwt.ParseRSAPublicKeyFromPEM(keyData)
	if err != nil {
		return ErrInvalidPubKey
	}
	mw.pubKey = key
	return nil
}

func (mw *AuthMiddleware) usingPublicKeyAlgo() bool {
	switch mw.SigningAlgorithm {
	case "RS256", "RS512", "RS384":
		return true
	}
	return false
}

func (mw *AuthMiddleware) MiddlewareFunc() ghttp.HandlerFunc {
	return func(r *ghttp.Request) {
		mw.middlewareImpl(r)
	}
}

func (mw *AuthMiddleware) GetClaimsFromJWT(r *ghttp.Request) (MapClaims, error) {
	token, err := mw.ParseToken(r)

	if err != nil {
		return nil, err
	}

	if mw.SendAuthorization {
		token := r.GetString("JWT_TOKEN")
		if len(token) > 0 {
			r.Header.Set("Authorization", mw.TokenHeadName+" "+token)
		}
	}

	claims := MapClaims{}
	for key, value := range token.Claims.(jwt.MapClaims) {
		claims[key] = value
	}

	return claims, nil
}

func (mw *AuthMiddleware) middlewareImpl(r *ghttp.Request) {
	claims, err := mw.GetClaimsFromJWT(r)
	if err != nil {
		h.Failed(r, mw.HTTPStatusMessageFunc(err, r), http.StatusUnauthorized)
		return
	}

	if claims["exp"] == nil {
		h.Failed(r, mw.HTTPStatusMessageFunc(ErrMissingExpField, r), http.StatusBadRequest)
		return
	}

	if _, ok := claims["exp"].(float64); !ok {
		h.Failed(r, mw.HTTPStatusMessageFunc(ErrWrongFormatOfExp, r), http.StatusBadRequest)
		return
	}

	if int64(claims["exp"].(float64)) < mw.TimeFunc().Unix() {
		h.Failed(r, mw.HTTPStatusMessageFunc(ErrExpiredToken, r), http.StatusBadRequest)
		return
	}

	r.SetParam("JWT_PAYLOAD", claims)
	identity := mw.IdentityHandler(r)

	if identity != nil {
		r.SetParam(mw.IdentityKey, identity)
	}

	if !mw.Authorization(identity, r) {
		h.Failed(r, mw.HTTPStatusMessageFunc(ErrForbidden, r), http.StatusForbidden)
		return
	}

	//c.Next() todo
}

func (mw *AuthMiddleware) ParseToken(r *ghttp.Request) (*jwt.Token, error) {
	var token string
	var err error

	methods := strings.Split(mw.TokenLookup, ",")
	for _, method := range methods {
		if len(token) > 0 {
			break
		}
		parts := strings.Split(strings.TrimSpace(method), ":")
		k := strings.TrimSpace(parts[0])
		v := strings.TrimSpace(parts[1])
		switch k {
		case "header":
			token, err = mw.jwtFromHeader(r, v)
		case "query":
			token, err = mw.jwtFromQuery(r, v)
		case "cookie":
			token, err = mw.jwtFromCookie(r, v)
		case "param":
			token, err = mw.jwtFromParam(r, v)
		}
	}

	if err != nil {
		return nil, err
	}

	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod(mw.SigningAlgorithm) != t.Method {
			return nil, ErrInvalidSigningAlgorithm
		}
		if mw.usingPublicKeyAlgo() {
			return mw.pubKey, nil
		}

		// save token string if validate
		r.SetParam("JWT_TOKEN", token)

		return mw.Key, nil
	})
}

func (mw *AuthMiddleware) jwtFromHeader(r *ghttp.Request, key string) (string, error) {
	authHeader := r.Header.Get(key)

	if authHeader == "" {
		return "", ErrNotFoundToken
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == mw.TokenHeadName) {
		return "", ErrInvalidAuthHeader
	}

	return parts[1], nil
}

func (mw *AuthMiddleware) jwtFromQuery(r *ghttp.Request, key string) (string, error) {
	token := r.GetString(key)

	if token == "" {
		return "", ErrNotFoundToken
	}

	return token, nil
}

func (mw *AuthMiddleware) jwtFromCookie(r *ghttp.Request, key string) (string, error) {
	cookie := r.Cookie.Get(key)

	if cookie == "" {
		return "", ErrNotFoundToken
	}

	return cookie, nil
}

func (mw *AuthMiddleware) jwtFromParam(r *ghttp.Request, key string) (string, error) {
	token := r.GetString(key)
	if token == "" {
		return "", ErrNotFoundToken
	}

	return token, nil
}
