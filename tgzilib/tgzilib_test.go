package tgzilib

import (
	"reflect"
	"testing"
)

func TestGetFiles(t *testing.T) {
	inputDir := "."
	want := []string{"tgzilib.go", "tgzilib_test.go"}

	result, err := GetFiles(inputDir, true)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if !reflect.DeepEqual(result, want) {
		t.Fatalf("Got %v, wanted %v", result, want)
	}
}
