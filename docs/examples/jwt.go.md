
## Hmac - jwt.StandardClaims
```go
var sampleSecret = []byte("my_secret_key")

func newToken() (string, error) {
	claims := jwt.StandardClaims{
		Audience:  "*",
		ExpiresAt: time.Now().Add(24 * 7 * time.Hour).Unix(),
		Id:        "b90e07d6-8732-4857-9433-5b9a6715719e",
		IssuedAt:  time.Now().Unix(),
		Issuer:    "zlab.dev",
		NotBefore: time.Now().Add(-60 * time.Second).Unix(),
		Subject:   "json web token",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString(sampleSecret)
}

func parseToken(tokenString string) (*jwt.StandardClaims, error) {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return sampleSecret, nil
	})

	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
```


## Hmac - jwt.MapClaims
```go
var sampleSecret = []byte("my_secret_key")

func newToken() (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo":  "bar",
		"name": "zlab.dev",
		"nbf":  time.Now().Add(-60 * time.Second).Unix(),
	})
	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString(sampleSecret)
}

func parseToken(tokenString string) (jwt.MapClaims, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return sampleSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
```


## Hmac - CustomClaimsType
```go

var sampleSecret = []byte("my_secret_key")

type MyCustomClaims struct {
	Foo string `json:"foo"`
	jwt.StandardClaims
}

func newToken() (string, error) {

	claims := MyCustomClaims{
		"bar",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 3600,
			Issuer:    "zlab.dev",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(sampleSecret)
}

func parseToken(tokenString string) (*MyCustomClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return sampleSecret, nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
```



## docs
https://pkg.go.dev/github.com/golang-jwt/jwt#pkg-examples
