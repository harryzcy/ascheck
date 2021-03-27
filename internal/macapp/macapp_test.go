package macapp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplication_GetExecutableName(t *testing.T) {
	app := Application{
		Path: "/System/Applications/Maps.app",
	}

	exec, err := app.GetExecutableName()
	assert.Nil(t, err)
	assert.Equal(t, "Maps", exec)
}

func TestApplication_GetArchitectures(t *testing.T) {
	app := Application{
		Path: "/System/Applications/Maps.app",
	}
	arch, err := app.GetArchitectures()
	assert.Nil(t, err)
	assert.Equal(t, []string{"Intel 64"}, arch)
}

func TestGetAllApplications(t *testing.T) {
	apps, err := GetAllApplications(nil)
	assert.Nil(t, err)
	assert.IsType(t, []Application{}, apps)
}
