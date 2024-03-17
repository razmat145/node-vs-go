package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateGarbage(c echo.Context) error {
	createGarbage()

	return c.String(http.StatusOK, "ok!")
}

type Garbage struct {
	a string
	b string
	c string
	d string
	e string
	f string
	g string
	h string
	i string
	j string
	k string
	l string
	m string
	n string
	o string
	p string
	q string
	r string
	s string
	t string
	u string
	v string
	w string
	x string
	y string
	z string
}

func createGarbage() {
	for i := 0; i < 10_000; i++ {
		g := Garbage{
			a: "a",
			b: "b",
			c: "c",
			d: "d",
			e: "e",
			f: "f",
			g: "g",
			h: "h",
			i: "i",
			j: "j",
			k: "k",
			l: "l",
			m: "m",
			n: "n",
			o: "o",
			p: "p",
			q: "q",
			r: "r",
			s: "s",
			t: "t",
			u: "u",
			v: "v",
			w: "w",
			x: "x",
			y: "y",
			z: "z",
		}
		_ = g
	}
}
