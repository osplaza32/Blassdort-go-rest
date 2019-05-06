package Tablas

import "time"

var Usuario struct{
	ID string `gorm:"primary_key;uuid"`
	Tarjetaid string
	Nombre string
	Email string
	Apellido string
	Cargo string
	Foto string
	Telefono string
	Role string
	HashPassword string
	Pincode string
	CreatedAt time.Time `sql:"DEFAULT:current_timestamp"`
	EditAt time.Time


}