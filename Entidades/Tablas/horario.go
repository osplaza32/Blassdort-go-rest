package Tablas

import "time"

type Horario struct{
	Model
	Inicio time.Time
	Fin time.Time
}
