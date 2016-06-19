package explorer

import (
	"encoding/json"
	"io/ioutil"
    "os"
)

func LoadConfig(filename string) (*App, error) {
    a := &App{}
	file, err := os.Open(filename)
	if err != nil {
		return a, err
	}

	jsonParser := json.NewDecoder(file)
	err = jsonParser.Decode(&a)
    return a, err
}

func (a *App) SaveConfig(filename string) error {
	b, err := json.Marshal(a)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, []byte(b), 0644)
	return err
}
