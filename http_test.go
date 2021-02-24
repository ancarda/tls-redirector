package main

import (
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"testing/fstest"
	"testing/quick"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	TextPlain = "text/plain"
	TextHTML  = "text/html; charset=utf-8"
)

func randomString() string {
	v, ok := quick.Value(reflect.TypeOf(""),
		rand.New(rand.NewSource(time.Now().Unix())))

	if !ok {
		panic("wasn't able to generate a string")
	}

	return v.String()
}

func assertionsCommonToAllResponses(t *testing.T, res *http.Response) {
	assert.Equal(t, "nosniff", res.Header.Get("X-Content-Type-Options"))
	assert.Equal(t, "tls-redirector/2.4", res.Header.Get("Server"))
}

func TestServer_ServeACME_404(t *testing.T) {
	rr := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodGet, "http://nowhere.invalid/.well-known/acme-challenge/z", nil)
	app{fstest.MapFS{}, true}.ServeHTTP(rr, r)
	res := rr.Result()

	assert.Equal(t, http.StatusNotFound, res.StatusCode)
	assert.Equal(t, TextHTML, res.Header.Get("Content-Type"))
	assertionsCommonToAllResponses(t, res)

	body, _ := ioutil.ReadAll(res.Body)
	assert.Contains(t, string(body), "File Not Found")
}

func TestServer_ServeACME_HappyPath(t *testing.T) {
	mapFS := fstest.MapFS{
		"ok": &fstest.MapFile{Data: []byte("12345678")},
	}

	rr := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodGet, "http://nowhere.invalid/.well-known/acme-challenge/ok", nil)
	app{mapFS, true}.ServeHTTP(rr, r)
	res := rr.Result()

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, TextPlain, res.Header.Get("Content-Type"))
	assertionsCommonToAllResponses(t, res)

	body, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, "12345678", string(body))
}

func TestServer_NoHostHeader_WillError(t *testing.T) {
	rr := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodGet, "", nil)
	app{}.ServeHTTP(rr, r)
	res := rr.Result()

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Equal(t, TextHTML, res.Header.Get("Content-Type"))
	assertionsCommonToAllResponses(t, res)

	body, _ := ioutil.ReadAll(res.Body)
	assert.Contains(t, string(body), "header is empty or wasn't sent")
}

func TestServer_IPAddressHostHeader_IsRejected(t *testing.T) {
	rr := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodGet, "http://127.0.0.1/", nil)
	app{}.ServeHTTP(rr, r)
	res := rr.Result()

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Equal(t, TextHTML, res.Header.Get("Content-Type"))
	assertionsCommonToAllResponses(t, res)

	body, _ := ioutil.ReadAll(res.Body)
	assert.Contains(t, string(body), "looks like an IP address")
}

func TestServer_HappyPath(t *testing.T) {
	rr := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodGet, "http://nowhere.invalid/", nil)
	app{}.ServeHTTP(rr, r)
	res := rr.Result()

	assert.Equal(t, http.StatusMovedPermanently, res.StatusCode)
	assert.Equal(t, TextHTML, res.Header.Get("Content-Type"))
	assertionsCommonToAllResponses(t, res)
	assert.Equal(t, "https://nowhere.invalid/", res.Header.Get("Location"))
}
