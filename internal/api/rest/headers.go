package rest

import "strings"

const (
	authHeader     = "Authorization"
	usernameHeader = "X-Username"
	creatorHeader  = "X-Is-Creator"
)

// bearerAuthHeader trims the "Bearer" prefix and returns the token.
// If there is no "Bearer" prefix, it will return an empty string.
func bearerAuthHeader(authHeader string) string {
	headerParts := strings.Split(authHeader, "Bearer")
	if len(headerParts) != 2 {
		return ""
	}

	return strings.TrimSpace(headerParts[1])
}
