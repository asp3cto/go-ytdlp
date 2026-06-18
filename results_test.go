// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package ytdlp

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockServer struct {
	*httptest.Server

	fileURL string
}

func newMockServer(t *testing.T, fileName string) *mockServer {
	t.Helper()

	base := filepath.Base(fileName)

	// TODO: potentially replace with FileServer + Go 1.24 os.Root.
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, base) {
			http.ServeFile(w, r, fileName)
			return
		}

		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("not found"))
	}))
	t.Cleanup(server.Close)

	return &mockServer{
		Server:  server,
		fileURL: server.URL + "/" + base,
	}
}

func TestExtractedInfo(t *testing.T) {
	server := newMockServer(t, "testdata/sample-1.mp4")

	dir := t.TempDir()

	result, err := New().
		ForceOverwrites().
		Output(filepath.Join(dir, "%(extractor)s - %(title)s.%(ext)s")).
		PrintJSON().
		Run(context.TODO(), server.fileURL)
	if err != nil {
		t.Fatal(err)
		return
	}

	info, err := result.GetExtractedInfo()
	if err != nil {
		t.Fatal(err)
	}

	require.Len(t, info, 1, "expected 1 extracted info")
	require.NotNil(t, info[0].FormatID, "expected format id to be set")
	assert.Equal(t, "mp4", *info[0].FormatID, "expected format id to be mp4")

	require.NotNil(t, info[0].Protocol, "expected protocol to be set")
	assert.Equal(t, "http", *info[0].Protocol, "expected protocol to be http")

	require.NotNil(t, info[0].HTTPHeaders, "expected http headers to be set")
	assert.Contains(t, info[0].HTTPHeaders["User-Agent"], "Mozilla", "expected User-Agent header to be set and contain Mozilla")

	assert.Equal(t, "sample-1", info[0].ID, "expected id to be set")

	require.NotNil(t, info[0].Title, "expected title to be set")
	assert.Equal(t, "sample-1", *info[0].Title, "expected title to be set")

	require.Len(t, info[0].Formats, 1, "expected 1 format")
	require.NotNil(t, info[0].Formats[0].Extension, "expected format extension to be set")
	assert.Equal(t, "mp4", *info[0].Formats[0].Extension, "expected format extension to be mp4")

	require.NotNil(t, info[0].URL, "expected url to be set")
	assert.Equal(t, server.fileURL, *info[0].URL, "expected url to be set")
	require.NotNil(t, info[0].WebpageURL, "expected webpage url to be set")
	assert.Equal(t, server.fileURL, *info[0].WebpageURL, "expected webpage url to be set")

	require.NotNil(t, info[0].Filename, "expected filename to be set")
	assert.FileExists(t, *info[0].Filename, "expected file to exist")

	require.NotNil(t, info[0].Timestamp, "expected timestamp to be set")
	assert.Positive(t, *info[0].Timestamp, "expected timestamp to be set")

	require.NotNil(t, info[0].UploadDate, "expected upload date to be set")
	assert.Positive(t, *info[0].UploadDate, "expected upload date to be set")

	require.NotNil(t, info[0].Extractor, "expected extractor to be set")
	assert.Equal(t, "generic", *info[0].Extractor, "expected extractor to be generic")

	require.NotNil(t, info[0].ExtractorKey, "expected extractor key to be set")
	assert.Equal(t, "Generic", *info[0].ExtractorKey, "expected extractor key to be generic")
}

func TestExtractedInfoParse(t *testing.T) {
	noneStr := "none"

	tests := []struct {
		testName       string
		json           json.RawMessage
		expectedFormat ExtractedFormat
	}{
		{
			testName: "acodec/vcodec: preserve none",
			json:     json.RawMessage(`{"acodec": "none", "vcodec": "none"}`),
			expectedFormat: ExtractedFormat{
				ACodec: &noneStr,
				VCodec: &noneStr,
			},
		},
		{
			testName: "acodec/vcodec: parse null as nil",
			json:     json.RawMessage(`{"acodec": null, "vcodec": null}`),
			expectedFormat: ExtractedFormat{
				ACodec: nil,
				VCodec: nil,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {
			got, err := ParseExtractedInfo(&tc.json)
			require.NoError(t, err)

			if tc.expectedFormat.ACodec == nil {
				assert.Equal(t, tc.expectedFormat.ACodec, got.ACodec)
			} else {
				assert.Equal(t, *tc.expectedFormat.ACodec, *got.ACodec)
			}

			if tc.expectedFormat.VCodec == nil {
				assert.Equal(t, tc.expectedFormat.VCodec, got.VCodec)
			} else {
				assert.Equal(t, *tc.expectedFormat.VCodec, *got.VCodec)
			}
		})
	}

	t.Run("title: parse none as empty string", func(t *testing.T) {
		raw := json.RawMessage(`{"title": "none"}`)
		got, err := ParseExtractedInfo(&raw)
		require.NoError(t, err)

		assert.Equal(t, "", *got.Title)
	})

	t.Run("other field: parse none as nil", func(t *testing.T) {
		raw := json.RawMessage(`{"thumbnail": "none"}`)
		got, err := ParseExtractedInfo(&raw)
		require.NoError(t, err)

		assert.Nil(t, got.Thumbnail)
	})
}
