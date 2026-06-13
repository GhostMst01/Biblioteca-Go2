package repository

import "gestion-libros/domain"

// UserOperations es la interfaz que define qué acciones se pueden hacer
// en el módulo de gestión de usuarios.
type UserOperations interface {
	RegistrarUsuario(user *domain.User) error
	ListarUsuarios() []*domain.User
}

// CatalogOperations es la interfaz que define qué acciones se pueden hacer
// en el módulo de catálogo de libros.
type CatalogOperations interface {
	AgregarLibro(book *domain.Book) error
	ListarCatalogo() []*domain.Book
}
