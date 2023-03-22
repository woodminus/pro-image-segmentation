package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinOfTwoIntegers(t *testing.T) {
	assert.Equal(t, 3, MinI(3, 4))
	