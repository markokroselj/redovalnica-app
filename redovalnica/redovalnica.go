package redovalnica

import "fmt"

type Student struct {
	Ime     string
	Priimek string
	Ocene   []int
}

func DodajOceno(studenti map[string]Student, vpisnaStevilka string, ocena int) {
	if ocena < 0 || ocena > 10 {
		fmt.Printf("Ocena %d ni v ustreznem območju [0-10]!\n", ocena)
		return
	}

	s, ok := studenti[vpisnaStevilka]

	if !ok {
		fmt.Printf("Študenta z vpisno številko %s ni na seznamu!\n", vpisnaStevilka)
		return
	}

	s.Ocene = append(studenti[vpisnaStevilka].Ocene, ocena)
	studenti[vpisnaStevilka] = s
}

func povprecje(studenti map[string]Student, vpisnaStevilka string) float64 {
	student, ok := studenti[vpisnaStevilka]

	if !ok {
		return -1.0
	}

	if len(student.Ocene) < 6 {
		return 0.0
	}

	var sum float64
	for _, ocena := range student.Ocene {
		sum += float64(ocena)
	}

	return sum / float64(len(student.Ocene))
}

func IzpisRedovalnice(studenti map[string]Student) {
	fmt.Println("\nREDOVALNICA")

	for vpisna, student := range studenti {
		fmt.Printf("%s - %s %s: %v\n", vpisna, student.Ime, student.Priimek, student.Ocene)
	}

	fmt.Println()
}

func IzpisiKoncniUspeh(studenti map[string]Student) {
	for vpisna, student := range studenti {

		avg := povprecje(studenti, vpisna)

		var komentar string
		if avg >= 9 {
			komentar = "Odličen študent!"
		} else if avg >= 6 && avg < 9 {
			komentar = "Povprečen študent!"
		} else {
			komentar = "Neuspešen študent!"
		}

		fmt.Printf("%s %s: povprečna ocena %f -> %s\n", student.Ime, student.Priimek, avg, komentar)

	}
}