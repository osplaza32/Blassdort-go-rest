package main

import (
	"fmt"
	"osplaza32/ApiFastToTest/config"
	"osplaza32/Blassdort-go-rest/Entidades/Tablas"
)

func main() {
	db,err := Config.Conneccion()
	if err != nil {
		panic(fmt.Sprintf("No error should happen when connect database, but got %+v", err))
	}
	db.Exec("CREATE EXTENSION postgis")
	db.Exec("CREATE EXTENSION postgis_topology")
	db.DropTableIfExists(&Tablas.Usuario{},&Tablas.Hardware{},&Tablas.Empresa{},&Tablas.Plan{},&Tablas.EmpresaPlan{})
	db.AutoMigrate(&Tablas.Usuario{},&Tablas.Hardware{},&Tablas.Empresa{},&Tablas.Plan{},&Tablas.EmpresaPlan{})



}
