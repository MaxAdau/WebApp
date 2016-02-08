package main

import (

	"net/url"
)


type APIObj interface {
	Handler(*url.URL) string
}