package routes

import (
	"github.com/labstack/echo/v4"
	"ledgerApp/src/utils/app/services/auth"
	"ledgerApp/src/utils/app/services/transactions"
	"ledgerApp/src/utils/app/services/users"
	"ledgerApp/src/utils/app/services/wallets"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *echo.Echo {
	e := echo.New()

	e.POST("/app/services/users/create", users.CreateUserHandler(db))
	e.POST("/app/services/wallets/update", wallets.UpdateWalletBalanceHandler(db))
	e.POST("/app/services/transactions/create", transactions.CreateTransactionHandler(db))
	e.GET("/app/services/wallets/get", wallets.GetWalletHandler(db))
	e.GET("/app/services/transactions/get", transactions.GetTransactionsHandler(db))

	adminGroup := e.Group("/app/services", auth.AuthenticateUserMiddleware(db))
	adminGroup.POST("/users/update-user", users.UpdateUserHandler(db))
	adminGroup.POST("/users/delete-user", users.DeleteUserHandler(db))
	adminGroup.POST("/users/update-user-role", users.UpdateUserRoleHandler(db))
	adminGroup.POST("/wallets/update-wallet-balance-by-user-id", wallets.UpdateWalletBalanceByUserIDHandler(db))
	adminGroup.POST("/transactions/get-wallet-balance-at-time", wallets.GetWalletBalanceAtTimeHandler(db))
	adminGroup.GET("/users/get-all-users", users.GetAllUsersHandler(db))
	adminGroup.GET("/users/get-all-users-with-wallet", users.GetAllUsersWithWalletHandler(db))
	adminGroup.GET("/wallets/get-all-wallets", wallets.GetAllWallets(db))
	adminGroup.GET("/users/get-user-by-id/:id", users.GetUserByIDHandler(db))
	adminGroup.GET("/wallets/get-wallet-by-id/:id", wallets.GetWalletByID(db))

	return e
}
