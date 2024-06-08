package paciente

import (
	"database/sql"
	"enfermeria_go/types"
	"fmt"
)

// Storage es el repositorio de usuarios
type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CrearPaciente(user types.RegisterPaciente) error {
	_, err := s.db.Exec("INSERT INTO paciente("+
		"documento, "+
		"tipo, "+
		"primer_nombre, "+
		"segundo_nombre, "+
		"primer_apellido, "+
		"segundo_apellido, "+
		"sexo, "+
		"eps, "+
		"fecha_nacimiento, "+
		"fecha_creacion, "+
		"fecha_modificacion) "+
		"VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);",
		user.Documento,
		user.Tipo,
		user.Primer_nombre,
		user.Segundo_nombre,
		user.Primer_apellido,
		user.Segundo_apellido,
		user.Sexo,
		user.Eps,
		user.Fecha_nacimiento)

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetPacientePorDocumento(documento int) (*types.Paciente, error) {
	rows, err := s.db.Query(
		"select  "+
			"tp.detalle as tipo_documento, "+
			"p.documento as documento, "+
			"p.primer_nombre, "+
			"p.segundo_nombre, "+
			"p.primer_apellido, "+
			"p.segundo_apellido, "+
			"ts.detalle as sexo, "+
			"e.nombre as eps, "+
			"p.fecha_nacimiento, "+
			"p.fecha_creacion, "+
			"p.fecha_modificacion "+
			"from paciente as p "+
			"join tipo_documento as tp on p.tipo = tp.id "+
			"join tipo_sexo as ts on p.sexo = ts.id "+
			"join eps as e on p.eps = e.nit "+
			"where p.documento = $1;",
		documento)
	if err != nil {
		return nil, err
	}

	u := new(types.Paciente)
	for rows.Next() {
		u, err = scanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.Documento == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func scanRowsIntoUser(rows *sql.Rows) (*types.Paciente, error) {
	paciente := new(types.Paciente)

	err := rows.Scan(
		&paciente.Tipo,
		&paciente.Documento,
		&paciente.Primer_nombre,
		&paciente.Segundo_nombre,
		&paciente.Primer_apellido,
		&paciente.Segundo_apellido,
		&paciente.Sexo,
		&paciente.Eps,
		&paciente.Fecha_nacimiento,
		&paciente.Fecha_creacion,
		&paciente.Fecha_modificacion,
	)
	if err != nil {
		return nil, err
	}

	return paciente, nil
}
