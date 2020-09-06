package fancy

import (
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"
	"time"

	"github.com/stretchr/testify/assert"
)

func randomString() string {
	v, ok := quick.Value(reflect.TypeOf(""),
		rand.New(rand.NewSource(time.Now().Unix())))

	if !ok {
		panic("wasn't able to generate a string")
	}

	return v.String()
}

func TestErrorPage(t *testing.T) {
	title := randomString()
	message := randomString()
	techInfo := randomString()
	version := randomString()

	page := string(ErrorPage(title, message, techInfo, version))

	assert.Contains(t, page, "<title>"+title+"</title>")
	assert.Contains(t, page, "<h1>"+title+"</h1>")
	assert.Contains(t, page, message)
	assert.Contains(t, page, techInfo)
	assert.Contains(t, page, "tls-redirector/"+version)
}
