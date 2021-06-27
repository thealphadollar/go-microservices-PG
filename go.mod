module github.com/thealphadollar/go-microservices-PG

go 1.16

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/bradfitz/gomemcache v0.0.0-20190913173617-a41fca850d0b // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gogo/protobuf v1.2.2-0.20190723190241-65acae22fc9d // indirect
	github.com/golang/protobuf v1.5.2
	github.com/jmoiron/sqlx v1.3.4
	github.com/kr/pretty v0.2.0 // indirect
	github.com/lib/pq v1.10.2
	github.com/micro/go-micro v1.18.0 // indirect
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/broker/nats v0.0.0-20200119172437-4fe21aa238fd // indirect
	github.com/micro/go-plugins/broker/nats/v2 v2.9.1 // indirect
	github.com/satori/go.uuid v1.2.0
	github.com/teris-io/shortid v0.0.0-20171029131806-771a37caa5cf // indirect
	go.mongodb.org/mongo-driver v1.5.3
	golang.org/x/crypto v0.0.0-20200709230013-948cd5f35899
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
)
