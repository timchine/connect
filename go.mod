module github.com/timchine/connect

go 1.13

require (
	github.com/coreos/etcd v3.3.20+incompatible // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/go-redis/redis/v7 v7.2.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/google/uuid v1.1.1 // indirect
	github.com/jinzhu/gorm v1.9.12
	go.etcd.io/etcd v3.3.20+incompatible // indirect
	go.uber.org/zap v1.14.1 // indirect
	google.golang.org/grpc v1.28.1 // indirect
)

replace github.com/coreos/bbolt v1.3.4 => go.etcd.io/bbolt v1.3.4
