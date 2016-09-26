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

const DELIMITER_COMMAND string = "_"
const DELIMITER_PARAM string = "="
const DELIMITER_SLICE string = ","

// setup - Download and configuration file handling, as well as obtaining
// the parameters of the environment or the command line.
func (this *Tuner) setup(path string) error {
	this.commandLineParser()
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
		//fmt.Print("\n========= ", str, " ==============")
		str = strings.TrimSpace(str)
		if len(str) < 3 || str[:1] == "#" {
			continue
		}
		if str[:1] == "[" {
			str = this.delComment(str)
			str = this.cleanSection(str, "[", "]")
			section = str
		} else if !strings.Contains(str, DELIMITER_PARAM) {
			return errors.New("Error in string " + strconv.Itoa(num) + " file " + path)
			//continue
		} else {
			str = this.delComment(str)
			k, v, err := this.exstractParam(str)
			//fmt.Print("\n=== ??? !==", k, " || ", v, " || ", reflect.TypeOf(v))
			if reflect.TypeOf(v) == reflect.TypeOf(errors.New("")) || err != nil {
				//fmt.Print("\n===!!!!!!!==", err) //v.(error), "|||||", v)
				return err //v.(error) //return errors.New("Error in string " + strconv.Itoa(num) + " file " + path)
			}
			if section == "" {
				//log.Printf("Errror in string %s file %s\n", strconv.Itoa(num), path)
				return errors.New("Error in string " + strconv.Itoa(num) + " file " + path)
			}
			if this.params[section] == nil {
				this.params[section] = make(map[string]interface{})
			}
			//os.Setenv("Main_path", "konfa.ini")
			//os.Setenv("Main_massiv", `{2015,2016,"kuku"}`)
			cKey := section + DELIMITER_COMMAND + k
			// config file
			this.params[section][k] = v
			// environment
			if e := os.Getenv(cKey); e != "" {
				if res, err := this.setTypeEnvPar(e, v, cKey); err != nil { // this.parseValue(e)
					return err
				} else {
					this.params[section][k] = res
				}
			}
			// command line
			if v2, ok := this.args[cKey]; ok {
				res, err := this.parseValue(v2)
				if err != nil {
					return err
				} else {
					this.params[section][k] = res
				}
			}
		}
	}
	//fmt.Print("\n - - - -   ", this.params)
	return nil
}

// setTypeEnvPar - Get the parameter from the environment.
func (this *Tuner) setTypeEnvPar(e string, v interface{}, key string) (interface{}, error) {
	var out interface{}
	if reflect.TypeOf(v) == reflect.TypeOf(string("")) {
		e = `"` + e + `"`
	}
	out, err := this.parseValue(e)
	return out, err
}

// exstractParam - Separation line with the parameter on the key and value.
func (this *Tuner) exstractParam(str string) (string, interface{}, error) {
	arr := strings.Split(str, DELIMITER_PARAM)
	key := strings.TrimSpace(arr[0])
	//value2 := strings.TrimSpace(arr[1])
	value := strings.Replace(str, arr[0]+DELIMITER_PARAM, "", 1)
	value = strings.TrimSpace(value)
	//fmt.Print("\n--------------", value, " |", arr[0], "| ", value2)
	v, err := this.parseValue(value)
	//fmt.Print("\n--------------", key, " |", v, "| ", err)
	return key, v, err
}

// parseValue - Processing parameter value, taking into account
// the type: strings, int, float64, slice.
func (this *Tuner) parseValue(value string) (interface{}, error) {
	var out interface{}
	//out = errors.New("Error in method 'parseValue' with string: " + value)
	var er error
	//fmt.Print("\n-- |", string(value[0]), "| ", value)
	switch string(value[0]) {
	case `"`, `'`:
		//fmt.Print("\n-- 1")
		out = this.parseValueString(value)
	case `(`, `{`:
		//fmt.Print("\n-- 2")
		out = this.parseValueSlice(value)
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
				//fmt.Print("\n == == == ", value, " ", v)
			} else {
				er = err
			}
		}
	}
	return out, er
}

// parseValueString - line processing (can be made out single or double quotes).
func (this *Tuner) parseValueString(value string) string {
	if string(value[0]) == `"` && string(value[len(value)-1]) == `"` {
		value = strings.Trim(value, `"`)
	} else if string(value[0]) == `'` && string(value[len(value)-1]) == `'` {
		value = strings.Trim(value, `'`)
	}
	return value
}

// parseValueSlice - slice processing, it can be made out simple or braces.
func (this *Tuner) parseValueSlice(value string) []interface{} {
	var out []interface{}
	if string(value[0]) == `(` && string(value[len(value)-1]) == `)` {
		value = this.cleanSection(value, "(", ")")
	} else if string(value[0]) == `{` && string(value[len(value)-1]) == `}` {
		value = this.cleanSection(value, "{", "}")
	}
	out = this.parseValueStringToSlice(value)
	return out
}

// parseValueStringToSlice - Processing contents of the slice,
// nested slices are not allowed.
func (this *Tuner) parseValueStringToSlice(value string) []interface{} {
	var out []interface{}
	arr := strings.Split(value, DELIMITER_SLICE)
	for _, v := range arr {
		v = strings.TrimSpace(v)
		p, err := this.parseValue(v)
		if reflect.TypeOf(p) != reflect.TypeOf(errors.New("")) && err == nil {
			out = append(out, p)
		}
	}
	return out
}
