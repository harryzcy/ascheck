package remotecheck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	err := Init()
	assert.Nil(t, err)
}

func TestGetInfo(t *testing.T) {
	err := Init()
	assert.Nil(t, err)

	info, err := GetInfo("Go (golang)")
	assert.Nil(t, err)
	assert.Equal(t, "✅", info.ArmSupport)
}
