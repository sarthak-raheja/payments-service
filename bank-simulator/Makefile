.PHONY: docker-up
docker-up:
	docker-compose -f docker-compose.yaml up --build

.PHONY: docker-down
docker-down: ## Stop docker containers and clear artefacts.
	docker-compose -f docker-compose.yaml down
	docker system prune 

.PHONY: codegen.protos
codegen.protos:
	protoc --go_out=protos/v1/. --go-grpc_out=protos/v1/. protos/v1/bank_simulator.proto	