.PHONY: up down restart logs ps clean test db-shell k8s-build k8s-apply k8s-delete k8s-diff

# Boot entire local platform via Docker Compose
up:
	docker compose up -d --build

# Stop all containers
down:
	docker compose down

# Restart containers
restart:
	docker compose restart

# View streaming logs
logs:
	docker compose logs -f

# List container status
ps:
	docker compose ps

# Access MySQL interactive shell inside container
db-shell:
	docker compose exec alumni-db mysql -u alumni_user -palumni_secret alumni_db

# Run tests
test:
	cd backend && go test -v ./...
	cd frontend && npm run build

# Remove containers, volumes, and caches
clean:
	docker compose down -v --remove-orphans

# Kubernetes targets
k8s-build:
	docker build -t alumni-api:latest backend --target prod
	docker build -t alumni-web:latest frontend --target prod

k8s-diff:
	kubectl kustomize deploy/k8s

k8s-apply:
	kubectl apply -k deploy/k8s

k8s-delete:
	kubectl delete -k deploy/k8s

