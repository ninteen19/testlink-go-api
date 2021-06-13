package testlink

import (
	"net/http"
)

type Config struct {
	Url        string
	Key        string
	HttpClient *http.Client
}

var Conf Config = Config{
	HttpClient: &http.Client{},
}
