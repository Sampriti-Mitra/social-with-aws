package main

import (
	"weekend.side/SocialMedia/internal/infra/db"
	"weekend.side/SocialMedia/internal/router"
)

func main() {
	db.InitialiseDb()
	db.Session = db.ConnectAws()
	router.Initialize()
}
