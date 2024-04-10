package db

import (
	"context"
	"testing"
	"time"

	"github.com/pe-Gomes/short-url/util"
	"github.com/stretchr/testify/require"
)

func createRandomURL(t *testing.T, user User) ShortLink {
	arg := CreateShortURLParams{
		Url:    util.RandomString(10),
		Slug:   util.RandomString(6),
		UserID: user.ID,
	}

	url, err := testStore.CreateShortURL(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, url)
	require.Equal(t, arg.Url, url.Url)
	require.Equal(t, arg.Slug, url.Slug)
	require.NotZero(t, url.ID)
	require.NotZero(t, url.CreatedAt)
	require.NotZero(t, url.UpdatedAt)

	return url
}

func TestCreateShortURL(t *testing.T) {
	user := createRandomUser(t)
	createRandomURL(t, user)
}

func TestGetShortURl(t *testing.T) {
	user := createRandomUser(t)
	url1 := createRandomURL(t, user)

	url2, err := testStore.GetShortURL(context.Background(), url1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, url2)
	require.Equal(t, url1.ID, url2.ID)
	require.Equal(t, url1.Url, url2.Url)
	require.Equal(t, url1.Slug, url2.Slug)
	require.WithinDuration(t, url1.CreatedAt, url2.CreatedAt, time.Second)
	require.WithinDuration(t, url1.UpdatedAt, url2.UpdatedAt, time.Second)
}

func TestGetShortURLBySlug(t *testing.T) {
	user := createRandomUser(t)
	url1 := createRandomURL(t, user)

	url2, err := testStore.GetShortURLBySlug(context.Background(), url1.Slug)
	require.NoError(t, err)
	require.NotEmpty(t, url2)
	require.Equal(t, url1.ID, url2.ID)
	require.Equal(t, url1.Url, url2.Url)
	require.Equal(t, url1.Slug, url2.Slug)
	require.WithinDuration(t, url1.CreatedAt, url2.CreatedAt, time.Second)
	require.WithinDuration(t, url1.UpdatedAt, url2.UpdatedAt, time.Second)
}

func TestListShortURLs(t *testing.T) {
	user := createRandomUser(t)

	createRandomURL(t, user)
	createRandomURL(t, user)

	user2 := createRandomUser(t)

	arg := ListShortURLsParams{Limit: 2, Offset: 0, UserID: user.ID}

	urls, err := testStore.ListShortURLs(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, urls)
	require.Len(t, urls, 2)

	for _, url := range urls {
		require.NotEmpty(t, url)
		require.Equal(t, url.UserID, user.ID)
		require.NotEqual(t, url.UserID, user2.ID)
	}

	arg2 := ListShortURLsParams{Limit: 2, Offset: 0, UserID: user2.ID}

	urls2, err := testStore.ListShortURLs(context.Background(), arg2)
	require.NoError(t, err)
	require.Empty(t, urls2)
	require.Len(t, urls2, 0)
}

func TestUpdateShortURL(t *testing.T) {
	user := createRandomUser(t)
	url1 := createRandomURL(t, user)

	arg := UpdateShortURLParams{
		ID:   url1.ID,
		Url:  util.RandomString(10),
		Slug: util.RandomString(6),
	}

	url2, err := testStore.UpdateShortURL(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, url2)
	require.Equal(t, url1.ID, url2.ID)
	require.Equal(t, arg.Url, url2.Url)
	require.Equal(t, arg.Slug, url2.Slug)
	require.WithinDuration(t, url1.CreatedAt, url2.CreatedAt, time.Second)
	require.WithinDuration(t, url1.UpdatedAt, url2.UpdatedAt, time.Second)
}

func TestDeleteShortURL(t *testing.T) {
	user := createRandomUser(t)
	url1 := createRandomURL(t, user)

	err := testStore.DeleteShortURL(context.Background(), url1.ID)
	require.NoError(t, err)

	url2, err := testStore.GetShortURL(context.Background(), url1.ID)
	require.Error(t, err)
	require.Empty(t, url2)
}
