.PHONY: docker-up
docker-up:
	docker-compose -f docker-compose.yaml up --build

.PHONY: docker-down
docker-down: ## Stop docker containers and clear artefacts.
	docker-compose -f docker-compose.yaml down
	docker system prune 

.PHONY: grpcui.local
grpcui.local: ## Spin up GRPCUI
	grpcui -plaintext localhost:9000

.PHONY: codegen.protos
codegen.protos:
	protoc --go_out=api/v1/. --go-grpc_out=api/v1/. api/v1/service.proto	
