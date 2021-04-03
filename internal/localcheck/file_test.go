package localcheck

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsText(t *testing.T) {
	tests := []struct {
		in       []byte
		expected bool
	}{
		{[]byte("#!/bin/bash\n"), true},
		{[]byte("some string"), true},
		{[]byte("#!/usr/bin/env bash\n"), true},
		{[]byte(strings.Repeat("some string ", 100)), true},
		{[]byte{0x00, 0x01, 0x02, 0x3}, false},
	}

	for _, test := range tests {
		actual := IsText(test.in)
		assert.Equal(t, test.expected, actual)
	}
}
