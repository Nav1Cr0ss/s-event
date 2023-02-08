module github.com/Nav1Cr0ss/s-event

go 1.19

replace github.com/Nav1Cr0ss/s-lib => ./pkg/s-lib

require (
	github.com/Nav1Cr0ss/s-lib v0.0.0-20230207083835-67bd4cdb95e4
	github.com/envoyproxy/protoc-gen-validate v0.9.1
	github.com/lib/pq v1.10.7
	go.uber.org/fx v1.19.1
	google.golang.org/genproto v0.0.0-20230202175211-008b39050e57
	google.golang.org/grpc v1.52.3
	google.golang.org/protobuf v1.28.1
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/kelseyhightower/envconfig v1.4.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/dig v1.16.0 // indirect
	go.uber.org/multierr v1.9.0 // indirect
	go.uber.org/zap v1.24.0 // indirect
	golang.org/x/net v0.5.0 // indirect
	golang.org/x/sys v0.4.0 // indirect
	golang.org/x/text v0.6.0 // indirect
)
