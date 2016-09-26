// Tuner
// Documentation RU
//
/*
Тюнер предназначен для удобного конфигурирования приложения.
Параметры он берёт из (в порядке уменьшения приоритета):
- командной строки
- окружения
- файла конфигурации

Пример использования

	tun, err := tuner.New(confile)
	if err != nil {
		panic(err) //fmt.Sprintf("%v", i))
	}
	path := tun.Section("Main").Get("path").(string)
	lenght := tun.Section("Main").Get("lenght").(float64)
	mas := tun.Section("Main").Get("mas").([]interface{})
	first := mas[0].(int)

Конфигурационный файл, это обычный ini-файл, поддерживающий строки, целые числа, числа с плавающей точкой и списки.

Комментарии
Комментарии обозначаются символом решетки #. Всё, что следует за этим символом в строке, считается комментарием.

Строки
Строки обрамляются в одинарные либо двойные кавычки. Вложенные кавычки должны отличаться от наружных.

Целые числа
Переменная, состоящая только из цифр, трактуется как целое int. Целое число не может быть больше 2147483648

Числа с плавающей точкой
Переменная из цифр, содержащая точку (не запятую!) трактуется как float64

Списки
Списки могут содержать строки и оба вида чисел. Вложенные списки не поддерживаются. Списки обрамляются в простые либо фигурные скобки. Члены списка отделяются друг от друга запятыми.

Принцип формирования ключа
В переменных окружения и командной строки ключ переменной надо указывать в виде пары Секция_ключ, соединенных символом подрёркивания. Соединительный символ можно поменять в файле конфигурации. Строки надо также, как и в файле, обрамлять в кавычки, а списки в скобки.

Пример
Для секции Main задавать значение для переменной path нужно так: Main_path. Командная строка при этом могла бы выглядеть в Windows так:
tuner.exe -Main_path="config.ini" -Second_mas=(6,7,"abc","efj",3.14)

Получение значения
При получении значения обязательно использовать сначала метод `Section` а потом метод `Get`. Полученное значение нужно приводить к нужному типу. Списки возвращают список интерфейсов, и внутри списка вам также нужно будеть значения приводить к нужному типу.
path := tun.Section("Main").Get("path").(string)

Изменение значения
При необходимости можно изменять значение параметра конфигурации, содержимое файла конфигурации при этом не изменяется:
	tnr, err := New("config.ini")
	y := tnr.Section("Main").Set("path", "now.ini")
	if x := tnr.Section("Main").Get("path"); y != nil || x != "now.ini" {
		fmt.Print("Error in key `path`: ", x, "! ", err)
	}

*/
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>
package tuner