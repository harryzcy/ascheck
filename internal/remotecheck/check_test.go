package remotecheck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInfo(t *testing.T) {
	info, err := getInfo()
	assert.Nil(t, err)
	assert.IsType(t, []appInfo{}, info)
}
