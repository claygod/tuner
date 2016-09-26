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
	this := &Tuner{}
	this.params = make(map[string]map[string]interface{})
	this.args = make(map[string]string)
	err := this.setup(path)
	if err != nil {
		return nil, err
	}
	return this, nil
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
func (this *Tuner) Get(key string) interface{} {
	if this.current != "" &&
		this.params[this.current] != nil {
		if v, ok := this.params[this.current][key]; ok {
			this.current = ""
			return v
		}
	}
	this.current = ""
	return nil
}

// Set - Set new value (the key must exist!).
// This method sure to use only after the method `Section`.
// Both methods are used in a pair!
func (this *Tuner) Set(key string, value interface{}) error {
	if this.current != "" && this.params[this.current] != nil {
		this.params[this.current][key] = value
		return nil
	}
	this.current = ""
	return errors.New("Not set value")
}

// Section - Specify a section to produce or modify a variable.
// This method sure to use only before the methods `Set` or `Get`.
// Both methods are used in a pair!
func (this *Tuner) Section(key string) *Tuner {
	if this.params[key] != nil {
		this.current = key
	}
	//fmt.Print("\n111")
	return this
}
