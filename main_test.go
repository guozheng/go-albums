package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAlbums(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/albums", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	expected, err := json.Marshal(albums)
	if err != nil {
		t.Fatal("failed to create expected albums JSON")
	}
	assert.Equal(t, string(expected), w.Body.String())
}

func TestGetAlbumByID(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/albums/2", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	album := albums[1]
	expected, err := json.Marshal(album)
	if err != nil {
		t.Fatal("failed to create expected albums JSON")
	}
	assert.Equal(t, string(expected), w.Body.String())
}
