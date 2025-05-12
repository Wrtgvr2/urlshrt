package services

import (
	rep "github.com/wrtgvr/urlshrt/internal/repository"
)

type TokensCleanupServices struct {
	Repo rep.TokensCleanupRepo
}

func NewTokensCleanupServices(repo rep.TokensCleanupRepo) TokensCleanupServices {
	return TokensCleanupServices{
		Repo: repo,
	}
}

// Clean tokens which expired twice
func (s *TokensCleanupServices) CleanOldTokens() (int64, error) {
	deletedRows, err := s.Repo.DeleteTooOldTokens()
	if err != nil {
		return 0, err
	}

	return deletedRows, nil
}

func (s *TokensCleanupServices) RevokeExpiredTokens() (int64, error) {
	revokedRows, err := s.Repo.RevokeExpiredTokens()
	if err != nil {
		return 0, err
	}

	return revokedRows, nil
}
