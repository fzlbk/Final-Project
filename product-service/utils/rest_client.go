package utils

import "github.com/go-resty/resty/v2"

var UserClient = resty.New().SetHostURL("http://localhost:8081")
