package repository

import (
	"testing"

	"github.com/sssjuing/media-vault/internal/app/server/db"
)

func createActressRepository() ActressRepository {
	d, _ := db.NewDB(db.DatabaseTypeSQLite, "../../../../media-vault.db")
	return NewActressRepositoryImpl(d)
}

func TestFindVideos(t *testing.T) {
	actressRepo := createActressRepository()
	videos, _ := actressRepo.FindVideos(62)
	for _, v := range videos {
		t.Log("\n", v.SerialNumber, v.ReleaseDate)
		for _, nameSet := range v.ActressNameSets {
			if nameSet.Actress != nil {
				t.Log(nameSet.Actress.UniqueName)
			}
		}
	}
}
