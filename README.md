# 📋 API CRUD - Suscripciones

API REST desarrollada en Go para la gestión de planes y suscripciones de usuarios. Permite manejar planes, características de planes y suscripciones activas de usuarios.

---

## 📂 Estructura del proyecto

El proyecto sigue una arquitectura modular:

- `config/` → configuración de base de datos
- `models/` → definición de entidades (planes, plan_features, usuarios_suscripciones)
- `controllers/` → lógica de negocio y respuestas HTTP
- `routes/` → definición de endpoints
- `main.go` → punto de entrada de la aplicación

---

## 🛠️ Tecnologías utilizadas

- Go (Golang)
- Gorilla Mux
- PostgreSQL
- Git & GitHub

---

## 🔐 Variables de entorno

Configura las siguientes variables antes de ejecutar el proyecto:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=tu_password
DB_NAME=Jobsy
DB_SCHEMA=suscripciones
```

---

## ▶️ Ejecución del proyecto

1. Clonar el repositorio:

```bash
git clone https://github.com/tu_usuario/suscripciones-api.git
cd suscripciones-api
```

2. Instalar dependencias:

```bash
go mod tidy
```

3. Ejecutar el servidor:

```bash
go run main.go
```

El servidor levanta en el puerto **8084**.

---

## 📡 Endpoints principales

### 📦 Planes

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| GET | `/planes` | Listar todos los planes |
| GET | `/planes/{id}` | Obtener un plan por ID |
| POST | `/planes` | Crear un nuevo plan |
| PUT | `/planes/{id}` | Actualizar un plan |
| DELETE | `/planes/{id}` | Eliminar un plan |

### ✅ Características de planes

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| GET | `/plan-features` | Listar todas las características |
| GET | `/plan-features/{id}` | Obtener una característica por ID |
| POST | `/plan-features` | Crear una característica |
| PUT | `/plan-features/{id}` | Actualizar una característica |
| DELETE | `/plan-features/{id}` | Eliminar una característica |

### 👤 Suscripciones de usuarios

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| GET | `/usuarios-suscripciones` | Listar todas las suscripciones |
| GET | `/usuarios-suscripciones/{id}` | Obtener una suscripción por ID |
| POST | `/usuarios-suscripciones` | Crear una suscripción |
| PUT | `/usuarios-suscripciones/{id}` | Actualizar una suscripción |
| DELETE | `/usuarios-suscripciones/{id}` | Eliminar una suscripción |

---

## 📅 Desarrollo

El proyecto fue desarrollado entre el **28 de abril y el 1 de mayo de 2026**, utilizando un flujo de trabajo basado en Git Flow:

- `develop` → rama de integración principal
- `feature/planes` → implementación del módulo de planes
- `feature/plan-features` → implementación de características de planes
- `feature/usuarios-suscripciones` → implementación de suscripciones
- `main` → versión final lista para producción

---

## 👨‍💻 Autor

Juan Valenciagit 
