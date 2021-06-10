module github.com/thealphadollar/go-microservices-PG

go 1.16

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gogo/protobuf v1.2.2-0.20190723190241-65acae22fc9d // indirect
	github.com/golang/protobuf v1.5.2
	github.com/jmoiron/sqlx v1.3.4
	github.com/kr/pretty v0.2.0 // indirect
	github.com/lib/pq v1.10.2 // indirect
	github.com/micro/go-micro/v2 v2.9.1-0.20200723075038-fbdf1f2c1c4c
	github.com/satori/go.uuid v1.2.0
	go.mongodb.org/mongo-driver v1.5.3
	golang.org/x/crypto v0.0.0-20200709230013-948cd5f35899
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
)
