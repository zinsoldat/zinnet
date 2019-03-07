package util_test

import (
	"testing"

	"github.com/zinsoldat/zinnet-go/util"
)

func TestRandString(t *testing.T) {
	const stringLength = 5
	randomString := util.RandString(stringLength)

	if len(randomString) != stringLength {
		t.Errorf("the random string should have a length of %d but it has a length of %d", stringLength, len(randomString))
	}
}
