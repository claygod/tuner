package tuner

// Tuner
// API
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"errors"
	//"fmt"
)

// New - create a new Tuner-struct
func New(path string) (*Tuner, error) {
	my := &Tuner{}
	my.params = make(map[string]map[string]interface{})
	my.args = make(map[string]string)
	err := my.setup(path)
	if err != nil {
		return nil, err
	}
	return my, nil
}

// Tuner structure
type Tuner struct {
	confile string
	current string
	params  map[string]map[string]interface{}
	args    map[string]string
}

// Get - Specify the key that are looking for value.
// This method sure to use only after the method `Section`.
// Both methods are used in a pair!
func (my *Tuner) Get(key string) interface{} {
	if my.current != "" &&
		my.params[my.current] != nil {
		if v, ok := my.params[my.current][key]; ok {
			my.current = ""
			return v
		}
	}
	my.current = ""
	return nil
}

// Set - Set new value (the key must exist!).
// This method sure to use only after the method `Section`.
// Both methods are used in a pair!
func (my *Tuner) Set(key string, value interface{}) error {
	if my.current != "" && my.params[my.current] != nil {
		my.params[my.current][key] = value
		return nil
	}
	my.current = ""
	return errors.New("Not set value")
}

// Section - Specify a section to produce or modify a variable.
// This method sure to use only before the methods `Set` or `Get`.
// Both methods are used in a pair!
func (my *Tuner) Section(key string) *Tuner {
	if my.params[key] != nil {
		my.current = key
	}
	//fmt.Print("\n111")
	return my
}
