package require

import (
	"errors"
	"reflect"
	"testing"
)

func Equal(t *testing.T, expected any, actual any) {
	t.Helper()
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected %v but got %v", expected, actual)
	}
}

func NotEqual(t *testing.T, expected any, actual any) {
	t.Helper()
	if reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected %v to not equal %v", expected, actual)
	}
}

func Nil(t *testing.T, actual any) {
	t.Helper()
	if !isNil(actual) {
		t.Fatalf("expected nil but got (%T) %v", actual, actual)
	}
}

func NotNil(t *testing.T, actual any) {
	t.Helper()
	if isNil(actual) {
		t.Fatalf("expected not nil but got nil")
	}
}

func True(t *testing.T, actual bool) {
	t.Helper()
	if !actual {
		t.Fatalf("expected true but got false")
	}
}

func isNil(v any) bool {
	if v == nil {
		return true
	}

	value := reflect.ValueOf(v)
	kind := value.Kind()

	if kind >= reflect.Chan && kind <= reflect.Slice {
		return value.IsNil()
	}

	return false
}

func ErrorAs(t *testing.T, err error, target interface{}) {
	t.Helper()
	if !errors.As(err, target) {
		t.Fatalf("expected %T but got %T", target, err)
	}
}

func ErrorIs(t *testing.T, err, target error) {
	t.Helper()
	if !errors.Is(err, target) {
		t.Fatalf("expected %v but got %v", target, err)
	}
}
