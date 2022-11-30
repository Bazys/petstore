module petstore

go 1.16

require (
	github.com/go-chi/chi/v5 v5.0.7 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.14.0
	golang.org/x/sync v0.1.0
	google.golang.org/grpc v1.51.0
	google.golang.org/protobuf v1.28.1
	petstore/pkg/pb v0.0.0-00010101000000-000000000000
)

replace petstore/pkg/pb => ./pkg/pb
