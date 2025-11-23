package main

import (
	"github.com/markokroselj/redovalnica-app/redovalnica"
)

func main() {
	var students = make(map[string]redovalnica.Student)
	students["63210001"] = redovalnica.Student{Ime: "Janez", Priimek: "Novak",Ocene: []int{6, 7, 8, 7}}
	students["63210002"] = redovalnica.Student{Ime: "Ana",Priimek: "Kralj", Ocene: []int{6, 7, 8, 7, 10, 8, 7}}
	students["63210003"] = redovalnica.Student{Ime: "Franc", Priimek: "Novak", Ocene: []int{10, 7, 8, 10, 10, 8, 7}}
	students["63210004"] = redovalnica.Student{Ime: "Micka",Priimek: "Kovač", Ocene: []int{10, 10, 8, 10, 10, 8, 10}}

	students["63210007"] = redovalnica.Student{Ime: "Marko", Priimek:  "Krošelj", Ocene: []int{10, 7, 8, 9, 10, 7}}
	redovalnica.DodajOceno(students, "63210007", 7)
	redovalnica.DodajOceno(students, "63210007", 25)
	redovalnica.DodajOceno(students, "63210007", -1)
	redovalnica.DodajOceno(students, "123", 8)

	redovalnica.DodajOceno(students, "63210002", 7)

	redovalnica.IzpisRedovalnice(students)

	redovalnica.IzpisiKoncniUspeh((students))
}
