package localcheck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplication_GetExecutableName(t *testing.T) {
	exec, err := getExecutableName("./../../test/data/py_app.app")
	assert.Nil(t, err)
	assert.Equal(t, "run.py", exec)
}

func TestApplication_GetArchitectures(t *testing.T) {
	arch, err := GetArchitectures("./../../test/data/example_macho.app")
	assert.Nil(t, err)
	assert.EqualValues(t, 0b01, arch.Intel)
	assert.EqualValues(t, 0, arch.PowerPC)
	assert.EqualValues(t, 0, arch.Arm)
}

func TestGetInterpreterPath(t *testing.T) {
	tests := []struct {
		file         string
		expectedPath []string
		expectedOK   bool
	}{
		{"./../../test/data/bash.sh", []string{"/bin/bash"}, true},
		{"./../../test/data/env_bash.sh", []string{"/bin/bash", "/usr/bin/bash"}, true},
	}

	for _, test := range tests {
		path, ok := getInterpreterPath(test.file)
		assert.Contains(t, test.expectedPath, path)
		assert.Equal(t, test.expectedOK, ok)
	}
}
