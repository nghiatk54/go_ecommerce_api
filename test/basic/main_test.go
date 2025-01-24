package basic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestAddOne is a test function for AddOne
func TestAddOne(t *testing.T) {
	assert.Equal(t, AddOne(2), 3, "AddOne(2) should be 3")
	assert.NotEqual(t, AddOne(2), 4, "AddOne(2) should not be 4")
}

// Test feature require for testify
func TestRequire(t *testing.T) {
	require.Equal(t, 2, 3)
	fmt.Println("After require")
}

// Test feature assert for testify
func TestAssert(t *testing.T) {
	assert.Equal(t, 2, 3)
	fmt.Println("After assert")
}
