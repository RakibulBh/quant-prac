package main

import "encoding/json"

func (app *application) ReadJSON(jsonData []byte, container any) error {
	return json.Unmarshal([]byte(jsonData), container)
}
