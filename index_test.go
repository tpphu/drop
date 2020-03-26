package drop

import (
	"testing"
)

func TestLoad(t *testing.T) {
	v := Load(1)
	if v == nil {
		t.Error("Should not nil")
	}
}
