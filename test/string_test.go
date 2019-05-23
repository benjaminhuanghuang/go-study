package stringstudy

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringSplitJoin(t *testing.T) {
	s := "A, B, C"
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}

	t.Log(strings.Join(parts, "-"))
}

func TestStringToNumber(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("str " + s)

	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(i)
	}
}