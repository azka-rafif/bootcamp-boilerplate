//go:build wireinject
// +build wireinject

package main

import (
	"github.com/evermos/boilerplate-go/configs"
	// "github.com/evermos/boilerplate-go/event"
	// fooBarBazEvent "github.com/evermos/boilerplate-go/event/domain/foobarbaz"
	// "github.com/evermos/boilerplate-go/event/producer"
	"github.com/evermos/boilerplate-go/infras"
	// "github.com/evermos/boilerplate-go/internal/domain/foobarbaz"
	"github.com/evermos/boilerplate-go/internal/domain/materials"
	"github.com/evermos/boilerplate-go/internal/domain/products"
	"github.com/evermos/boilerplate-go/internal/handlers"
	"github.com/evermos/boilerplate-go/transport/http"
	// "github.com/evermos/boilerplate-go/transport/http/middleware"
	"github.com/evermos/boilerplate-go/transport/http/router"
	"github.com/google/wire"
)

// Wiring for configurations.
var configurations = wire.NewSet(
	configs.Get,
)

// Wiring for persistences.
var persistences = wire.NewSet(
	infras.ProvideMariaDBConn,
)

// Wiring for domain FooBarBaz.
// var domainFooBarBaz = wire.NewSet(
// 	// FooService interface and implementation
// 	foobarbaz.ProvideFooServiceImpl,
// 	wire.Bind(new(foobarbaz.FooService), new(*foobarbaz.FooServiceImpl)),
// 	// FooRepository interface and implementation
// 	foobarbaz.ProvideFooRepositoryMySQL,
// 	wire.Bind(new(foobarbaz.FooRepository), new(*foobarbaz.FooRepositoryMySQL)),
// 	// Producer interface and implementation
// 	producer.NewSNSProducer,
// 	wire.Bind(new(producer.Producer), new(*producer.SNSProducer)),
// )

var domainMaterials = wire.NewSet(
	materials.ProvideMaterialServiceImpl,
	wire.Bind(new(materials.MaterialService), new(*materials.MaterialServiceImpl)),
	materials.ProvideMaterialRepositoryMariaDB,
	wire.Bind(new(materials.MaterialRepository), new(*materials.MaterialRepositoryMariaDB)),
)

var domainProducts = wire.NewSet(
	products.ProvideProductServiceImpl,
	wire.Bind(new(products.ProductService), new(*products.ProductServiceImpl)),
	products.ProvideProductRepositoryMariaDB,
	wire.Bind(new(products.ProductRepository), new(*products.ProductRepositoryMariaDB)),
)

// Wiring for all domains.
var domains = wire.NewSet(
	domainMaterials, domainProducts,
)

// var authMiddleware = wire.NewSet(
// 	middleware.ProvideAuthentication,
// )

// Wiring for HTTP routing.
var routing = wire.NewSet(
	wire.Struct(new(router.DomainHandlers), "MaterialsHandler", "ProductHandler"),
	router.ProvideRouter,
	handlers.ProvideMaterialsHandler,
	handlers.ProvideProductHandler,
)

// Wiring for all domains event consumer.
// var evco = wire.NewSet(
// 	wire.Struct(new(event.Consumers), "FooBarBaz"),
// 	fooBarBazEvent.ProvideConsumerImpl,
// )

// Wiring for everything.
func InitializeService() *http.HTTP {
	wire.Build(
		// configurations
		configurations,
		// persistences
		persistences,
		// middleware
		// authMiddleware,
		// domains
		domains,
		// routing
		routing,
		// selected transport layer
		http.ProvideHTTP)
	return &http.HTTP{}
}

// Wiring the event needs.
// func InitializeEvent() event.Consumers {
// 	wire.Build(
// 		// configurations
// 		configurations,
// 		// persistences
// 		persistences,
// 		// domains
// 		domains,
// 		// event consumer
// 		evco)

// 	return event.Consumers{}
// }
