package service

import (
	"article/config"
	"article/database"
	"article/models"
	"article/repository"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
	"syreclabs.com/go/faker"
)

var s *PostService

func init() {
	cfg := config.LoadConfig()
	db := database.InitDB(cfg)
	repo := repository.NewPostsRepository(db)

	s = NewPostsService(repo)
}

func TestCreatePost(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		post := &models.Posts{
			Title:    faker.Lorem().Characters(20),
			Content:  faker.Lorem().Characters(200),
			Category: faker.Lorem().Characters(3),
			Status:   "publish",
		}

		err := s.Create(post)
		require.Nil(t, err)
	})

	t.Run("field required", func(t *testing.T) {
		post := &models.Posts{
			Title:    "",
			Content:  faker.Lorem().Characters(200),
			Category: faker.Lorem().Characters(3),
			Status:   "publish",
		}

		err := s.Create(post)
		require.NotNil(t, err)
		log.Println(err)
	})

	t.Run("status not valid", func(t *testing.T) {
		post := &models.Posts{
			Title:    faker.Lorem().Characters(20),
			Content:  faker.Lorem().Characters(200),
			Category: faker.Lorem().Characters(3),
			Status:   "abc",
		}

		err := s.Create(post)
		require.NotNil(t, err)
		log.Println(err)
	})
}

func TestUpdatePost(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		post := &models.Posts{
			Id:       2,
			Title:    faker.Lorem().Characters(20),
			Content:  faker.Lorem().Characters(200),
			Category: faker.Lorem().Characters(3),
			Status:   "publish",
		}

		err := s.Update(post)
		require.Nil(t, err)
	})

	t.Run("field required", func(t *testing.T) {
		post := &models.Posts{
			Id:       2,
			Title:    faker.Lorem().Characters(20),
			Content:  "",
			Category: faker.Lorem().Characters(3),
			Status:   "publish",
		}

		err := s.Update(post)
		require.NotNil(t, err)
		log.Println(err)
	})

	t.Run("status not valid", func(t *testing.T) {
		post := &models.Posts{
			Id:       2,
			Title:    faker.Lorem().Characters(20),
			Content:  faker.Lorem().Characters(200),
			Category: faker.Lorem().Characters(3),
			Status:   "qwe",
		}

		err := s.Update(post)
		require.NotNil(t, err)
		log.Println(err)
	})
}

func TestDeletePost(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		id := 4
		err := s.Delete(id)
		require.Nil(t, err)
	})
}

func TestGetPost(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		posts, err := s.GetPosts(2, 0)
		require.Nil(t, err)
		log.Println(posts)
	})
}

func TestGetPostById(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		posts, err := s.GetPostsById(3)
		log.Println(posts)
		require.Nil(t, err)
	})
}
