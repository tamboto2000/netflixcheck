// Package netflixcheck is a lib for Netflix account checker
package netflixcheck

import (
	"os"

	"github.com/tamboto2000/json5extract"
)

// Account contains information about Netflix account
type Account struct {
	Username   string `json:"username"`
	Name       string `json:"name"`
	Screens    int    `json:"screens"`
	Language   string `json:"language"`
	ValidUntil string `json:"validUntil"`
	Working    bool   `json:"working"`
}

func TestLoginPage() error {
	raw, err := get(nil, nil, "/login")
	if err != nil {
		return err
	}

	f, err := os.Create("login.html")
	if err != nil {
		return err
	}

	if _, err := f.Write(raw); err != nil {
		return err
	}

	jsons, err := json5extract.FromBytes(raw)
	if err != nil {
		return err
	}

	return json5extract.Save(jsons)
}
