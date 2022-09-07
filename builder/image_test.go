package builder_test

import (
	"testing"

	"github.com/victorfernandesraton/hq-now-dowloader/builder"
)

func TestResizeImage(t *testing.T) {
	t.Run("resize simple image for specific output", func(t *testing.T) {
		builder.ResizeImage("test.jpg", "test_output.jpg")
	})
}
