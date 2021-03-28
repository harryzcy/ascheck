package remotecheck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	err := Init()
	assert.Nil(t, err)
}
