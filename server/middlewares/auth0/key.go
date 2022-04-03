package auth0

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JSONWebKeys struct {
	Kty string   `json:"kty,omitempty"`
	Kid string   `json:"kid,omitempty"`
	Use string   `json:"use,omitempty"`
	N   string   `json:"n,omitempty"`
	E   string   `json:"e,omitempty"`
	X5c []string `json:"x_5_c,omitempty"`
}

type JWKS struct {
	Keys []JSONWebKeys `json:"keys"`
}

func FetchJWKS(auth0Domain string) (*JWKS, error) {
	resp, err := http.Get(fmt.Sprintf("https://%s/.well-known/jwks.json", auth0Domain))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	jwks := &JWKS{}
	err = json.NewDecoder(resp.Body).Decode(jwks)

	return jwks, err
}
