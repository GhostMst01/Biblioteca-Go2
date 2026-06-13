package main

import (
	"fmt"
	"log"

	"gestion-libros/domain"
	"gestion-libros/fp"
	"gestion-libros/repository"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// CONEXIÓN A BASE DE DATOS Y PROGRAMACIÓN FUNCIONAL
	// 1. DEFINE LA CADENA DE CONEXIÓN
	dsn := "root:test@tcp(127.0.0.1:3306)/BaseBibliotecaria?charset=utf8mb4&parseTime=True&loc=Local"

	// 2. INTENTO DE CONEXIÓN (Manejo de errores exigido)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("\n[ERROR DE CONEXIÓN] No se pudo acceder. Revisa si DBeaver/MySQL está encendido.\nDetalle técnico: %v", err)
	}

	// 3. CONSULTA DE RECOLECCIONES DE DATA
	var books []domain.Book
	db.Find(&books)

	fmt.Println("\n----------------------------------------------")
	fmt.Println("PROCESAMIENTO FUNCIONAL CON DATA DE DBEAVER")
	fmt.Println("----------------------------------------------")

	// FILTER: Filtrar los libros de la base de datos que sean de la categoría "Novela"
	categoriaObjetivo := "Novela"
	novelas := fp.Filter(books, func(b domain.Book) bool {
		return b.Categoria == categoriaObjetivo
	})

	fmt.Printf("\n[Catálogo] Filtrando por categoría '%s':\n", categoriaObjetivo)
	fmt.Println("----------------------------------------------------------------------")
	fmt.Printf("| %-35s | %-25s |\n", "TÍTULO DEL LIBRO", "AÑO DE PUBLICACIÓN")
	fmt.Println("----------------------------------------------------------------------")
	for _, b := range novelas {
		fmt.Printf("| %-35s | Año: %-20s |\n", b.Titulo, b.Anio_Publicacion)
	}
	fmt.Println("----------------------------------------------------------------------")

	// REDUCE: Contar de forma funcional cuántos registros totales procesó el sistema
	totalLibros := fp.Reduce(books, 0, func(accumulator int, b domain.Book) int {
		return accumulator + 1
	})

	fmt.Printf("\n[Reportes] Registros procesados: %d libros.\n", totalLibros)
	fmt.Println("---------------------------------------------")

	// INTERFACES, ENCAPSULACIÓN Y NUEVOS MÓDULOS
	fmt.Println("\n-------------------------------------")
	fmt.Println("INTERFACES, ENCAPSULACIÓN Y MÓDULOS")
	fmt.Println("-------------------------------------")

	// 1. Instanciamos el almacenamiento que implementa de forma orientada a objetos las interfaces
	storage := repository.NewMemoryStorage()

	// ----GESTIÓN DE USUARIOS ----
	fmt.Println("\n[Módulo Usuarios] Creando registros con validación y encapsulamiento:")

	// Usamos el constructor NewUser que encapsula y valida errores
	u1, errUser := domain.NewUser(1, "1725544332", "Mateo Baca", "mateo@gmail.com", "Administrador")
	if errUser == nil {
		// Acoplamos el almacenamiento a la interfaz correspondiente
		var userInterface repository.UserOperations = storage

		// Guardamos a través de la interfaz
		userInterface.RegistrarUsuario(u1)

		// CONTROL DE ERRORES: Intentamos registrar el mismo usuario para forzar el error de duplicados
		errDuplicado := userInterface.RegistrarUsuario(u1)
		if errDuplicado != nil {
			fmt.Printf("Control de Error Capturado (Cédula Duplicada): %v\n", errDuplicado)
		}

		// Listamos los usuarios activos mediante el contrato de la interfaz
		for _, user := range userInterface.ListarUsuarios() {
			fmt.Printf("Usuario Registrado Exitosamente: %s | Rol: %s\n", user.Nombre, user.Rol)
		}
	}

	// CONTROL DE ERRORES: Intentamos forzar un error creando un usuario con datos inválidos
	_, errValidacion := domain.NewUser(2, "", "Invalido", "correoMalFormado", "Estudiante")
	if errValidacion != nil {
		fmt.Printf("Control de Error Capturado (Validación de Atributos): %v\n", errValidacion)
	}

	// ---- DEMOSTRACIÓN MÓDULO ACTUALIZADO: GESTIÓN DE CATÁLOGO ----
	fmt.Println("\n[Módulo Catálogo] Insertando registros a través de interfaces:")

	var catalogInterface repository.CatalogOperations = storage

	// Usamos el constructor controlado de libros para simular nuevas inserciones seguras
	bNew1, _ := domain.NewBook(201, "Estructuras de Datos en Go", "Ingeniería")
	catalogInterface.AgregarLibro(bNew1)

	for _, libro := range catalogInterface.ListarCatalogo() {
		fmt.Printf("Libro Añadido mediante Interfaz: %s (ID: %d)\n", libro.Titulo, libro.ID_Libro)
	}

	fmt.Println("\n--------------------")
	fmt.Println("COMPILADO CON ÉXITO")
	fmt.Println("--------------------")
}
