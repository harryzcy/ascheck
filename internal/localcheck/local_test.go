package localcheck

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplication_GetArchitectures(t *testing.T) {
	info, err := GetAppInfo("./../../test/data/example_macho.app", "")
	assert.Nil(t, err)
	arch, err := info.GetArchitectures()
	assert.Nil(t, err)
	assert.EqualValues(t, 0b01, arch.Intel)
	assert.EqualValues(t, 0, arch.PowerPC)
	assert.EqualValues(t, 0, arch.Arm)

	info, err = GetAppInfo("./../../test/data/example_fat.app", "")
	assert.Nil(t, err)
	arch, err = info.GetArchitectures()
	assert.Nil(t, err)
	assert.EqualValues(t, 0b10, arch.Intel)
	assert.EqualValues(t, 0, arch.PowerPC)
	assert.EqualValues(t, 0, arch.Arm)

	info, err = GetAppInfo("./../../test/data/sh_app.app", "")
	assert.Nil(t, err)
	arch, err = info.GetArchitectures()
	if err == nil { // should pass on macOS
		assert.NotEmpty(t, arch)
	} else { // would failed on linux
		assert.Equal(t, errors.New("unknown file type"), err)
		assert.Empty(t, arch)
	}
}

func TestApplication_GetArchitectures_Error(t *testing.T) {
	// Invalid interpreter
	info, err := GetAppInfo("./../../test/data/invalid_interpreter.app", "")
	assert.Nil(t, err)
	arch, err := info.GetArchitectures()
	assert.NotNil(t, err)
	assert.Equal(t, errors.New("unable to get executable path"), err)
	assert.Empty(t, arch)

	// Unknown file type
	info, err = GetAppInfo("./../../test/data/unknown_type.app", "")
	assert.Nil(t, err)
	arch, err = info.GetArchitectures()
	assert.NotNil(t, err)
	assert.Equal(t, errors.New("unknown file type"), err)
	assert.Empty(t, arch)
}

func TestGetInterpreterPath(t *testing.T) {
	tests := []struct {
		file         string
		expectedPath []string
		expectedOK   bool
	}{
		{"./../../test/data/bash.sh", []string{"/bin/bash"}, true},
		{"./../../test/data/env_bash.sh", []string{"/bin/bash", "/usr/bin/bash"}, true},
		{"./../../test/data/invalid.sh", []string{""}, false},
		{"./../../test/data/env_invalid.sh", []string{""}, false},
		{"./../../test/data/invalid_shebang.sh", []string{""}, false},
	}

	for _, test := range tests {
		path, ok := getInterpreterPath(test.file)
		assert.Contains(t, test.expectedPath, path)
		assert.Equal(t, test.expectedOK, ok)
	}
}
