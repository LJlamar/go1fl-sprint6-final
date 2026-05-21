package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func TextType(str string) (res string, err error) {

	//ПРОВЕРЯЕМ НА НАЛИЧИЕ СИМВОЛОВ ВНЕ ТАБЛИЦЫ ПЕРЕВОДА
	//В ВЕРХНИЙ РЕГИСТР

	//наглядная проверка
	//fmt.Println(str)
	str = strings.ToUpper(strings.TrimRight(str, "\r\n"))
	runes := []rune(str)
	//ПРОВЕРЯЕМ СИМВОЛЫ СТРОКИ

	for _, chk := range runes {
		_, ok := morse.DefaultMorse[chk]

		//ПРОВЕРКА НА СООТВЕТСТВИЕ СИМВОЛА В ТЕКСТЕ СИМВОЛАМ СЛОВАРЯ
		//ТАКЖЕ, НА НАЛИЧИЕ ПРОБЕЛА, КОТОРЫЙ ОТЧЕГО-ТО В СЛОВАРЕ НЕ ПРОПИСАН,
		//ВОЗМОЖНО ОТТОГО ЧТО НЕ ТРЕБУЕТ ПЕРЕВОДА
		if !ok && chk != 32 {

			//наглядная проверка
			//fmt.Println(chk)

			return "", errors.New("ошибка: строка для разбора содержит символы, отсутствующие в словаре")
		}
	}

	if len(str) == 0 {
		return "", errors.New("ошибка: строка для разбора не содержит текста")
	}
	if strings.Contains(str, "......") || strings.Contains(str, ".-") || strings.Contains(str, "-.") || strings.Contains(str, "- .") || strings.Contains(str, ". -") {
		//fmt.Printf("Это строка Морзе: %s\n", s)
		res := morse.ToText(str)
		//fmt.Printf("Это её перевод: %s\n", res)
		return res, err
	} else {
		//fmt.Printf("Это строка текста: %s\n", s)
		res := morse.ToMorse(str)
		//fmt.Printf("Это её перевод: %s\n", res)
		return res, err
	}
}
