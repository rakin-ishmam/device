APP_NAME=my-go-app
MAIN_FILE=./app/api/main.go
DOCKER_COMPOSE_FILE=docker-compose.yaml

.PHONY: build
build:
	@echo "Building the Go application..."
	go build -o $(APP_NAME) $(MAIN_FILE)

.PHONY: docker-up
docker-up:
	@echo "Starting Docker Compose services..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) up -d

.PHONY: docker-down
docker-down:
	@echo "Stopping Docker Compose services..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) down

.PHONY: run
run: 
	@echo "Running the Go application..."
	./$(APP_NAME)