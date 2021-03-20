package main

import (
	"encoding/json"
	"errors"
)

type Foo struct {
	Type   string      `json:"type"`
	Object interface{} `json:"object"`
}

type A struct {
	A string `json:"a"`
}

type B struct {
	B string `json:"b"`
}

func (f *Foo) UnmarshalJSON(data []byte) error {
	type cloneType Foo

	rawMsg := json.RawMessage{}
	f.Object = &rawMsg

	if err := json.Unmarshal(data, (*cloneType)(f)); err != nil {
		return err
	}

	switch f.Type {
	case "a":
		params := new(A)
		if err := json.Unmarshal(rawMsg, params); err != nil {
			return err
		}
		f.Object = params
	case "b":
		params := new(B)
		if err := json.Unmarshal(rawMsg, params); err != nil {
			return err
		}
		f.Object = params
	default:
		return errors.New("nonsupport type")
	}

	return nil
}
