package tuner

// Tuner
// Helper
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	//"fmt"
	"os"
	"strings"
)

func (my *Tuner) commandLineParser() {
	var arg []string
	arg = os.Args
	for i := len(arg) - 1; i > 0; i-- {
		arg[i] = strings.TrimLeft(arg[i], "-")
		twoSide := strings.Split(arg[i], DELIMITER_PARAM)
		key := twoSide[0]
		value := strings.Replace(arg[i], key+DELIMITER_PARAM, "", 1)
		my.args[key] = value
	}
}

func (my *Tuner) delComment(str string) string {
	wos := strings.Split(str, "#")
	out := strings.TrimSpace(wos[0])
	out = strings.TrimSpace(out)
	return out
}

func (my *Tuner) cleanSection(str string, left string, right string) string {
	str = strings.TrimSpace(str)
	str = strings.TrimLeft(str, left)
	str = strings.TrimRight(str, right)
	return str
}
