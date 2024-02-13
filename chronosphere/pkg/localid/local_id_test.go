package localid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsLocalID(t *testing.T) {
	assert.True(t, IsLocalID(NewLocalID()))
	assert.True(t, IsLocalID(NewImportedID()))
	assert.True(t, IsLocalID("1234567")) // Older format, no longer used but possibly present in state files

	assert.False(t, IsLocalID(""))
	assert.False(t, IsLocalID("foo"))
	assert.False(t, IsLocalID("foo-123"))
	assert.False(t, IsLocalID("foo-bar"))
}
