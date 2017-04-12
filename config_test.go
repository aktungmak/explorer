package explorer

import (
	"github.com/aktungmak/odata-client"
	"net/url"
	"os"
	"reflect"
	"testing"
)

const TEST_FILENAME = "__test_conf_file.json"

func makeMockApp() (*App, error) {
	sr, _ := url.Parse("https://hostess/rest/v0/")
	c := odata.NewBaClient("user", "pass", true)
	// TODO test other client types
	a, err := NewApp(sr, c)
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
	} else {
		os.Remove(TEST_FILENAME)
	}
}
