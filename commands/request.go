package commands

import (
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"io"
	"mime/multipart"
	"strings"
)

var ErrorFormData = errors.New("invalid form data")

func (x *HttpRequest) GetHeaderValue(k string) string {
	for _, key := range x.Header {
		if strings.ToLower(key.GetKey()) == strings.ToLower(k) {
			if len(key.GetValue()) > 0 {
				return key.GetValue()[0]
			}
		}
	}

	return ""
}

func (x *HttpRequest) BasicAuth() (username, password string, ok bool) {
	const prefix = "Basic "
	auth := x.GetHeaderValue("Authorization")
	if auth == "" {
		return "", "", false
	}

	c, err := base64.StdEncoding.DecodeString(auth[len(prefix):])
	if err != nil {
		return "", "", false
	}

	cs := string(c)
	username, password, ok = strings.Cut(cs, ":")
	if !ok {
		return "", "", false
	}
	return username, password, true
}

func (x *HttpRequest) CheckBasicAuth(login, pass string) bool {
	reqLogin, reqPass, ok := x.BasicAuth()
	if ok {
		requestLoginHash := sha256.Sum256([]byte(reqLogin))
		reqquestPassHash := sha256.Sum256([]byte(reqPass))
		loginHash := sha256.Sum256([]byte(login))
		passHash := sha256.Sum256([]byte(pass))

		uMatch := subtle.ConstantTimeCompare(requestLoginHash[:], loginHash[:]) == 1
		pMatch := subtle.ConstantTimeCompare(reqquestPassHash[:], passHash[:]) == 1

		if uMatch && pMatch {
			return true
		}
	}

	return false
}

func (x *HttpRequest) IsJson() bool {
	return x.GetHeaderValue("Content-Type") == "application/json"
}

func (x *HttpRequest) IsFormData() (boundary string, ok bool) {
	const formDataPrefix = "multipart/form-data"
	header := x.GetHeaderValue("Content-Type")
	contType, boundary, ok := strings.Cut(header, ";")
	if ok && strings.ToLower(contType) == formDataPrefix {
		_, boundary, ok = strings.Cut(boundary, "=")
		if ok {
			return boundary, ok
		}
	}
	return "", false
}

func (x *HttpRequest) ParseFormData(r io.Reader) (*multipart.Form, error) {
	boundary, ok := x.IsFormData()
	if !ok {
		return nil, ErrorFormData
	}

	reader := multipart.NewReader(r, boundary)
	return reader.ReadForm(0)
}

func (x *HttpRequest) IsBrowserRequest() bool {
	userAgent := x.GetHeaderValue("User-Agent")
	if strings.Contains(userAgent, "Mozilla") {
		return true
	}

	return false
}
