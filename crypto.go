package fxa

// https://github.com/mozilla/PyFxA/blob/6b5d33686176b1730b31951c4da15a61faed131e/fxa/crypto.py

import (
	"crypto/sha256"
	"fmt"
	"golang.org/x/crypto/hkdf"
	"golang.org/x/crypto/pbkdf2"
	"io"
)

func QuickStretchPassword(email string, password string) ([]byte, error) {

	salt := hkdf_namespace("quickStretch", email)

	k := pbkdf2.Key([]byte(password), []byte(salt), 1000, 32, sha256.New)
	return k, nil
}

func DeriveKey(secret []byte, namespace string) ([]byte, error) {

	length := 32 // sudo make me a param?

	salt := ""
	info := hkdf_namespace(namespace)

	hkdf := hkdf.New(sha256.New, secret, []byte(salt), []byte(info))

	k := make([]byte, length)
	_, err := io.ReadFull(hkdf, k)

	if err != nil {
		return nil, err
	}

	return k, nil
}

func hkdf_namespace(name string, extra ...string) string {

	ns := fmt.Sprintf("identity.mozilla.com/picl/v1/%s", name)

	for _, e := range extra {
		ns = fmt.Sprintf("%s:%s", ns, e)
	}

	return ns
}
