package Tablas

import "time"

type Hardware struct{
	Model
	Serial string
	FechaInstalacion time.Time
	TheLastRevision time.Time
}
