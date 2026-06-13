package domain

import "errors"

// Book representa la estructura de un libro mapeada con DBeaver/MariaDB
type Book struct {
	ID_Libro         uint   `gorm:"column:ID_Libro;primaryKey" json:"id_libro"`
	Titulo           string `gorm:"column:Titulo" json:"titulo"`
	Categoria        string `gorm:"column:Categoria" json:"categoria"`
	Anio_Publicacion string `gorm:"column:Anio_Publicacion" json:"anio_publicacion"`
	ID_Editorial     uint   `gorm:"column:ID_Editorial" json:"id_editorial"`
}

// TableName le indica a GORM que busque la tabla exactamente como "LIBRO"
func (Book) TableName() string {
	return "LIBRO"
}

//ENCAPSULACIÓN Y ERRORES
// NewBook actúa como constructor controlado. Aplica la regla de negocio:
// No se pueden registrar libros si los campos principales están vacíos.
func NewBook(id uint, titulo, categoria string) (*Book, error) {
	if titulo == "" || categoria == "" {
		// Retorna un error si no se cumplen los requisitos del dominio
		return nil, errors.New("el título y la categoría no pueden estar vacíos")
	}
	return &Book{
		ID_Libro:  id,
		Titulo:    titulo,
		Categoria: categoria,
	}, nil
}
