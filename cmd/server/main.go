package main

import (
	"github.com/nghiatk54/go_ecommerce_api/internal/router"
)

func main() {
	r := router.NewRouter()
	r.Run(":8002") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
