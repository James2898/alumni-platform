# GEMINI.md — Antigravity Workspace Guidelines & Context

This file is automatically loaded for **Google Antigravity** sessions in `alumni-platform`. It defines the architecture, project conventions, commands, and rules for executing the 4-milestone roadmap.

---

## 1. Project Overview & Architecture

* **Origin**: Clean modernization of [James2898/PHP-CI-Alumni-Management-System](https://github.com/James2898/PHP-CI-Alumni-Management-System).
* **Architecture**: Decoupled 3-Tier Cloud-Native Architecture:
  * **Frontend**: React 19 + TypeScript + Vite + Tailwind CSS v4 (`frontend/`, port `5173`)
  * **Backend**: Go 1.25+ Clean Architecture with Gin + GORM (`backend/`, port `8080`)
  * **Database**: MySQL 8.0 relational schema with foreign keys and UUIDs (`database/`, port `3306`)
  * **Object Storage**: MinIO S3-compatible local server (`alumni-media` bucket, ports `9000` API, `9001` Console)
* **Master Plan**: See [`docs/alumni-platform-master-plan.md`](./docs/alumni-platform-master-plan.md) for the 4-milestone roadmap.

---

## 2. Key Architecture Standards

### S3 Presigned URLs (AWS SAA-C03 Pattern)
- Never route raw binary file uploads through backend container memory.
- The Go backend generates an authorized 15-minute AWS S3 presigned `PUT` ticket via `/api/alumni/upload-ticket`.
- The frontend uploads directly to MinIO/S3 using standard HTTP `PUT`. Only the resulting `object_key` is persisted in the database.

### Kubernetes & Container Probes (CKAD Pattern)
- The Go backend must always maintain:
  - `GET /healthz`: Liveness probe (returns HTTP 200 `{"status":"healthy"}`).
  - `GET /readyz`: Readiness probe (actively pings MySQL socket before admitting traffic).

### Clean Architecture Boundaries
- **Backend Directory Structure**:
  - `cmd/api/main.go`: Server bootstrap and router assembly.
  - `internal/domain/`: Pure data models, DTOs, and interface contracts. No database or HTTP dependencies.
  - `internal/repository/`: GORM database queries and persistence.
  - `internal/service/`: Business logic, password hashing, JWT creation.
  - `internal/handler/`: Gin HTTP request binding, status codes, JSON responses.
  - `internal/middleware/`: JWT verification and RBAC (`admin`, `alumni`).
  - `internal/storage/`: MinIO / AWS S3 client and presigned URL generator.

---

## 3. Developer Commands & Workflows

Run from repository root:

```bash
make up         # Start all 5 Docker containers (Go, React, MySQL, MinIO, Storage-Init)
make down       # Stop all containers
make logs       # Follow streaming logs across all containers
make ps         # Check health status of all containers
make db-shell   # Open MySQL interactive shell inside alumni-db
make test       # Run Go backend tests and TypeScript frontend build
make clean      # Destroy containers, networks, and persistent data volumes
```

### Local Dev Ports & URLs
* **Frontend**: `http://localhost:5173`
* **Backend API**: `http://localhost:8080`
* **Health Liveness**: `http://localhost:8080/healthz`
* **Readiness**: `http://localhost:8080/readyz`
* **MinIO Console**: `http://localhost:9001` (`minioadmin` / `minioadmin123`)
* **MySQL**: `localhost:3306` (`alumni_user` / `alumni_secret`, DB: `alumni_db`)

### Default Seed Accounts
* **Admin**: `admin@alumni.edu` / `admin123`
* **Alumni (Active)**: `jiro.martinez@alumni.edu` / `password123`
* **Alumni (Pending Verification)**: `jolan.clemente@alumni.edu` / `password123`

---

## 4. Multi-Milestone Roadmap Status

- [x] **Milestone 1**: Core Modern Platform (Go REST API + React 19 + MySQL 8.0 + MinIO + Docker Compose)
- [x] **Milestone 2**: Kubernetes Orchestration (CKAD manifests in `deploy/k8s/`, resource limits, probes, ingress)
- [-] **Milestone 3**: Infrastructure as Code (Terraform) — *Deferred / Skipped by user direction*
- [ ] **Milestone 4**: Intelligent AI Capabilities (AWS AI Practitioner: PDF resume extractor worker) — *In Progress (Branch `feat/ai/resume-extractor-milestone-four`)*

---

## 5. Kubernetes & Environment Testing Policy

* **Primary Dev Environment**: Local development is exclusively run via Docker Compose (`make up`).
* **Kubernetes Validation**: Manifests in `deploy/k8s/` are validated via client-side dry-run (`kubectl apply -k deploy/k8s/ --dry-run=client` or `make k8s-diff`).
* **Live Cluster Deferral**: Do not require or prompt for live cluster rollouts (Docker Desktop K8s / Kind / Minikube) during milestone implementation; live cluster testing is deferred to the end of development.

---

## 6. Git & Pull Request Conventions

* **PR Assignee**: Pull requests must always be assigned to the user (`--add-assignee "@me"`).


