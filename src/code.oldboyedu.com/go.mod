module code.oldboyedu.com

go 1.13

require (
	github.com/Shopify/sarama v1.29.1
	github.com/garyburd/redigo v1.6.2
	github.com/gin-gonic/gin v1.7.2
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/go-sql-driver/mysql v1.6.0
	github.com/golang/protobuf v1.5.2
	github.com/jmoiron/sqlx v1.3.4
	github.com/micro/go-micro v1.18.0
	github.com/nsqio/go-nsq v1.0.8
	github.com/olivere/elastic/v7 v7.0.26
	github.com/satori/go.uuid v1.2.0
	golang.org/x/net v0.0.0-20210726213435-c6fcb2dbf985
	google.golang.org/grpc v1.36.0
	google.golang.org/grpc/examples v0.0.0-20210727191318-245ad25715e0 // indirect
	gopkg.in/ini.v1 v1.62.0
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
