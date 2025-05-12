package main

import (
	"fmt"

	"github.com/wrtgvr/urlshrt/internal/app"
	rep "github.com/wrtgvr/urlshrt/internal/repository"
	"github.com/wrtgvr/urlshrt/internal/services"
)

func main() {
	app.LoadEnvVars()

	repo := rep.NewPostgresTokensCleanupRepo(rep.InitDatabase())
	s := services.NewTokensCleanupServices(
		repo,
	)
	deletedRows, err := s.CleanOldTokens()
	if err != nil {
		fmt.Printf("Can't clean up tokens: %v\n", err)
		return
	}
	revokedRows, err := s.RevokeExpiredTokens()
	if err != nil {
		fmt.Printf("Can't revoke expired tokens: %v\n", err)
		return
	}

	fmt.Printf("Deleted rows: %d", deletedRows)
	fmt.Printf("Revoked rows: %d", revokedRows)
}
