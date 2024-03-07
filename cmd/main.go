package main

import (
	"github.com/labstack/echo/v4"

	// "github.com/smkdev-id/promotion_tracking_dashboard/internal/app/configs"
	// "github.com/smkdev-id/promotion_tracking_dashboard/internal/app/delivery"
	// "github.com/smkdev-id/promotion_tracking_dashboard/internal/app/repositories"
	// "github.com/smkdev-id/promotion_tracking_dashboard/internal/app/services"
	"github.com/iniardien/submission_promotion_api/internal/app/repositories"
	"github.com/iniardien/submission_promotion_api/internal/app/services"
	"github.com/iniardien/submission_promotion_api/internal/configs"
	"github.com/iniardien/submission_promotion_api/internal/delivery"
)

func main() {

	configs.LoadViperEnv()

	db := configs.InitDatabase()

	e := echo.New()

	PromotionRepo := repositories.NewPromotionRepository(db)

	PromoService := services.NewPromotionService(PromotionRepo)

	delivery.PromotionRoute(e, PromoService)

	e.Logger.Fatal(e.Start(":8080"))
}
