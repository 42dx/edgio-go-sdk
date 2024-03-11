package property_test

import (
	"edgio/common"
	"edgio/property"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	mux := http.NewServeMux()

	server := httptest.NewServer(mux)
	defer server.Close()

	server2 := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		_, err := rw.Write([]byte(`{"access_token": "test_token"}`))
		if err != nil {
			t.Fatal(err)
		}
	}))
	defer server2.Close()

	mux.HandleFunc("/accounts/v0.1/properties", func(rw http.ResponseWriter, _ *http.Request) {
		_, err := rw.Write([]byte(`{
            "total_items": 2,
            "items": [
                {
                    "id": "some-id",
                    "slug": "some-slug",
                    "created_at": "2019-08-24T14:15:22Z",
                    "updated_at": "2019-08-24T14:15:22Z"
                },
                {
                    "id": "another-id",
                    "slug": "another-slug",
                    "created_at": "2019-08-24T14:15:22Z",
                    "updated_at": "2019-08-24T14:15:22Z"
                }
            ]
        }`))
		if err != nil {
			t.Fatal(err)
		}
	})

	params := property.ClientParams{
		Credentials: common.Creds{
			Key:     "key",
			Secret:  "secret",
			Scopes:  "scopes",
			AuthURL: server2.URL,
		},
		Config: common.ClientConfig{URL: server.URL, OrgID: "some-org-id"},
	}

	client, _ := property.NewClient(params)
	result, _ := client.List()

	assert.Len(t, result.Items, 2)
	assert.Equal(t, result.Items[0].Slug, "some-slug")
	assert.Equal(t, result.Items[1].Slug, "another-slug")
}
