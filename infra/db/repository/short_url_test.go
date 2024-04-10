package db

import (
	"context"
	"testing"
	"time"

	"github.com/pe-Gomes/short-url/util"
	"github.com/stretchr/testify/require"
)

func createRandomURL(t *testing.T) ShortLink {
	user := createRandomUser(t)

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
	createRandomURL(t)
}

func TestGetShortURl(t *testing.T) {
	url1 := createRandomURL(t)

	url2, err := testStore.GetShortURL(context.Background(), url1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, url2)
	require.Equal(t, url1.ID, url2.ID)
	require.Equal(t, url1.Url, url2.Url)
	require.Equal(t, url1.Slug, url2.Slug)
	require.WithinDuration(t, url1.CreatedAt, url2.CreatedAt, time.Second)
	require.WithinDuration(t, url1.UpdatedAt, url2.UpdatedAt, time.Second)
}

func TestListShortURLs(t *testing.T) {
	_ = createRandomURL(t)
	_ = createRandomURL(t)

	arg := ListShortURLsParams{Limit: 2, Offset: 0}

	urls, err := testStore.ListShortURLs(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, urls)
	require.Len(t, urls, 2)
}
