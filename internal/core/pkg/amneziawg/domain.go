package amneziawg

import (
	"fmt"
	"os"
)

type Obfuscation struct {
	Jc   string
	Jmin string
	Jmax string
	S1   string
	S2   string
	H1   string
	H2   string
	H3   string
	H4   string
}

func NewObfuscation() (*Obfuscation, error) {
	js := os.Getenv("JS")
	if len(js) == 0 {
		return nil, fmt.Errorf("JS is not set")
	}
	jmin := os.Getenv("JMIN")
	if len(jmin) == 0 {
		return nil, fmt.Errorf("JMIN is not set")
	}
	jmax := os.Getenv("JMAX")
	if len(jmax) == 0 {
		return nil, fmt.Errorf("JMAX is not set")
	}
	s1 := os.Getenv("S1")
	if len(s1) == 0 {
		return nil, fmt.Errorf("S1 is not set")
	}
	s2 := os.Getenv("S2")
	if len(s2) == 0 {
		return nil, fmt.Errorf("S2 is not set")
	}
	h1 := os.Getenv("H1")
	if len(h1) == 0 {
		return nil, fmt.Errorf("H1 is not set")
	}
	h2 := os.Getenv("H2")
	if len(h2) == 0 {
		return nil, fmt.Errorf("H2 is not set")
	}
	h3 := os.Getenv("H3")
	if len(h3) == 0 {
		return nil, fmt.Errorf("H3 is not set")
	}
	h4 := os.Getenv("H4")
	if len(h4) == 0 {
		return nil, fmt.Errorf("H4 is not set")
	}
	return &Obfuscation{
		Jc:   js,
		Jmin: jmin,
		Jmax: jmax,
		S1:   s1,
		S2:   s2,
		H1:   h1,
		H2:   h2,
		H3:   h3,
		H4:   h4,
	}, nil
}
