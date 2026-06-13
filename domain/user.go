package domain

import (
	"errors"
	"strings"
)

// User representa el modelo para el nuevo módulo de Gestión de Usuarios
type User struct {
	ID     uint   `json:"id"`
	Cedula string `json:"cedula"`
	Nombre string `json:"nombre"`
	Correo string `json:"correo"`
	Rol    string `json:"rol"` // Administrador, Bibliotecario, Estudiante
}

// NewUser es el constructor del módulo que valida la integridad de los datos
func NewUser(id uint, cedula, nombre, correo, rol string) (*User, error) {
	// Validación de campos obligatorios (Manejo de Errores)
	if strings.TrimSpace(cedula) == "" || strings.TrimSpace(nombre) == "" {
		return nil, errors.New("la cédula y el nombre son campos obligatorios")
	}
	// Validación de formato de negocio
	if !strings.Contains(correo, "@") {
		return nil, errors.New("el correo electrónico provisto no es válido")
	}

	return &User{
		ID:     id,
		Cedula: cedula,
		Nombre: nombre,
		Correo: correo,
		Rol:    rol,
	}, nil
}
