package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorPage(t *testing.T) {
	title := randomString()
	message := randomString()
	techInfo := randomString()
	version := randomString()

	page := string(errorPage(title, message, techInfo, version))

	assert.Contains(t, page, "<title>"+title+"</title>")
	assert.Contains(t, page, "<h1>"+title+"</h1>")
	assert.Contains(t, page, message)
	assert.Contains(t, page, techInfo)
	assert.Contains(t, page, "tls-redirector/"+version)
}
