// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// Code generated by cmd/codegen. DO NOT EDIT.
//
// Extractor Option Group

package ytdlp

// Number of retries for known extractor errors (default is %default), or
// "infinite"
//
// ExtractorRetries maps to cli flags: --extractor-retries=RETRIES.
func (c *Command) ExtractorRetries(retries string) *Command {
	c.addFlag(&Flag{
		ID:   "extractor_retries",
		Flag: "--extractor-retries",
		Args: []string{retries},
	})
	return c
}

// Process dynamic DASH manifests (default) (Alias: --no-ignore-dynamic-mpd)
//
// AllowDynamicMpd maps to cli flags: --allow-dynamic-mpd/--no-ignore-dynamic-mpd.
func (c *Command) AllowDynamicMpd() *Command {
	c.addFlag(&Flag{
		ID:   "dynamic_mpd",
		Flag: "--allow-dynamic-mpd",
		Args: nil,
	})
	return c
}

// Do not process dynamic DASH manifests (Alias: --no-allow-dynamic-mpd)
//
// IgnoreDynamicMpd maps to cli flags: --ignore-dynamic-mpd/--no-allow-dynamic-mpd.
func (c *Command) IgnoreDynamicMpd() *Command {
	c.addFlag(&Flag{
		ID:   "dynamic_mpd",
		Flag: "--ignore-dynamic-mpd",
		Args: nil,
	})
	return c
}

// Split HLS playlists to different formats at discontinuities such as ad breaks
//
// HlsSplitDiscontinuity maps to cli flags: --hls-split-discontinuity.
func (c *Command) HlsSplitDiscontinuity() *Command {
	c.addFlag(&Flag{
		ID:   "hls_split_discontinuity",
		Flag: "--hls-split-discontinuity",
		Args: nil,
	})
	return c
}

// Do not split HLS playlists to different formats at discontinuities such as ad
// breaks (default)
//
// NoHlsSplitDiscontinuity maps to cli flags: --no-hls-split-discontinuity.
func (c *Command) NoHlsSplitDiscontinuity() *Command {
	c.addFlag(&Flag{
		ID:   "hls_split_discontinuity",
		Flag: "--no-hls-split-discontinuity",
		Args: nil,
	})
	return c
}

// Pass ARGS arguments to the IE_KEY extractor. See "EXTRACTOR ARGUMENTS" for
// details. You can use this option multiple times to give arguments for different
// extractors
//
// ExtractorArgs maps to cli flags: --extractor-args=IE_KEY:ARGS.
func (c *Command) ExtractorArgs(ieKeyargs string) *Command {
	c.addFlag(&Flag{
		ID:   "extractor_args",
		Flag: "--extractor-args",
		Args: []string{ieKeyargs},
	})
	return c
}

// YoutubeIncludeDashManifest sets the "youtube-include-dash-manifest" flag (no description specified).
//
// YoutubeIncludeDashManifest maps to cli flags: --youtube-include-dash-manifest/--no-youtube-skip-dash-manifest.
func (c *Command) YoutubeIncludeDashManifest() *Command {
	c.addFlag(&Flag{
		ID:   "youtube_include_dash_manifest",
		Flag: "--youtube-include-dash-manifest",
		Args: nil,
	})
	return c
}

// YoutubeSkipDashManifest sets the "youtube-skip-dash-manifest" flag (no description specified).
//
// YoutubeSkipDashManifest maps to cli flags: --youtube-skip-dash-manifest/--no-youtube-include-dash-manifest.
func (c *Command) YoutubeSkipDashManifest() *Command {
	c.addFlag(&Flag{
		ID:   "youtube_include_dash_manifest",
		Flag: "--youtube-skip-dash-manifest",
		Args: nil,
	})
	return c
}

// YoutubeIncludeHlsManifest sets the "youtube-include-hls-manifest" flag (no description specified).
//
// YoutubeIncludeHlsManifest maps to cli flags: --youtube-include-hls-manifest/--no-youtube-skip-hls-manifest.
func (c *Command) YoutubeIncludeHlsManifest() *Command {
	c.addFlag(&Flag{
		ID:   "youtube_include_hls_manifest",
		Flag: "--youtube-include-hls-manifest",
		Args: nil,
	})
	return c
}

// YoutubeSkipHlsManifest sets the "youtube-skip-hls-manifest" flag (no description specified).
//
// YoutubeSkipHlsManifest maps to cli flags: --youtube-skip-hls-manifest/--no-youtube-include-hls-manifest.
func (c *Command) YoutubeSkipHlsManifest() *Command {
	c.addFlag(&Flag{
		ID:   "youtube_include_hls_manifest",
		Flag: "--youtube-skip-hls-manifest",
		Args: nil,
	})
	return c
}
