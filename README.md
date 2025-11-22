# 🎓 Sistema de Control Escolar - API REST

![Go Version](https://img.shields.io/badge/Go-1.25-00ADD8?style=flat&logo=go)
![Gin](https://img.shields.io/badge/Gin-v1.11-00ADD8?style=flat)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-14+-336791?style=flat&logo=postgresql)
![GORM](https://img.shields.io/badge/GORM-v1.31-00ADD8?style=flat)

API RESTful desarrollada en **Go (Golang)** utilizando el framework **Gin Gonic** y **GORM** como ORM. Sistema completo de gestión escolar que permite administrar estudiantes, materias y calificaciones con persistencia de datos en **PostgreSQL**.

## 📋 Tabla de Contenidos

- [Características](#-características)
- [Tecnologías Utilizadas](#-tecnologías-utilizadas)
- [Estructura del Proyecto](#-estructura-del-proyecto)
- [Requisitos Previos](#-requisitos-previos)
- [Instalación](#-instalación)
- [Configuración](#-configuración)
- [Ejecución](#-ejecución)
- [Endpoints de la API](#-endpoints-de-la-api)
- [Modelos de Datos](#-modelos-de-datos)
- [Ejemplos de Uso](#-ejemplos-de-uso)
- [Validaciones](#-validaciones)
- [Manejo de Errores](#-manejo-de-errores)

## ✨ Características

- ✅ **CRUD completo** para Estudiantes, Materias y Calificaciones
- ✅ **Validaciones automáticas** con Gin Validator
- ✅ **Validación de integridad referencial** (Foreign Keys)
- ✅ **Validación de rango de calificaciones** (0-10)
- ✅ **Validación de email único** para estudiantes
- ✅ **Relaciones entre entidades** (Student → Grades ← Subject)
- ✅ **Consultas con JOIN** usando GORM Preload
- ✅ **Migraciones automáticas** de base de datos
- ✅ **Manejo de errores robusto** con códigos HTTP apropiados
- ✅ **API RESTful** siguiendo mejores prácticas

## 🛠 Tecnologías Utilizadas

| Tecnología | Versión | Propósito |
|------------|---------|-----------|
| **Go** | 1.25+ | Lenguaje de programación |
| **Gin Gonic** | 1.11.0 | Framework web HTTP |
| **GORM** | 1.31.1 | ORM para Go |
| **PostgreSQL** | 14+ | Base de datos relacional |
| **Gin Validator** | 10.28.0 | Validación de datos |

## 📁 Estructura del Proyecto

```
Control_Escolar/
│
├── config/
│   └── db.go                 # Configuración de base de datos
│
├── models/
│   ├── student.go            # Modelo de Estudiante
│   ├── subject.go            # Modelo de Materia
│   └── grade.go              # Modelo de Calificación
│
├── controllers/
│   ├── student_controller.go # Controladores de Estudiantes
│   ├── subject_controller.go # Controladores de Materias
│   └── grade_controller.go   # Controladores de Calificaciones
│
├── routes/
│   └── routes.go             # Definición de rutas de la API
│
├── .env                      # Variables de entorno (no incluido en Git)
├── .gitignore                # Archivos ignorados por Git
├── main.go                   # Punto de entrada de la aplicación
├── go.mod                    # Dependencias del proyecto
├── go.sum                    # Checksums de dependencias
└── README.md                 # Este archivo
```

## 📋 Requisitos Previos

Antes de comenzar, asegúrate de tener instalado:

- **Go** versión 1.20 o superior → [Descargar Go](https://golang.org/dl/)
- **PostgreSQL** versión 14 o superior → [Descargar PostgreSQL](https://www.postgresql.org/download/)
- **Git** → [Descargar Git](https://git-scm.com/downloads)
- Un cliente HTTP como **Postman** o **Thunder Client** (opcional)

## 🚀 Instalación

### 1. Clonar el repositorio

```bash
git clone https://github.com/IanB28/Control_Escolar
cd Control_Escolar
```

### 2. Instalar dependencias

```bash
go mod download
go mod tidy
```

### 3. Verificar instalación

```bash
go version
```

## ⚙️ Configuración

### 1. Configurar PostgreSQL

Inicia PostgreSQL y crea la base de datos:

```sql
CREATE DATABASE control_escolar;
```

### 2. Configurar variables de entorno

Crea un archivo `.env` en la raíz del proyecto:

```env
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=
DB_NAME=control_escolar
DB_PORT=5432
DB_SSLMODE=disable
```

### 3. Configurar conexión a la base de datos

Actualiza el archivo `config/db.go` para usar las variables de entorno:

```go
// Recomendación: Usa github.com/joho/godotenv para cargar .env
dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
    os.Getenv("DB_HOST"),
    os.Getenv("DB_USER"),
    os.Getenv("DB_PASSWORD"),
    os.Getenv("DB_NAME"),
    os.Getenv("DB_PORT"),
    os.Getenv("DB_SSLMODE"),
)
```

## ▶️ Ejecución

### Modo desarrollo

```bash
go run main.go
```

### Compilar y ejecutar

```bash
go build -o control_escolar.exe
./control_escolar.exe
```

El servidor iniciará en: **http://localhost:8081**

Verás el mensaje:
```
Conexion a Base de Datos Exitosa
Tablas migradas correctamente
[GIN-debug] Listening and serving HTTP on :8081
```

## 🔌 Endpoints de la API

### 📚 Estudiantes

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| `GET` | `/api/students` | Obtener todos los estudiantes |
| `GET` | `/api/students/:id` | Obtener un estudiante por ID |
| `POST` | `/api/students` | Crear un nuevo estudiante |
| `PUT` | `/api/students/:id` | Actualizar un estudiante |
| `DELETE` | `/api/students/:id` | Eliminar un estudiante |

### 📖 Materias

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| `GET` | `/api/subjects/:id` | Obtener una materia por ID |
| `POST` | `/api/subjects` | Crear una nueva materia |
| `PUT` | `/api/subjects/:id` | Actualizar una materia |
| `DELETE` | `/api/subjects/:id` | Eliminar una materia |

### 📊 Calificaciones

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| `POST` | `/api/grades` | Registrar una calificación |
| `PUT` | `/api/grades/:id` | Actualizar una calificación |
| `DELETE` | `/api/grades/:id` | Eliminar una calificación |
| `GET` | `/api/grades/student/:student_id` | Obtener todas las calificaciones de un estudiante (con detalles de materia) |
| `GET` | `/api/grades/:id/student/:student_id` | Obtener una calificación específica de un estudiante |

## 📊 Modelos de Datos

### Student (Estudiante)

```json
{
  "student_id": 1,
  "name": "Jake Buzzo",
  "group": "A",
  "email": "jake@gmail.com"
}
```

**Validaciones:**
- `name`: Requerido
- `group`: Requerido
- `email`: Requerido, formato email válido, único

### Subject (Materia)

```json
{
  "subject_id": 1,
  "name": "Matematicas"
}
```

**Validaciones:**
- `name`: Requerido

### Grade (Calificación)

```json
{
  "grade_id": 1,
  "student_id": 1,
  "subject_id": 1,
  "grade": 9.5
}
```

**Validaciones:**
- `student_id`: Requerido, debe existir en la tabla students
- `subject_id`: Requerido, debe existir en la tabla subjects
- `grade`: Requerido, rango válido 0-10

### Relaciones

```
Student (1) ──────< (N) Grade (N) >────── (1) Subject
         uno a muchos              muchos a uno
```

## 📝 Ejemplos de Uso

### Crear un Estudiante

**Request:**
```http
POST http://localhost:8081/api/students
Content-Type: application/json

{
  "name": "Juan Pérez",
  "group": "A",
  "email": "juan@test.com"
}
```

**Response (201 Created):**
```json
{
  "student_id": 1,
  "name": "Juan Pérez",
  "group": "A",
  "email": "juan@test.com"
}
```

### Crear una Materia

**Request:**
```http
POST http://localhost:8081/api/subjects
Content-Type: application/json

{
  "name": "Matemáticas"
}
```

**Response (201 Created):**
```json
{
  "subject_id": 1,
  "name": "Matemáticas"
}
```

### Registrar una Calificación

**Request:**
```http
POST http://localhost:8081/api/grades
Content-Type: application/json

{
  "student_id": 1,
  "subject_id": 1,
  "grade": 9.5
}
```

**Response (201 Created):**
```json
{
  "grade_id": 1,
  "student_id": 1,
  "subject_id": 1,
  "grade": 9.5
}
```

### Obtener Calificaciones de un Estudiante

**Request:**
```http
GET http://localhost:8081/api/grades/student/1
```

**Response (200 OK):**
```json
[
  {
    "grade_id": 1,
    "student_id": 1,
    "subject_id": 1,
    "grade": 9.5,
    "Student": {
      "student_id": 1,
      "name": "Jake Buzzo",
      "group": "A",
      "email": "jake@gmail.com"
    },
    "Subject": {
      "subject_id": 1,
      "name": "Matemáticas"
    }
  },
  {
    "grade_id": 2,
    "student_id": 1,
    "subject_id": 2,
    "grade": 8.7,
    "Student": {
      "student_id": 1,
      "name": "Jake Buzzo",
      "group": "A",
      "email": "jake@gmail.com"
    },
    "Subject": {
      "subject_id": 2,
      "name": "Geografia"
    }
  }
]
```

### Actualizar un Estudiante

**Request:**
```http
PUT http://localhost:8081/api/students/1
Content-Type: application/json

{
  "name": "Ian Buzzo",
  "group": "B",
  "email": "ian@gmail.com"
}
```

**Response (200 OK):**
```json
{
  "student_id": 1,
  "name": "Ian Buzzo",
  "group": "B",
  "email": "ian@gmail.com"
}
```

### Eliminar una Calificación

**Request:**
```http
DELETE http://localhost:8081/api/grades/1
```

**Response (200 OK):**
```json
{
  "message": "Calificación eliminada"
}
```

## ✅ Validaciones

El sistema implementa las siguientes validaciones:

### Validaciones de Estudiante
- ✅ **Nombre**: Campo requerido
- ✅ **Grupo**: Campo requerido
- ✅ **Email**: Campo requerido, formato válido, único en la base de datos

### Validaciones de Materia
- ✅ **Nombre**: Campo requerido

### Validaciones de Calificación
- ✅ **Student ID**: Campo requerido, debe existir en la tabla students
- ✅ **Subject ID**: Campo requerido, debe existir en la tabla subjects
- ✅ **Calificación**: Campo requerido, rango válido entre 0 y 10

### Validación de Integridad Referencial
El sistema valida que los IDs de estudiantes y materias existan antes de crear/actualizar calificaciones.

## ⚠️ Manejo de Errores

### Códigos de Estado HTTP

| Código | Significado | Ejemplo |
|--------|-------------|---------|
| `200 OK` | Operación exitosa | GET, PUT, DELETE exitosos |
| `201 Created` | Recurso creado | POST exitoso |
| `400 Bad Request` | Error de validación | JSON inválido, campos faltantes |
| `404 Not Found` | Recurso no encontrado | ID no existe |
| `500 Internal Server Error` | Error del servidor | Error de base de datos |

### Ejemplos de Errores

**Email duplicado (400):**
```json
{
  "error": "UNIQUE constraint failed: students.email"
}
```

**Calificación fuera de rango (400):**
```json
{
  "error": "Key: 'Grade.Grade' Error:Field validation for 'Grade' failed on the 'lte' tag"
}
```

**Estudiante no encontrado (400):**
```json
{
  "error": "El ID de estudiante no existe"
}
```

**Recurso no encontrado (404):**
```json
{
  "error": "Estudiante no encontrado"
}
```

## 🔒 Seguridad

### Recomendaciones implementadas:
- ✅ Credenciales en archivo `.env` (no incluido en Git)
- ✅ Validación de entrada de datos
- ✅ Uso de ORM para prevenir SQL Injection

### TODO - Mejoras de seguridad:
- [ ] Implementar autenticación JWT
- [ ] Rate limiting
- [ ] CORS configurado
- [ ] HTTPS en producción

## 🧪 Testing

Para ejecutar pruebas (próximamente):

```bash
go test ./...
```

## 📦 Despliegue

### Variables de entorno requeridas en producción:
- `DB_HOST`
- `DB_USER`
- `DB_PASSWORD`
- `DB_NAME`
- `DB_PORT`
- `DB_SSLMODE`

### Compilar para producción:

```bash
# Windows
GOOS=windows GOARCH=amd64 go build -o control_escolar.exe

# Linux
GOOS=linux GOARCH=amd64 go build -o control_escolar

# macOS
GOOS=darwin GOARCH=amd64 go build -o control_escolar
```

Proyecto Final, Taller Fundamentos de Go.


---
