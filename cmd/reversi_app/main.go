package main

import (
	"github.com/narakosen-festival-info-2020/reversi-back/pkg/api"
)

func main() {
	api.SeedInit()
	api.ServerUp()
}
