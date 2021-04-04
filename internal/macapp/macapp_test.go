package macapp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllApplications(t *testing.T) {
	apps, err := GetAllApplications(nil, "")
	assert.Nil(t, err)
	assert.IsType(t, []Application{}, apps)
}
