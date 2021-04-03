package localcheck

import (
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
	}

	for _, test := range tests {
		actual := IsText(test.in)
		assert.Equal(t, test.expected, actual)
	}
}
