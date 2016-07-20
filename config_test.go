package explorer

import (
	"os"
	"reflect"
	"testing"
)

const TEST_FILENAME = "__test_conf_file.json"

func makeMockApp() (*App, error) {
	// c, _ := url.Parse("/rest/v0/endpoint")
	a, err := NewApp("https://hostess/rest/v0/", "user", "pass", true)
	return a, err
}

func TestLoadSave(t *testing.T) {
	a, err := makeMockApp()
	if err != nil {
		t.Errorf("error creating mock app: %s", err)
	}
	err = a.SaveConfig(TEST_FILENAME)
	if err != nil {
		t.Errorf("error saving config: %s", err)
	}
	b, err := LoadConfig(TEST_FILENAME)
	if err != nil {
		t.Errorf("error loading config: %s", err)
	}
	if !reflect.DeepEqual(a, b) {
		t.Error("the loaded and saved configs were different!")
	}
	os.Remove(TEST_FILENAME)
}
