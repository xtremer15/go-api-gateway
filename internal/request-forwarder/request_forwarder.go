package requestforwarder

// This package will act as a standalone request router towards specific services.
// It provides minimal routing and authentication middleware.

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

// Route represents the target URL of the service we want to route the request to.
type Route struct {
	TargetURL string
}

// Routes defines the current mapping of prefixes to internal service targets.
var Routes = map[string]Route{
	"/path1/": {TargetURL: "http://auth-service:port_nr"},
	"/path2/": {TargetURL: "http://payment-service:port_nr"},
	"/path3/": {TargetURL: "http://analytics-service:port_nr"},
}

// RouteHandler forwards incoming requests to the matching target service.
func RouteHandler(w http.ResponseWriter, r *http.Request) {
	for prefix, route := range Routes {
		if strings.HasPrefix(r.URL.Path, prefix) {
			ProxyRequest(w, r, route.TargetURL)
			return
		}
	}
	http.Error(w, "Not Found", http.StatusNotFound)
}

// ProxyRequest creates a reverse proxy to the specified targetURL and serves the request.
func ProxyRequest(w http.ResponseWriter, r *http.Request, targetURL string) {
	target, err := url.Parse(targetURL)
	if err != nil {
		http.Error(w, "Bad Gateway: Invalid target URL", http.StatusBadGateway)
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.ServeHTTP(w, r)
}

// AuthMiddleware intercepts requests to validate Bearer JWT tokens.
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Parse the token and validate the signing method
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate the algorithm (HMAC in this case context)
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("your-secret-key"), nil // Replace with real secret in production
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}
