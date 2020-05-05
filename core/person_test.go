package core

import (
	"testing"
)

func TestIsInSameTeam(t *testing.T) {
	result := isInSameTeam("a b  c ", "  c x y")
	if result == false {
		t.Errorf("Should be same team")
	}

	result = isInSameTeam("a b c", "d e f")
	if result == true {
		t.Errorf("Should be different team")
	}
}
