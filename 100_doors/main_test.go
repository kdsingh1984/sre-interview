package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpenDoorsList(t *testing.T) {
	result := openDoorsList(map[int]bool{1: true, 2: false, 3: true, 4: false})
	assert.Equal(t, []int{1, 2, 3}, result)
}
