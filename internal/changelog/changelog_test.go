package changelog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetURL(t *testing.T) {
	t.Run("is semantic version tag", func(t *testing.T) {
		want := "https://github.com/leozhantw/gojira/releases/tag/v1.2.3"
		got := GetURL("1.2.3")
		assert.Equal(t, want, got)
	})

	t.Run("is not semantic version tag", func(t *testing.T) {
		want := "https://github.com/leozhantw/gojira/releases/latest"
		got := GetURL("dev")
		assert.Equal(t, want, got)
	})
}
