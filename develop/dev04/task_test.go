package main

import (
	"reflect"
	"testing"
)

func Test_anagramma(t *testing.T) {
	itogMap := make(map[string][]string)
	a1 := []string{"кот", "кто", "окт", "ток"}
	a2 := []string{"кит", "тик"}
	a3 := []string{"крона", "норка"}
	a4 := []string{"пятка", "тяпка"}
	itogMap["ток"] = a1
	itogMap["тик"] = a2
	itogMap["норка"] = a3
	itogMap["тяпка"] = a4

	t.Run("test", func(t *testing.T) {
		massiv := &[]string{"Тяпка", "пятка", "ток", "коТ", "окт", "тик", "кИт", "Яблоко", "КТО", "НОРКа", "КроНА"}
		gotMap := anagramma(massiv)
		for o, oo := range *gotMap {
			if !reflect.DeepEqual(oo, itogMap[o]) {
				t.Errorf("anagramma() = %v, want %v", gotMap, itogMap)
			}
		}
	})
}
