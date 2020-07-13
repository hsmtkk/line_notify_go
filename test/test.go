package test

import "testing"

func AssertNil(t *testing.T, v interface{}) {
	t.Helper()
	if v != nil {
		t.Errorf("want nil, got %v", v)
	}
}

func AssertEqualString(t *testing.T, want, got string) {
	t.Helper()
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}
