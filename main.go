package main

import (
	"github.com/AugustoPersonalProjects/gin-api-rest/models"
	"github.com/AugustoPersonalProjects/gin-api-rest/routes"
)

func main() {
	models.Alunos = []models.Aluno{
		{Nome: "Augusto Lopes", CPF: "111.222.333-45", RG: "10.234.567-8"},
		{Nome: "Franco Baresi", CPF: "120.340.450-68", RG: "20.345.678.9"},
	}
	routes.HandleRequests()
}
