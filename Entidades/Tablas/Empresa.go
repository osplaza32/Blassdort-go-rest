package Tablas
var Empresa struct{
	ID string `gorm:"primary_key;uuid"`
	Nombre string
	Plan string

}