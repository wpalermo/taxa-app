package util

import "time"

//StringToDate faz o parde de uma string para uma time Time.
//O formato do go segue esses dias fixos para saber o que é dia, mes, ano, etc...
//Ex: 2006-01-02 (ano sempre sera 2006, dia 02 e mes 01)
func StringToDate(input string, layout string) time.Time {

	if layout == "" {
		layout = "02/01/2006"
	}

	returnable, err := time.Parse(layout, input)

	if err != nil {
		panic(err)
	}

	return returnable
}
