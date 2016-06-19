package explorer

import (
    "testing"
    "reflect"
    "net/url"
)
const TEST_FILENAME = "__test_conf_file.json"
func makeMockApp() *App {
    return &App{
        Root: url.Parse("https://hostess/rest/v0/")
        Current: url.Parse("/rest/v0/endpoint")

    }
}

func TestLoadSave(t *testing.T) {
    a := makeMockApp()
    err := a.SaveConfig(TEST_FILENAME)
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
}

