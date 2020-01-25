package placeholder

import (
	"testing"
)

func TestPlaceholder(t *testing.T) {
	if placeholder("Placeholder") != "Placeholder" {
		t.Error("Test error placeholder")
	}
}
