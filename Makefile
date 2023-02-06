.PHONY: rebuild
rebuild:
	docker build -t ngoctd/ecommerce-notify:latest . && \
	docker push ngoctd/ecommerce-notify

.PHONY: redeploy
redeploy:
	kubectl rollout restart deployment depl-notify

.PHONY: protogen
protogen:
	protoc --proto_path=proto proto/notify_service.proto proto/general.proto \
	--go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative

.PHONY: redis
redis:
	docker run --name redis -p 6379:6379 -d redis:7-alpine