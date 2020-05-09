package main

import (
	"io/ioutil"
	"math/rand"
	"strings"
)

type ymsh struct {
	y, m, s, h []string
}

func newYMSH() (*ymsh, error) {
	buf, err := ioutil.ReadFile("assets/words_alpha.txt")
	if err != nil {
		return nil, err
	}

	ymsh := &ymsh{
		y: []string{},
		m: []string{},
		s: []string{},
		h: []string{},
	}

	words := strings.Fields(string(buf))
	for _, word := range words {
		lower := strings.ToLower(word)

		if strings.HasPrefix(lower, "y") {
			ymsh.y = append(ymsh.y, lower)
			continue
		}
		if strings.HasPrefix(lower, "m") {
			ymsh.m = append(ymsh.m, lower)
			continue
		}
		if strings.HasPrefix(lower, "s") {
			ymsh.s = append(ymsh.s, lower)
			continue
		}
		if strings.HasPrefix(lower, "h") {
			ymsh.h = append(ymsh.h, lower)
			continue
		}
	}

	return ymsh, nil
}

func (ymsh *ymsh) String(rand *rand.Rand) string {
	yWord := ymsh.y[rand.Intn(len(ymsh.y))]
	mWord := ymsh.m[rand.Intn(len(ymsh.m))]
	sWord := ymsh.s[rand.Intn(len(ymsh.s))]
	hWord := ymsh.h[rand.Intn(len(ymsh.h))]

	builder := []string{
		strings.Title(strings.ToLower(yWord)),
		strings.Title(strings.ToLower(mWord)),
		strings.Title(strings.ToLower(sWord)),
		strings.Title(strings.ToLower(hWord)),
	}

	return strings.Join(builder, " ")
}
