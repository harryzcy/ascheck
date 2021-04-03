package remotecheck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSupport_Parse(t *testing.T) {
	tests := []struct {
		in       string
		expected Support
	}{
		{"✅", SupportNative},
		{"✳️", SupportTransition},
		{"⏹", SupportInDevelopment},
		{"🚫", SupportNotYet},
		{"🔶", SupportUnknown},
		{"some other", SupportUndefined},
	}

	for _, test := range tests {
		var support Support
		actual := support.Parse(test.in)

		assert.Equal(t, test.expected, support)
		assert.Equal(t, test.expected, actual)
	}
}

func TestSupport_String(t *testing.T) {
	tests := []struct {
		support  Support
		expected string
	}{
		{SupportNative, "Supported"},
		{SupportTransition, "Supported*"},
		{SupportInDevelopment, "Unsupported"},
		{SupportNotYet, "Unsupported"},
		{SupportUnknown, "Unknown"},
		{SupportUndefined, "Unknown"},
	}

	for _, test := range tests {
		actual := test.support.String()
		assert.Equal(t, test.expected, actual)
	}
}
