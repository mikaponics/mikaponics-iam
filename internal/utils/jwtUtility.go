package utils

import (
    // "context"
    "time"
    "fmt"
    jwt "github.com/dgrijalva/jwt-go"
)

// The user claim embedded with the JWT standard claim.
type MyCustomClaims struct {
    UserId int64 `json:"user_id"`
    jwt.StandardClaims
}



func GenerateJWTToken(userId int64) (string, error) {
    mySigningKey := []byte("AllYourBase")

    // Create the Claims
    claims := MyCustomClaims{
        userId,
        jwt.StandardClaims{
            ExpiresAt: time.Now().Add(time.Minute * 15).Unix(), // MAKE SHORT-LIVED
            Issuer:    "test",
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    ss, err := token.SignedString(mySigningKey)
    return ss, err
}


func VerifyJWTTokenString(tokenString string) (bool, error) {
    token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte("AllYourBase"), nil
    })

    if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
        // fmt.Printf("%v %v", claims.UserId, claims.StandardClaims.ExpiresAt)
        fmt.Printf("> %v <", claims)
    } else {
        return false, err
    }
    return true, nil
}
