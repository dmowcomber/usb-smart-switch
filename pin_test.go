package main

import (
	"testing"

	"github.com/brutella/hc"
	"github.com/stretchr/testify/assert"
)

func TestGeneratePin(t *testing.T) {
	// generated pins should be valid
	for i := 0; i < 10000; i++ {
		pin, _ := generatePin()
		formattedPin, err := hc.ValidatePin(pin)
		assert.NoError(t, err)
		assert.NotEmpty(t, formattedPin)
	}
}
