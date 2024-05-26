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
	t := template.Must(template.New("CursoTemplate").Parse("Curso: {{.Nome}} - Cargo Horária: {{.CargoHoraria}} horas.\n"))
	err := t.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}
