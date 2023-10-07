package main

import (
	_ "file-system/internal/database"
	"file-system/internal/global"
	"file-system/internal/web/router"
	"fmt"
)

func main() {
	r := router.NewRouter()

	if err := r.Run(fmt.Sprintf("%s:%d", global.WebAddr, global.WebPort)); err != nil {
		fmt.Println("Server Start Failed ... :", err)
	}
}
