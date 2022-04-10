package routes

import "github.com/unrolled/render"

var renderer *render.Render

func init() {
	renderer = render.New()
}