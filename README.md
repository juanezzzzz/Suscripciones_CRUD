# 💳 Suscripciones API — Microservicio de Membresías (Go + PostgreSQL)

![Go](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-4169E1?style=flat&logo=postgresql&logoColor=white)
![Gorilla Mux](https://img.shields.io/badge/Gorilla_Mux-425A5B?style=flat&logo=go&logoColor=white)
![REST API](https://img.shields.io/badge/REST-API-6BA539?style=flat)

> Microservicio **REST** que administra los **planes de membresía** de la plataforma **Jobsy**, sus características y las suscripciones de cada usuario.

---

## 📖 ¿Qué es?

**Suscripciones API** es un backend REST escrito en **Go** que gestiona el modelo de **membresías** de la plataforma Jobsy. Expone **3 recursos** con operaciones **CRUD completas** sobre **PostgreSQL**, con una arquitectura modular por capas y configuración por **variables de entorno**.

## 🧩 ¿Qué hace?

- 📦 Administra los **planes** disponibles (nombre, precio mensual y anual, descripción).
- ✅ Define las **características** incluidas en cada plan (y si están o no incluidas).
- 👤 Gestiona la **suscripción de cada usuario**: qué plan tiene, su modalidad (mensual/anual), su estado (activa, cancelada, expirada), fechas de inicio/fin y renovación automática.

---

## 🏗️ Parte de la plataforma Jobsy

Este servicio es uno de los **microservicios** de **Jobsy**. Comparte la base de datos `Jobsy` en PostgreSQL, aislado en su propio **schema `suscripciones`**. Su servicio hermano:

- 🛒 [**Tienda API**](https://github.com/juanezzzzz/Tienda_CRUD) — e-commerce completo: catálogo, carritos, pagos y monedero (schema `tienda`, puerto `8083`).

---

## 🧱 Arquitectura

Estructura organizada por capas (separación de responsabilidades):

```
Suscripciones_CRUD/
├── config/          # Conexión a PostgreSQL vía variables de entorno
│   └── db.go
├── models/          # 3 structs → planes, plan_features, usuarios_suscripciones
├── controllers/     # Handlers HTTP (lógica CRUD) + helpers.go
├── routes/          # Registro de endpoints por recurso
├── main.go          # Entry point: router, CORS y arranque (puerto 8084)
├── go.mod / go.sum  # Módulo y dependencias
└── README.md
```

**Flujo de una petición:** `main.go` (router + CORS) → `routes/` (mapea la URL) → `controllers/` (procesa y responde JSON) → `models/` + `config/` (acceso a datos en PostgreSQL).

---

## 🛠️ Stack

| Tecnología | Uso |
|---|---|
| **Go** | Lenguaje del servicio |
| **Gorilla Mux** | Enrutador HTTP |
| **PostgreSQL** (`lib/pq`) | Base de datos relacional |
| **CORS middleware** | Acceso desde el frontend |

---

## 🗃️ Modelo de datos (3 recursos)

| Recurso | Endpoint base | Qué almacena |
|---|---|---|
| **Planes** | `/planes` | Planes de membresía: nombre, precio mensual, precio anual, descripción |
| **Características de plan** | `/plan-features` | Características de cada plan (texto de la feature + si está incluida) |
| **Suscripciones de usuario** | `/usuarios-suscripciones` | Suscripción por usuario: plan, modalidad, estado, inicio, fin y auto-renovación |

> 🧷 Cada entidad incluye un campo `activo` como bandera de estado (baja lógica). Campos opcionales (como `precio_anual` o `fin_en`) pueden llegar como `null`.

---

## 📡 Referencia de endpoints

**Los 3 recursos siguen el mismo contrato REST.** Sustituye `<recurso>` por `planes`, `plan-features` o `usuarios-suscripciones`:

| Método | Ruta | Acción |
|---|---|---|
| `GET` | `/<recurso>` | Listar todos |
| `GET` | `/<recurso>/{id}` | Obtener uno por ID |
| `POST` | `/<recurso>` | Crear |
| `PUT` | `/<recurso>/{id}` | Actualizar por ID |
| `DELETE` | `/<recurso>/{id}` | Eliminar por ID |

Todas las respuestas son **JSON** (`Content-Type: application/json`).

### Ejemplo — crear un plan

```http
POST http://localhost:8084/planes
Content-Type: application/json
```

```json
{
  "nombre": "Premium",
  "precio_mensual": 29900,
  "precio_anual": 299000,
  "descripcion": "Acceso completo a todas las funciones de Jobsy",
  "activo": true
}
```

---

## ⚙️ Configuración y ejecución

### Requisitos

- **Go** 1.21+
- **PostgreSQL** con la base `Jobsy` y el schema `suscripciones` creados.

### Variables de entorno

La conexión se configura por variables de entorno (con valores por defecto en [`config/db.go`](config/db.go)):

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=********
DB_NAME=Jobsy
DB_SCHEMA=suscripciones
```

> 💡 Define estas variables en tu entorno (o un archivo `.env`) para no exponer credenciales en el código.

### Ejecutar

```bash
# 1. Clonar
git clone https://github.com/juanezzzzz/Suscripciones_CRUD.git
cd Suscripciones_CRUD

# 2. Dependencias
go mod tidy

# 3. Levantar el servidor
go run main.go
```

El servidor queda escuchando en **`http://localhost:8084`** 🚀

---

## 🗺️ Roadmap

- [ ] Autenticación con **JWT** y control de acceso
- [ ] Capa **services / repositories** (separar lógica de negocio del acceso a datos)
- [ ] **Validación** de entrada y manejo de errores estandarizado
- [ ] **Dockerización** + `docker-compose` junto al resto de la plataforma
- [ ] Pruebas automatizadas

---

## 🧭 Flujo de desarrollo

Proyecto construido con **Git Flow** (28 de abril – 1 de mayo de 2026): ramas `feature/planes`, `feature/plan-features` y `feature/usuarios-suscripciones` → `develop` → `main` (estable), integrando vía Pull Requests.

---

## 👤 Autor

**Juan Esteban Valencia A.** — [@juanezzzzz](https://github.com/juanezzzzz)
