package core

import (
	`fmt`
	`path/filepath`

	`github.com/goexl/gfx`
)

func (c *Convert) SegmentFilename(home string) string {
	return fmt.Sprintf(`%s-%%d.%s`, filepath.Join(c.DestDir(home), gfx.Filename(c.File.Name)), hlsSegmentExt)
}

func (c *Convert) DestDir(home string) string {
	return filepath.Join(home, `dest`, filepath.Base(c.File.Name))
}
