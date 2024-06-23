package types

import "time"

type Paciente struct {
	Documento          int       `json:"documento"`
	Tipo               string    `json:"tipo"`
	Primer_nombre      string    `json:"primer_nombre"`
	Segundo_nombre     string    `json:"segundo_nombre"`
	Primer_apellido    string    `json:"primer_apellido"`
	Segundo_apellido   string    `json:"segundo_apellido"`
	Sexo               string    `json:"sexo"`
	Eps                string    `json:"eps"`
	Fecha_nacimiento   time.Time `json:"fecha_nacimiento"`
	Fecha_creacion     time.Time `json:"fecha_creacion"`
	Fecha_modificacion time.Time `json:"fecha_modificacion"`
}

type RegisterPaciente struct {
	Documento        int       `json:"documento"`
	Tipo             int       `json:"tipo"`
	Primer_nombre    string    `json:"primer_nombre"`
	Segundo_nombre   string    `json:"segundo_nombre"`
	Primer_apellido  string    `json:"primer_apellido"`
	Segundo_apellido string    `json:"segundo_apellido"`
	Sexo             int       `json:"sexo"`
	Eps              int       `json:"eps"`
	Fecha_nacimiento time.Time `json:"fecha_nacimiento"`
}

type PacienteStore interface {
	GetPacientePorDocumento(int) (*Paciente, error)
	CrearPaciente(RegisterPaciente) error
	UpdatePaciente(RegisterPaciente) error
	DeltePaciente(int) error
}
