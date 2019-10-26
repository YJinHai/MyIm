package cmd

import (
	"../app/router"
)

func main() {
	r := router.Load()

	r.Run(":4990")
}
