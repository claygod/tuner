package tuner

// Tuner
// Test
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	//"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	tnr, err := New("config.ini")
	if tnr == nil || err != nil {
		t.Error("Error init Tuner struct")
	}
}

func TestNothingFile(t *testing.T) {
	tnr, err := New("2016.ini")
	if tnr != nil || err == nil {
		t.Error("Error processing file is missing")
	}
}

func TestWrongFile(t *testing.T) {
	tnr, err := New("wrong.ini")
	if tnr != nil || err == nil {
		t.Error("Do not handle the error in a special file with errors")
	}
}

func TestGetCorrectKey(t *testing.T) {
	tnr, _ := New("config.ini")
	if x := tnr.Section("Main").Get("path"); x != "config.ini" {
		t.Error("Error get correct get key `path`: ", x)
	}
}

func TestGetUncorrectKey(t *testing.T) {
	tnr, err := New("config.ini")
	if x := tnr.Section("Main").Get("pathhh"); x != nil {
		t.Error("Error get uncorrect get key `pathhh`: ", err)
	}
}

func TestGetUncorrectSection(t *testing.T) {
	tnr, err := New("config.ini")
	if x := tnr.Section("Mainnn").Get("path"); x != nil {
		t.Error("Error get uncorrect get section `Mainnn`: ", err)
	}
}

func TestOnlySection(t *testing.T) {
	tnr, err := New("config.ini")
	if x := tnr.Section("Main"); x != tnr {
		t.Error("Error only method `Section`: ", err)
	}
}

func TestOnlyGet(t *testing.T) {
	tnr, err := New("config.ini")
	if x := tnr.Get("path"); x != nil {
		t.Error("Error only method `Get`", err)
	}
}

func TestList(t *testing.T) {
	tnr, _ := New("config.ini")
	m := tnr.Section("Main").Get("mas").([]interface{})
	y := m[0].(int)
	if y != 6 {
		t.Error("Error get element from list")
	}

}
