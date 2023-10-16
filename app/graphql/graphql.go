package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/nsqio/go-nsq"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"vehicle-log-graphql-api/app/graphql/graph"
	graphCommon "vehicle-log-graphql-api/common/graphql/graph"
	"vehicle-log-graphql-api/config"
	grpcClient "vehicle-log-graphql-api/internal/common/grpc"
	grpcCommon "vehicle-log-graphql-api/internal/common/grpc/impl"
	"vehicle-log-graphql-api/internal/common/publisher"
	publisherImpl "vehicle-log-graphql-api/internal/common/publisher/impl"
	logHandlerIMpl "vehicle-log-graphql-api/internal/handlers/log/impl"
	townHandlerImpl "vehicle-log-graphql-api/internal/handlers/town/impl"
	typesHandlerImpl "vehicle-log-graphql-api/internal/handlers/types/impl"
	vehicleHandlerImpl "vehicle-log-graphql-api/internal/handlers/vehicles/impl"
	"vehicle-log-graphql-api/internal/repository/redis"
	redisImpl "vehicle-log-graphql-api/internal/repository/redis/impl"
	logUc "vehicle-log-graphql-api/internal/usecase/log/impl"
	townUc "vehicle-log-graphql-api/internal/usecase/town/impl"
	typesUc "vehicle-log-graphql-api/internal/usecase/types/impl"
	vehicleUc "vehicle-log-graphql-api/internal/usecase/vehicle/impl"
)

type Server struct {
	redis     redis.Redis
	cfg       *config.Config
	grpcCon   grpcClient.Client
	publisher publisher.Publisher
	producer  *nsq.Producer
}

func (g Server) prepareResolver() *graph.Resolver {
	// Use Case
	townUC := townUc.New(g.grpcCon, g.cfg, g.redis)
	typeUC := typesUc.New(g.grpcCon, g.cfg, g.redis)
	logUC := logUc.New(g.grpcCon, g.cfg, g.redis, g.publisher)
	vehicleUC := vehicleUc.New(g.grpcCon, g.cfg, g.redis)

	// Handler
	handlerTown := townHandlerImpl.New(townUC)
	handlerTypes := typesHandlerImpl.New(typeUC)
	handlerLog := logHandlerIMpl.New(logUC)
	handlerVehicle := vehicleHandlerImpl.New(vehicleUC)
	r := graph.Resolver{
		TownHandler:    handlerTown,
		TypeHandler:    handlerTypes,
		LogHandler:     handlerLog,
		VehicleHandler: handlerVehicle,
	}

	return &r
}

func (g Server) Start() {

	defer g.producer.Stop()

	srv := handler.NewDefaultServer(
		graphCommon.NewExecutableSchema(
			graphCommon.Config{Resolvers: g.prepareResolver()}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", disableCors(srv))

	log.Printf("connect to http://localhost%s/ for GraphQL playground", g.cfg.Server.Port)
	log.Fatal(http.ListenAndServe(g.cfg.Server.Port, nil))
}

func disableCors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, Accept-Encoding")
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Max-Age", "86400")
			w.WriteHeader(http.StatusOK)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func New() Server {
	cfg := config.LoadDefault()
	r := redisImpl.New(cfg)
	conn, err := grpc.Dial(cfg.Grpc.Url, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	gClient := grpcCommon.New(conn)
	configNsq := nsq.NewConfig()
	producer, err := nsq.NewProducer(cfg.Nsq.PublisherAddr, configNsq)
	if err != nil {
		panic(err)
	}
	p := publisherImpl.New(producer)
	return Server{
		cfg:       cfg,
		redis:     r,
		grpcCon:   gClient,
		publisher: p,
		producer:  producer,
	}
}
