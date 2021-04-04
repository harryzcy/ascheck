package localcheck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppInfo_GetDisplayName(t *testing.T) {
	info := AppInfo{CFBundleDisplayName: "test-display-name"}
	assert.Equal(t, "test-display-name", info.GetDisplayName())

	info = AppInfo{CFBundleDisplayName: "", CFBundleName: "test"}
	assert.Equal(t, "test", info.GetDisplayName())
}

func TestAppInfo_GetExecutable(t *testing.T) {
	info := AppInfo{CFBundleExecutable: "test-executable"}
	assert.Equal(t, "test-executable", info.GetExecutableName())
}

func TestGetAppInfo(t *testing.T) {
	tests := []struct {
		path        string
		name        string
		displayName string
		executable  string
		hasErr      bool
	}{
		{"./../../test/data/example_macho.app", "Test App", "", "example", false},
		{"./../../test/data/invalid.app", "", "", "", true},
		{"./../../test/data/invalid_plist.app", "", "", "", true},
	}

	for _, test := range tests {
		info, err := GetAppInfo(test.path, "")
		assert.Equal(t, test.path, info.path)
		assert.Equal(t, test.name, info.CFBundleName)
		assert.Equal(t, test.displayName, info.CFBundleDisplayName)
		assert.Equal(t, test.executable, info.CFBundleExecutable)
		if test.hasErr {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}
	}
}
