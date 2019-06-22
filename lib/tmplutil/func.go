package tmplutil

import (
	"github.com/dustin/go-humanize"
	"github.com/flosch/pongo2"
)

func init() {
	pongo2.RegisterFilter("comma", CommaInt)
}

// CommaInt ...
func CommaInt(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	if in == nil || in.IsNil() || !in.IsNumber() {
		return pongo2.AsValue(""), nil
	}

	n := humanize.Comma(int64(in.Integer()))
	return pongo2.AsValue(n), nil
}
