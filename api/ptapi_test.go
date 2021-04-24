package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func skipShort(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
}

func TestGetUNIXUTCFor(t *testing.T) {
	unixTimestamp, err := getUNIXUTCFor("Europe/Vienna", 2021, 4, 24, 15, 26, 00)
	if assert.NoError(t, err) {
		assert.Equal(t, int64(1619270760), unixTimestamp)
	}
}
