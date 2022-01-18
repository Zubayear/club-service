module club-service

go 1.16

require (
	github.com/Microsoft/go-winio v0.5.1 // indirect
	github.com/ProtonMail/go-crypto v0.0.0-20220113124808-70ae35bab23f // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.1 // indirect
	github.com/fsnotify/fsnotify v1.5.1 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/kevinburke/ssh_config v1.1.0 // indirect
	github.com/miekg/dns v1.1.45 // indirect
	github.com/sergi/go-diff v1.2.0 // indirect
	github.com/xanzy/ssh-agent v0.3.1 // indirect
	go-micro.dev/v4 v4.5.0
	golang.org/x/crypto v0.0.0-20220112180741-5e0467b6c7ce // indirect
	golang.org/x/net v0.0.0-20220114011407-0dd24b26b47d // indirect
	golang.org/x/sys v0.0.0-20220114195835-da31bd327af9 // indirect
	golang.org/x/tools v0.1.8 // indirect
	google.golang.org/protobuf v1.27.1
	gopkg.in/jeevatkm/go-model.v1 v1.1.0 // indirect
	gorm.io/driver/mysql v1.2.3
	gorm.io/gorm v1.22.5
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
