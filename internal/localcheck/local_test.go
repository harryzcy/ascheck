package localcheck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplication_GetExecutableName(t *testing.T) {
	exec, err := getExecutableName("/System/Applications/Maps.app")
	assert.Nil(t, err)
	assert.Equal(t, "Maps", exec)
}

func TestApplication_GetArchitectures(t *testing.T) {
	arch, err := GetArchitectures("/System/Applications/Maps.app")
	assert.Nil(t, err)
	assert.Equal(t, []string{"Intel 64"}, arch)
}
