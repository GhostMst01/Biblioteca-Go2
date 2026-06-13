package repository

import (
	"errors"
	"gestion-libros/domain"
)

// MemoryStorage es la estructura ("clase") que almacenará los datos en memoria.
// Implementa de forma automática las dos interfaces que creamos.
type MemoryStorage struct {
	usuarios []*domain.User
	libros   []*domain.Book
}

// NewMemoryStorage es el constructor que inicializa las listas vacías.
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		usuarios: make([]*domain.User, 0),
		libros:   make([]*domain.Book, 0),
	}
}

// Implementacion de interfaz: UserOperations

// RegistrarUsuario guarda un usuario y maneja errores si la cédula ya existe.
func (m *MemoryStorage) RegistrarUsuario(user *domain.User) error {
	for _, u := range m.usuarios {
		if u.Cedula == user.Cedula {
			return errors.New("error de negocio: la cédula ya se encuentra registrada")
		}
	}
	m.usuarios = append(m.usuarios, user)
	return nil
}

// ListarUsuarios devuelve todos los usuarios guardados en el sistema.
func (m *MemoryStorage) ListarUsuarios() []*domain.User {
	return m.usuarios
}

// Implementacion de interfaz: CatalogOperations

// AgregarLibro añade un nuevo libro al catálogo en memoria.
func (m *MemoryStorage) AgregarLibro(book *domain.Book) error {
	m.libros = append(m.libros, book)
	return nil
}

// ListarCatalogo devuelve todos los libros registrados.
func (m *MemoryStorage) ListarCatalogo() []*domain.Book {
	return m.libros
}
