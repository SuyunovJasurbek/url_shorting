package postgres_test

import (
	"context"
	"math/rand"

	"testing"

	"github.com/SuyunovJasurbek/url_shorting/models"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func createRandomUrl(t *testing.T) *models.Url {

	user := createRandomuser(t)

	mockurl := &models.Url{
		UserID:     user.ID,
		OrgPath:    uuid.NewString(),
		ShortPath:  uuid.NewString(),
		Counter:    rand.Intn(100),
		Status:     rand.Intn(100),
		QrCodePath: uuid.NewString(),
		CreatedAt:  int64(rand.Intn(100)),
		UpdatedAt:  int64(rand.Intn(100)),
	}

	url, err := strg.Url().CreateUrl(context.Background(), mockurl)
	require.NoError(t, err)
	require.NotEmpty(t, url)
	require.Equal(t, mockurl.UserID, url.UserID)
	require.Equal(t, mockurl.OrgPath, url.OrgPath)
	require.Equal(t, mockurl.ShortPath, url.ShortPath)
	return url
}

func deleteRandomUrl(t *testing.T, url *models.Url) {
	err := strg.Url().DeleteUrlByID(context.Background(), url.ID)
	require.NoError(t, err)
	defer deleteRandomUser(t, url.UserID)
}

func Test_CreateUrl(t *testing.T) {
	deleteRandomUrl(t, createRandomUrl(t))
}

func Test_DeleteUrlByID(t *testing.T) {
	url := createRandomUrl(t)

	// bad case 1
	err := strg.Url().DeleteUrlByID(context.Background(), "invalid-uuid")
	require.Error(t, err)

	// bad case 2
	err = strg.Url().DeleteUrlByID(context.Background(), uuid.NewString())
	require.Error(t, err)

	// good case
	err = strg.Url().DeleteUrlByID(context.Background(), url.ID)
	require.NoError(t, err)
}

func Test_DeleteUrlByShortURL(t *testing.T) {
	url := createRandomUrl(t)

	// bad case 1
	err := strg.Url().DeleteUrlByShortURL(context.Background(), "no-short-url")
	require.Error(t, err)

	// good case

	err = strg.Url().DeleteUrlByShortURL(context.Background(), url.ShortPath)
	require.NoError(t, err)
}

func Test_GetUrlByShortPath(t *testing.T) {
	url := createRandomUrl(t)

	// bad case 1

	url2, err := strg.Url().GetUrlByShortPath(context.Background(), "no-short-url")
	require.Error(t, err)
	require.Empty(t, url2)

	// good case
	url2, err = strg.Url().GetUrlByShortPath(context.Background(), url.ShortPath)
	require.NoError(t, err)
	require.NotEmpty(t, url2)
	require.Equal(t, url.ID, url2.ID)

}

func Test_GetUrlByID(t *testing.T) {
	url := createRandomUrl(t)

	// bad case 1

	url2, err := strg.Url().GetUrlByID(context.Background(), "no-short-url")
	require.Error(t, err)
	require.Empty(t, url2)

	// good case
	url2, err = strg.Url().GetUrlByID(context.Background(), url.ID)
	require.NoError(t, err)
	require.NotEmpty(t, url2)
	require.Equal(t, url.ID, url2.ID)

}

func Test_GetUrls(t *testing.T) {

	url := createRandomUrl(t)
	defer deleteRandomUrl(t, url)
	// bad case 1
	urls, err := strg.Url().GetUrls(context.Background(), "invalid-uuid")
	require.Error(t, err)
	require.Empty(t, urls)

	urls, err = strg.Url().GetUrls(context.Background(), url.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, urls)
	require.Len(t, urls, 1)

}
