.PHONY: up down restart logs ps clean test db-shell

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
