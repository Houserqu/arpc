package arpc

import "testing"

func TestParsePageSize(t *testing.T) {
	offset, size := ParsePageSize(1, 10)
	if offset != 0 || size != 10 {
		t.Errorf("Expected offset=0, size=10 but got offset=%d, size=%d", offset, size)
	}

	offset, size = ParsePageSize(2, 30, int32(20))
	if offset != 20 || size != 20 {
		t.Errorf("Expected offset=20, size=20 but got offset=%d, size=%d", offset, size)
	}

	offset, size = ParsePageSize(3, 50)
	if offset != 40 || size != 20 {
		t.Errorf("Expected offset=40, size=20 but got offset=%d, size=%d", offset, size)
	}

	offset, size = ParsePageSize(0, 0)
	if offset != 0 || size != 20 {
		t.Errorf("Expected offset=40, size=20 but got offset=%d, size=%d", offset, size)
	}

	offset, size = ParsePageSize(0, 0, 50)
	if offset != 0 || size != 50 {
		t.Errorf("Expected offset=40, size=20 but got offset=%d, size=%d", offset, size)
	}

	offset, size = ParsePageSize(nil, "a", 50)
	if offset != 0 || size != 50 {
		t.Errorf("Expected offset=40, size=20 but got offset=%d, size=%d", offset, size)
	}
}
