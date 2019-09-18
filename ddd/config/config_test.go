package config

import (
	"os"
	"reflect"
	"testing"
)

func TestNewConfig(t *testing.T) {

	os.Setenv("MYAPP_PORT", "8080")
	got, err := NewConfig()
	if err != nil {
		t.Error(err)
	}
	want := &Config{
		Port: 8080,
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Unexpected configuration; got: %v, want: %v", got, want)
	}
}
