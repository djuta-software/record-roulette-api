package httpclient

import (
	"encoding/json"
	"errors"
	"net/http"
)

// GetJSON decodes json to an interface
func GetJSON(url string, out interface{}) error {
	response, err := http.Get(url)
	if response.StatusCode != http.StatusOK {
		err = errors.New("oops")
	}
	if err != nil {
		return err
	}
	json.NewDecoder(response.Body).Decode(out)
	response.Body.Close()
	return nil
}
