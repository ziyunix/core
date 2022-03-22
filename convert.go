package core

import (
	`fmt`
	`path/filepath`

	`github.com/goexl/gfx`
	`github.com/storezhang/media`
)

func (c *Convert) SegmentPattern() string {
	return fmt.Sprintf("%s\\.%s|%s-\\d+\\.%s", c.File.Name, hlsPlaylistExt, c.File.Name, hlsSegmentExt)
}

func (c *Convert) SegmentFilename(home string) string {
	return fmt.Sprintf(`%s-%%d.%s`, filepath.Join(c.DestDir(home), gfx.Filename(c.File.Name)), hlsSegmentExt)
}

func (c *Convert) DestDir(home string) string {
	return filepath.Join(home, `dest`, filepath.Base(c.File.Name))
}

func (c *Convert) DestFilename(home string) (filename string) {
	dir := c.DestDir(home)
	if nil != c.Hls {
		filename = filepath.Join(dir, gfx.Filename(c.File.Name, gfx.Ext(hlsPlaylistExt)))
	} else if media.Type_AUDIO == c.Type {
		filename = filepath.Join(dir, gfx.Filename(c.File.Name, gfx.Ext(c.Audio.Format.String())))
	} else if media.Type_VIDEO == c.Type {
		filename = filepath.Join(dir, gfx.Filename(c.File.Name, gfx.Ext(c.Video.Format.String())))
	}

	return
}
