package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargoHoraria int
}

type Cursos []Curso

func main() {
	t := template.Must(template.New("template.html").ParseFiles("template.html"))
	err := t.Execute(os.Stdout, Cursos{
		{"GO", 40},
		{"Java", 60},
		{"Python", 10},
	})
	if err != nil {
		panic(err)
	}
}
