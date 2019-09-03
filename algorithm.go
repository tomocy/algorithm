package algorithm

import (
	"testing"
)

func Reportln(t *testing.T, name string, actual, expected interface{}) {
	t.Errorf("unexpected %s: got %v, expect %v\n", name, actual, expected)
}
