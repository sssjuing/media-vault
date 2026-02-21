package repository

import (
	"testing"

	"github.com/sssjuing/media-vault/internal/app/server/db"
	"github.com/sssjuing/media-vault/internal/pkg/config"
)

func createActressRepository() ActressRepository {
	dsn := config.GetPostgresDsn()
	d := db.NewDB(dsn, config.GetPostgresReplicas())
	return NewActressRepositoryImpl(d)
}

func TestFindVideos(t *testing.T) {
	actressRepo := createActressRepository()
	videos, _ := actressRepo.FindVideos(62)
	for _, v := range videos {
		t.Log("\n", v.SerialNumber, v.ReleaseDate)
		for _, a := range v.Actresses {
			t.Log(a.UniqueName)
		}
	}
}
