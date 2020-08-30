package socketactivation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstantErrorType(t *testing.T) {
	e := err("category 55 emergency doomsday crisis")

	assert.EqualError(t, e, "category 55 emergency doomsday crisis")
}
