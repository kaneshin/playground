package trick

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBoolTimer(t *testing.T) {
	bt := BoolTimer(time.Second)

	assert.False(t, bt())
	assert.True(t, bt())
	assert.True(t, bt())
	assert.True(t, bt())
	assert.True(t, bt())
	assert.True(t, bt())
	assert.True(t, bt())

	time.Sleep(time.Second)

	assert.False(t, bt())
	assert.True(t, bt())
	assert.True(t, bt())
	assert.True(t, bt())
	assert.True(t, bt())
	assert.True(t, bt())
	assert.True(t, bt())
}
