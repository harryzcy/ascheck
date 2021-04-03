package localcheck

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplication_GetExecutableName(t *testing.T) {
	exec, err := getExecutableName("./../../test/data/sh_app.app")
	assert.Nil(t, err)
	assert.Equal(t, "run.sh", exec)
}

func TestApplication_GetArchitectures(t *testing.T) {
	arch, err := GetArchitectures("./../../test/data/example_macho.app")
	assert.Nil(t, err)
	assert.EqualValues(t, 0b01, arch.Intel)
	assert.EqualValues(t, 0, arch.PowerPC)
	assert.EqualValues(t, 0, arch.Arm)

	arch, err = GetArchitectures("./../../test/data/example_fat.app")
	assert.Nil(t, err)
	assert.EqualValues(t, 0b10, arch.Intel)
	assert.EqualValues(t, 0, arch.PowerPC)
	assert.EqualValues(t, 0, arch.Arm)

	arch, err = GetArchitectures("./../../test/data/sh_app.app")
	assert.Nil(t, err)
	assert.NotEmpty(t, arch)
	if err != nil {
		t.Error(err)
	}
}

func TestApplication_GetArchitectures_Error(t *testing.T) {
	// Invalid path
	arch, err := GetArchitectures("./../../test/data/invalid.app")
	assert.NotNil(t, err)
	assert.True(t, os.IsNotExist(err))
	assert.Empty(t, arch)

	// Invalid plist
	arch, err = GetArchitectures("./../../test/data/invalid_plist.app")
	assert.NotNil(t, err)
	assert.False(t, os.IsNotExist(err))
	assert.Empty(t, arch)

	// Invalid interpreter
	arch, err = GetArchitectures("./../../test/data/invalid_interpreter.app")
	assert.NotNil(t, err)
	assert.Equal(t, errors.New("unable to get executable path"), err)
	assert.Empty(t, arch)

	// Unknown file type
	arch, err = GetArchitectures("./../../test/data/unknown_type.app")
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
