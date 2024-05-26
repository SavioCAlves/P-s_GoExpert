package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargoHoraria int
}

func main() {
	curso := Curso{"Go", 40}
	tmp := template.New("CursoTemplate")
	tmp, _ = tmp.Parse("Curso: {{.Nome}} - Cargo Horária: {{.CargoHoraria}} horas.\n")
	err := tmp.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}
