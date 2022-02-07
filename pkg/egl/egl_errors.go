package egl

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

type Error struct {
	Text    string
	Code    string `json:"errorCode"`
	Message string `json:"errorMessage"`
}

func (err *Error) Error() string {
	if err.Text != "" {
		return err.Text
	}
	return err.Message
}

func FindError(buf io.ReadCloser) error {
	raw, err := ioutil.ReadAll(buf)
	if err != nil {
		return err
	}
	buf.Close()

	data := &Error{}
	if err = json.Unmarshal(raw, data); err != nil {
		return err
	}
	return data
}
