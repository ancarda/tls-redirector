package main

import (
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
	"testing/quick"
	"time"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

const (
	TextPlain = "text/plain"
	TextHtml  = "text/html"
)

func randomString() string {
	v, ok := quick.Value(reflect.TypeOf(""),
		rand.New(rand.NewSource(time.Now().Unix())))

	if !ok {
		panic("wasn't able to generate a string")
	}

	return v.String()
}

func TestNewApp_UsesRealFileSystem(t *testing.T) {
	app := NewApp("")

	exists, err := afero.DirExists(app.fs, os.TempDir())
	assert.True(t, exists)
	assert.Nil(t, err)
}

func TestNewApp_StoresEnvSettings(t *testing.T) {
	someString := randomString()

	app := NewApp(someString)
	assert.Equal(t, someString, app.acmeChallengeDir)
}

func TestServer_ServeACME_404(t *testing.T) {
	rr := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodGet, "http://nowhere.invalid"+acmeChallengeUrlPrefix+"z", nil)
	App{afero.NewMemMapFs(), "/"}.ServeHTTP(rr, r)
	res := rr.Result()

	assert.Equal(t, http.StatusNotFound, res.StatusCode)
	assert.Equal(t, TextPlain, res.Header.Get("Content-Type"))
	assert.Equal(t, "nosniff", res.Header.Get("X-Content-Type-Options"))

	body, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, "File Not Found\n", string(body))

}

func TestServer_ServeACME_HappyPath(t *testing.T) {
	fs := afero.NewMemMapFs()
	f, _ := fs.Create("/ok")
	f.Write([]byte("12345678"))

	rr := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodGet, "http://nowhere.invalid"+acmeChallengeUrlPrefix+"ok", nil)
	App{fs, "/"}.ServeHTTP(rr, r)
	res := rr.Result()

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, TextPlain, res.Header.Get("Content-Type"))
	assert.Equal(t, "nosniff", res.Header.Get("X-Content-Type-Options"))

	body, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, "12345678", string(body))
}

func TestServer_NoHostHeader_WillError(t *testing.T) {
	rr := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodGet, "", nil)
	App{}.ServeHTTP(rr, r)
	res := rr.Result()

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Equal(t, TextPlain, res.Header.Get("Content-Type"))
	assert.Equal(t, "nosniff", res.Header.Get("X-Content-Type-Options"))

	body, _ := ioutil.ReadAll(res.Body)
	assert.Contains(t, string(body), "no `host` header was sent")
}

func TestServer_IPAddressHostHeader_IsRejected(t *testing.T) {
	rr := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodGet, "http://127.0.0.1/", nil)
	App{}.ServeHTTP(rr, r)
	res := rr.Result()

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Equal(t, TextPlain, res.Header.Get("Content-Type"))
	assert.Equal(t, "nosniff", res.Header.Get("X-Content-Type-Options"))

	body, _ := ioutil.ReadAll(res.Body)
	assert.Contains(t, string(body), "cannot redirect IP addresses")
}

func TestServer_HappyPath(t *testing.T) {
	rr := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodGet, "http://nowhere.invalid/", nil)
	App{}.ServeHTTP(rr, r)
	res := rr.Result()

	assert.Equal(t, http.StatusMovedPermanently, res.StatusCode)
	assert.Equal(t, TextHtml, res.Header.Get("Content-Type"))
	assert.Equal(t, "nosniff", res.Header.Get("X-Content-Type-Options"))
	assert.Equal(t, "https://nowhere.invalid/", res.Header.Get("Location"))
}
