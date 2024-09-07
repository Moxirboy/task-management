package usecase

import (
	"database/sql"
	"food-delivery/internal/configs"
	"food-delivery/internal/service/storage/postgres"
	"food-delivery/pkg/logger"
)

type IUseCase interface {
	IAccountUseCase() IAccountUseCase
	IAuthUseCase() IAuthUseCase
	ProductUsecase() IProductUseCase
	CartUsecase() ICartUseCase
	IOrderUseCase() IOrderUseCase
}

type UseCase struct {
	connections map[string]interface{}
}

const (
	_AuthUseCase    = "auth_use_case"
	_AccountUseCase = "account_use_case"
	_productUseCase = "product_use_case"
	_cartUseCase    = "cart_use_case"
	_orderUseCase   = "order_use_case"
)

func New(
	cfg *configs.Config,
	pg *sql.DB,
	logger logger.Logger,
) IUseCase {
	var connections = make(map[string]interface{})
	connections[_AuthUseCase] = NewAuthUseCase(
		postgres.NewAuthTokenRepository(
			pg,
			logger,
		),
		logger,
	)
	connections[_AccountUseCase] = NewAccountUseCase(
		postgres.NewUserRepository(
			pg,
			logger,
		),
		logger,
		cfg,
	)
	connections[_productUseCase] = NewProductUsecase(
		postgres.NewProduct(
			pg,
			logger,
		),
		logger,
	)

	connections[_cartUseCase] = NewCartUsecase(
		postgres.NewCart(
			pg,
			logger,
		),
		logger,
	)
	connections[_orderUseCase] = NewOrderUseCase(
		postgres.NewOrder(
			pg,
			logger,
		),
		logger,
	)

	return &UseCase{
		connections: connections,
	}
}

func (c *UseCase) IAuthUseCase() IAuthUseCase {
	return c.connections[_AuthUseCase].(IAuthUseCase)
}
func (c *UseCase) IAccountUseCase() IAccountUseCase {
	return c.connections[_AccountUseCase].(IAccountUseCase)
}

func (c *UseCase) ProductUsecase() IProductUseCase {
	return c.connections[_productUseCase].(IProductUseCase)
}
func (c *UseCase) CartUsecase() ICartUseCase { return c.connections[_cartUseCase].(ICartUseCase) }

func (c *UseCase) IOrderUseCase() IOrderUseCase {
	return c.connections[_orderUseCase].(IOrderUseCase)
}
