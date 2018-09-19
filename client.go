package fxa

import (
	"encoding/hex"
	"log"
	_ "net/http"
)

type Login struct {
	Email  string `json:"email"`
	AuthPW string `json"authPW"`
}

type Client struct {
}

func NewClient() (*Client, error) {

	cl := Client{}

	return &cl, nil
}

// https://github.com/mozilla/PyFxA/blob/6b5d33686176b1730b31951c4da15a61faed131e/fxa/core.py

func (cl *Client) Login(email string, password string) error {

	s_password, err := QuickStretchPassword(email, password)

	if err != nil {
		return err
	}

	k, err := DeriveKey(s_password, "authPW")

	if err != nil {
		return err
	}

	l := Login{
		Email:  email,
		AuthPW: hex.EncodeToString(k),
	}

	log.Println(l)
	return nil
}
