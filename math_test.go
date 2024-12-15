package gouse

import (
	"testing"
)

/*Testing for formulas*/
func TestLog(t *testing.T) {
	result := Log(10, 10)
	if result != 1 {
		t.Errorf("Log() = %d; want 1", result)
	}
}
