package utils

import (
    // "context"
    "time"
    "errors"
    "os"
    // "fmt"
    jwt "github.com/dgrijalva/jwt-go"
)

// The user claim embedded with the JWT standard claim.
type AccessTokenClaims struct {
    UserId int64 `json:"user_id"`
    ThingId int64 `json:"ThingId"`
    jwt.StandardClaims
}


// Function returns the `secret key` used by our web service for signing
// our JWT standard claims.
func GetAccessTokenSecretKey() ([]byte) {
    accessTokenString := os.Getenv("MIKAPONICS_SECRET_KEY")
    accessTokenBytes := []byte(accessTokenString)
    return accessTokenBytes
}


// Function returns token string which was generated.
func GenerateAccessToken(
    userId int64,
    thingId int64,
    durationInMinutes time.Duration,
) (string, error) {
    jwtKey := GetAccessTokenSecretKey()

    // Create the JWT claims, which includes the userId and expir.
    claims := AccessTokenClaims{
        userId,
        thingId,
        jwt.StandardClaims{
            // In JWT, the expiry time is expressed as unix milliseconds
            ExpiresAt: time.Now().Add(time.Minute * durationInMinutes).Unix(), // MAKE SHORT-LIVED
            Issuer:    "test",
        },
    }

    // Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    // Create the JWT string
	tokenString, err := token.SignedString(jwtKey)

    return tokenString, err
}


// Function returns the verified claims if the JWT claim was verified, else
// return the error.
func VerifyAccessTokenString(tokenString string) (*AccessTokenClaims, error) {
    // Initialize a new instance of `Claims`
	claims := &AccessTokenClaims{}

    // Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return GetAccessTokenSecretKey(), nil
	})

    // If the token is invalid then we'll generate an error else return our results.
    if !token.Valid {
        return nil, errors.New("Invalid token")
    }
    return claims, err
}
