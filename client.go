package fxa

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type Credentials struct {
	Email  string `json:"email"`
	AuthPW string `json"authPW"`
}

type Client struct {
	AuthServer string
}

func NewClient() (*Client, error) {

	cl := Client{
		AuthServer: "https://api.accounts.firefox.com/v1",
	}

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

	creds := Credentials{
		Email:  email,
		AuthPW: hex.EncodeToString(k),
	}

	enc_creds, err := json.Marshal(creds)

	if err != nil {
		return err
	}

	fh := bytes.NewReader(enc_creds)

	url := fmt.Sprintf("%s/account/login", cl.AuthServer)

	rsp, err := http.Post(url, "application/json", fh)

	if err != nil {
		return err
	}

	if rsp.StatusCode != 200 {
		return errors.New(rsp.Status)
	}

	log.Println(rsp)
	return nil
}
