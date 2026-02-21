package repository

import (
	"testing"

	"github.com/sssjuing/media-vault/internal/app/server/db"
	"github.com/sssjuing/media-vault/internal/pkg/config"
)

func createVideoRepository() VideoRepository {
	dsn := config.GetPostgresDsn()
	d := db.NewDB(dsn, config.GetPostgresReplicas())
	return NewVideoRepositoryImpl(d)
}

func TestFindAll(t *testing.T) {
	videoRepo := createVideoRepository()
	videos, _ := videoRepo.FindAll(VidoesQueryOptions{Tags: []string{"OL"}})
	t.Log(len(videos))
	for _, v := range videos {
		t.Log("video:", v.SerialNumber)
		for _, a := range v.Actresses {
			t.Log("actress:", a.UniqueName)
		}
	}
	allVideos, _ := videoRepo.FindAll(VidoesQueryOptions{})
	t.Log(len(allVideos))
}

func TestPaginateAll(t *testing.T) {
	videoRepo := createVideoRepository()
	videos, _ := videoRepo.PaginateAll(VidoesQueryOptions{Page: 1, Size: 10, Search: "RO", Tags: []string{"OL"}})
	for _, v := range videos {
		title := "-"
		if v.Title != nil {
			title = *v.Title
		}
		t.Log(v.SerialNumber, title)
	}
}

func TestCount(t *testing.T) {
	videoRepo := createVideoRepository()
	count1, _ := videoRepo.Count(VidoesQueryOptions{})
	t.Log(count1)
	count2, _ := videoRepo.Count(VidoesQueryOptions{Tags: []string{"OL"}})
	t.Log(count2)
}
