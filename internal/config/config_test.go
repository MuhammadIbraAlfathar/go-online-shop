package config

import (
	"github.com/test-go/testify/require"
	"log"
	"testing"
)

func TestLoadConfig(t *testing.T) {

	t.Run("Success", func(t *testing.T) {
		filename := "../../cmd/config.yaml"
		err := LoadConfig(filename)

		require.Nil(t, err)
		log.Printf("%+v\n", Cfg)
	})
}
