module github.com/ThomasVonGera/shippy-service-consignment

go 1.15

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
//replace github.com/ThomasVonGera/shippy-service-vessel => ../shippy-service-vessel
require (
	github.com/ThomasVonGera/shippy-service-vessel v0.0.0-20200830143317-8564bca0d2dd
	github.com/golang/protobuf v1.4.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.5.1 // indirect
	go.mongodb.org/mongo-driver v1.4.0
	golang.org/x/crypto v0.0.0-20200709230013-948cd5f35899 // indirect
	golang.org/x/net v0.0.0-20200822124328-c89045814202 // indirect
	golang.org/x/sys v0.0.0-20200824131525-c12d262b63d8 // indirect
	google.golang.org/protobuf v1.25.0
	gopkg.in/yaml.v2 v2.2.8 // indirect
)
