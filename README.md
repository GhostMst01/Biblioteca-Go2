# Biblioteca-Go2
## Datos del Estudiante
* **Nombre:** Mateo Steven Baca Cedeño
* **Carrera:** Ingeniería en Sistemas
* **Asignatura:** Programación Orientada a Objetos
* **Fecha:**  13 Junio 2026

## Objetivo del Programa
El objetivo principal de este sistema es automatizar la administración de una biblioteca universitaria, permitiendo la gestión eficiente del catálogo de libros y el registro de usuarios. El enfoque de esta entrega es aplicar arquitecturas de software limpias mediante el uso de interfaces, encapsulación y un control estricto de las reglas de negocio y manejo de errores en el lenguaje Go.

## Principales Funcionalidades del Código
El sistema está diseñado bajo una arquitectura modular y desacoplada que incluye:

1. **Gestión de Catálogo (Módulo Libros):**
   * Conexión funcional y persistencia real con la base de datos MariaDB a través del ORM GORM.
   * Recuperación automatizada de registros de la tabla `LIBRO` en DBeaver.
   * Procesamiento de datos avanzado utilizando programación funcional (filtros por categoría y cálculo de reportes totales).

2. **Gestión de Usuarios (Módulo de Simulación Seguro):**
   * Implementación de la estructura `User` con encapsulación mediante funciones constructoras (`NewUser`).
   * Validación estricta de **Reglas de Negocio** en el dominio (control de campos obligatorios como cédula y nombre, y formato de correo electrónico).
   * Sistema avanzado de **Manejo de Errores** nativos para evitar el ingreso de datos duplicados o corruptos.

3. **Arquitectura basada en Interfaces:**
   * Uso de los contratos abstractos `UserOperations` y `CatalogOperations` para independizar la lógica de la aplicación del almacenamiento final.
   * Implementación de un **Escenario de Pruebas en Memoria** (`memory.go`) para validar el comportamiento del módulo de usuarios de forma segura y aislada de la base de datos de producción, garantizando la escalabilidad total para la integración de servicios Web en la Unidad 4.

## Estructura del Proyecto
El código se encuentra organizado bajo las mejores prácticas de estructuración en Go:
* `/domain`: Contiene las entidades del negocio y sus reglas de validación (`user.go`).
* `/repository`: Contiene las definiciones de las interfaces (`interfaces.go`) y la simulación de almacenamiento en memoria RAM (`memory.go`).
* `main.go`: Punto de entrada que coordina el flujo del sistema e interactúa con los componentes únicamente a través de los contratos abstractos.
