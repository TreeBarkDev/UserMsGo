PROTO_DIR=proto
OUT_DIR=api
REPO=github.com/TreeBarkDev/UserMsGo

proto-gen:
	protoc -I=$(PROTO_DIR) \
		--go_out=$(OUT_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative \
		--go_opt=Mproto/userapi/user-api.proto=$(REPO)/$(OUT_DIR)/userpb \
		--go-grpc_opt=Mproto/userapi/user-api.proto=$(REPO)/$(OUT_DIR)/userpb \
		$(PROTO_DIR)/proto/userapi/*.proto
