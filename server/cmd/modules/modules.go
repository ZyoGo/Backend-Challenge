package modules

import (
	authBusi "github.com/ZyoGo/Backend-Challange/internal/auth/business"
	authApiHttp "github.com/ZyoGo/Backend-Challange/internal/auth/delivery/http"
	productBusi "github.com/ZyoGo/Backend-Challange/internal/products/business"
	productApiHttp "github.com/ZyoGo/Backend-Challange/internal/products/delivery/http"
	productRepo "github.com/ZyoGo/Backend-Challange/internal/products/repository/postgreSQL"
	userBusi "github.com/ZyoGo/Backend-Challange/internal/users/business"
	userIntra "github.com/ZyoGo/Backend-Challange/internal/users/delivery/intraprocess"
	userRepo "github.com/ZyoGo/Backend-Challange/internal/users/repository/postgreSQL"
	"github.com/ZyoGo/Backend-Challange/pkg/hash"
	"github.com/ZyoGo/Backend-Challange/pkg/jwt"
	"github.com/ZyoGo/Backend-Challange/pkg/ulid"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterModules(r *mux.Router, db *pgxpool.Pool) {
	genUlidID := ulid.NewULIDGenerator()
	jwt := jwt.NewBusiness()
	hash := hash.NewHash()

	// User modules
	userRepository := userRepo.NewPostgreSQL(db)
	userBusiness := userBusi.NewBusiness(userRepository, genUlidID)
	userIntraprocess := userIntra.NewIntraprocess(userBusiness)

	// Auth Modules
	authBusiness := authBusi.NewBusiness(userIntraprocess, jwt, hash)
	authHandler := authApiHttp.NewHandler(authBusiness)

	// Products Modules
	productRepository := productRepo.NewPostgreSQL(db)
	productBusiness := productBusi.NewBusiness(productRepository)
	productHandler := productApiHttp.NewHandler(productBusiness)

	authApiHttp.RegisterPath(r, authHandler)
	productApiHttp.RegisterPath(r, productHandler)
}
