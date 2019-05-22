package config

import (
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig(t *testing.T) {
	err := Load(V)
	assert.NoError(t, err)
	assert.Equal(t, V.Port, "3000")
	assert.Equal(t, viper.Get("PORT"), "3000")
}
