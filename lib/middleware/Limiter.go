package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
	"log"
)

func LimiterMiddleware() gin.HandlerFunc {
	rate, err := limiter.NewRateFromFormatted("10-S")
	if err != nil {
		log.Fatal(err)
	}
	store := memory.NewStore()
	instance := limiter.New(store, rate)
	middleware := mgin.NewMiddleware(instance)
	return middleware
}
