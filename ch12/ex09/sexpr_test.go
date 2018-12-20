// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package sexpr

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

// Test verifies that encoding and decoding a complex data value
// produces an equal result.
//
// The test does not make direct assertions about the encoded output
// because the output depends on map iteration order, which is
// nondeterministic.  The output of the t.Log statements can be
// inspected by running the test with the -v flag:
//
// 	$ go test -v gopl.io/ch12/sexpr
//
func Test(t *testing.T) {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}
	data, err := Marshal(strangelove)
	if err != nil {
		fmt.Println("EOF")
		t.Fatalf("Marshal failed: %v", err)
	}
	fmt.Printf("%s\n", data)
	decoder := NewDecoder(bytes.NewReader(data))
	for {
		tok, err := decoder.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			t.Errorf("unexpected syntax %v\n", err)
		}
		switch tok := tok.(type) {
		case String:
			{
				fmt.Printf("%s ", tok.Val)
			}
		case Int:
			{
				fmt.Printf("%d ", tok.Val)
			}
		case Symbol:
			{
				fmt.Printf("%s: ", tok.Name)
			}
		case StartList:
			{
				fmt.Printf("[")
			}
		case EndList:
			{
				fmt.Printf("]")
			}
		}
	}

}
