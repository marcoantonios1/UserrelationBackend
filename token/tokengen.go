package token

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SignedDetails struct {
	Uid         primitive.ObjectID
	Logged      bool
	Blocked     bool
	DeviceId    string
	DeviceType  string
	OpSys       string
	Model       string
	Country     string
	Locality    string
	AdminLevel1 string
	AdminLevel2 string
	Permission  []string
	jwt.RegisteredClaims
}

func ValidateToken(signedtoken string) (*SignedDetails, string) {
	// Read secret key from a secure location (e.g., encrypted configuration file)
	secretKey, err := GetSecretKey()
	if err != nil {
		return nil, "failed to retrieve secret key: " + err.Error()
	}

	// Parse the token with claims
	token, err := jwt.ParseWithClaims(signedtoken, &SignedDetails{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, "invalid token signature"
		}
		// Type assert claims to SignedDetails
		claims, ok := token.Claims.(*SignedDetails)
		if !ok {
			return nil, "invalid token claims structure"
		}

		// Validate expiration using standard claims
		if !token.Valid {
			return nil, "token is expired"
		}

		if claims.Blocked {
			return nil, "blocked user"
		}

		// Validate specific permission (can be extended for other checks)
		if !hasPermissions(claims.Permission, "PROFILE") {
			return nil, "insufficient permission"
		}
		return nil, "failed to parse token: " + err.Error()
	}

	// Type assert claims to SignedDetails
	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		return nil, "invalid token claims structure"
	}

	// Validate expiration using standard claims
	if !token.Valid {
		return nil, "token is expired"
	}

	if claims.Blocked {
		return nil, "blocked user"
	}

	// Validate specific permission (can be extended for other checks)
	if !hasPermissions(claims.Permission, "PROFILE") {
		return nil, "insufficient permission"
	}

	return claims, ""
}

// Helper function to check permission existence
func hasPermissions(permissions []string, required string) bool {
	for _, perm := range permissions {
		if perm == required {
			return true
		}
	}
	return false
}

// GetSecretKey should be implemented to read the secret key from a secure location
// This example omits the implementation for brevity.
func GetSecretKey() (string, error) {
	secretKey := os.Getenv("SECRET_USER_KEY")
	if secretKey == "" {
		return "", errors.New("missing SECRET_USER_KEY environment variable")
	}
	return secretKey, nil
}
