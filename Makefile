COMPOSE_VALIDATOR := docker compose --env-file .env -f deploy/docker-compose-validator.yml
COMPOSE_MINER     := docker compose -f deploy/docker-compose-miner.yml
BIN_DIR           := bin

.PHONY: build-cli install-cli test clean \
        up-validator down-validator up-miner down-miner

build-cli:
	go build -o $(BIN_DIR)/sn4 ./cmd/cli

install-cli:
	go build -o $(shell go env GOPATH)/bin/sn4 ./cmd/cli

test:
	go test ./...

fmt:
	gofmt -w cmd internal

vet:
	go vet ./...

lint: fmt vet

clean:
	rm -rf $(BIN_DIR)

up-validator:
	$(COMPOSE_VALIDATOR) up -d --build --force-recreate

down-validator:
	$(COMPOSE_VALIDATOR) down --remove-orphans

up-miner:
	$(COMPOSE_MINER) up -d --build --force-recreate
	$(COMPOSE_MINER) logs -f

down-miner:
	$(COMPOSE_MINER) down --remove-orphans
