protoc -Igreet/proto --go_opt=module=github.com/smart-think-app/huy_grpc --go_out=. --go-grpc_opt=module=github.com/sm
art-think-app/huy_grpc --go-grpc_out=. greet/proto/*.proto