package tuner

// Tuner
// Work
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"errors"
	//"flag"
	//"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// DELIMITER_COMMAND - is used to bind the section name and key name
const DELIMITER_COMMAND string = "_"

// DELIMITER_PARAM - is used to separate key and value
const DELIMITER_PARAM string = "="

// DELIMITER_SLICE - is used to separate components of the List
const DELIMITER_SLICE string = ","

// setup - Download and configuration file handling, as well as obtaining
// the parameters of the environment or the command line.
func (my *Tuner) setup(path string) error {
	my.commandLineParser()
	//out := true
	confFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Print(err)
		return err
	}
	confStr := strings.Split(string(confFile), "\n")
	var section string
	for num, strng := range confStr {
		str := strng
		str = strings.TrimSpace(str)
		if len(str) < 3 || str[:1] == "#" {
			continue
		}
		if str[:1] == "[" {
			str = my.delComment(str)
			str = my.cleanSection(str, "[", "]")
			section = str
		} else if !strings.Contains(str, DELIMITER_PARAM) {
			return errors.New("Error in string " + strconv.Itoa(num) + " file " + path)
		} else {
			str = my.delComment(str)
			k, v, err := my.exstractParam(str)
			if reflect.TypeOf(v) == reflect.TypeOf(errors.New("")) || err != nil {
				return err
			}
			if section == "" {
				return errors.New("Error in string " + strconv.Itoa(num) + " file " + path)
			}
			if my.params[section] == nil {
				my.params[section] = make(map[string]interface{})
			}
			//os.Setenv("Main_path", "konfa.ini")
			//os.Setenv("Main_massiv", `{2015,2016,"kuku"}`)
			cKey := section + DELIMITER_COMMAND + k
			// config file
			my.params[section][k] = v
			// environment
			if e := os.Getenv(cKey); e != "" {
				if res, err := my.setTypeEnvPar(e, v, cKey); err != nil {
					return err
				}
				my.params[section][k] = res

			}
			// command line
			if v2, ok := my.args[cKey]; ok {
				res, err := my.parseValue(v2)
				if err != nil {
					return err
				}
				my.params[section][k] = res
			}
		}
	}
	//fmt.Print("\n - - - -   ", my.params)
	return nil
}

// setTypeEnvPar - Get the parameter from the environment.
func (my *Tuner) setTypeEnvPar(e string, v interface{}, key string) (interface{}, error) {
	var out interface{}
	if reflect.TypeOf(v) == reflect.TypeOf(string("")) {
		e = `"` + e + `"`
	}
	out, err := my.parseValue(e)
	return out, err
}

// exstractParam - Separation line with the parameter on the key and value.
func (my *Tuner) exstractParam(str string) (string, interface{}, error) {
	arr := strings.Split(str, DELIMITER_PARAM)
	key := strings.TrimSpace(arr[0])
	value := strings.Replace(str, arr[0]+DELIMITER_PARAM, "", 1)
	value = strings.TrimSpace(value)
	v, err := my.parseValue(value)
	return key, v, err
}

// parseValue - Processing parameter value, taking into account
// the type: strings, int, float64, slice.
func (my *Tuner) parseValue(value string) (interface{}, error) {
	var out interface{}
	var er error
	switch string(value[0]) {
	case `"`, `'`:
		out = my.parseValueString(value)
	case `(`, `{`:
		out = my.parseValueSlice(value)
	default:
		if string(value) == "true" || string(value) == "false" {
			if v, err := strconv.ParseBool(value); err == nil {
				out = v
			} else {
				er = err
			}
		} else if strings.Count(value, ".") == 1 {
			if v, err := strconv.ParseFloat(value, 64); err == nil {
				out = v
			} else {
				er = err
			}
		} else if strings.Count(value, ".") == 0 {
			if v, err := strconv.Atoi(value); err == nil {
				out = v
			} else {
				er = err
			}
		}
	}
	return out, er
}

// parseValueString - line processing (can be made out single or double quotes).
func (my *Tuner) parseValueString(value string) string {
	if string(value[0]) == `"` && string(value[len(value)-1]) == `"` {
		value = strings.Trim(value, `"`)
	} else if string(value[0]) == `'` && string(value[len(value)-1]) == `'` {
		value = strings.Trim(value, `'`)
	}
	return value
}

// parseValueSlice - slice processing, it can be made out simple or braces.
func (my *Tuner) parseValueSlice(value string) []interface{} {
	var out []interface{}
	if string(value[0]) == `(` && string(value[len(value)-1]) == `)` {
		value = my.cleanSection(value, "(", ")")
	} else if string(value[0]) == `{` && string(value[len(value)-1]) == `}` {
		value = my.cleanSection(value, "{", "}")
	}
	out = my.parseValueStringToSlice(value)
	return out
}

// parseValueStringToSlice - Processing contents of the slice,
// nested slices are not allowed.
func (my *Tuner) parseValueStringToSlice(value string) []interface{} {
	var out []interface{}
	arr := strings.Split(value, DELIMITER_SLICE)
	for _, v := range arr {
		v = strings.TrimSpace(v)
		p, err := my.parseValue(v)
		if reflect.TypeOf(p) != reflect.TypeOf(errors.New("")) && err == nil {
			out = append(out, p)
		}
	}
	return out
}
