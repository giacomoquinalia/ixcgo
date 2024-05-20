package main

import (
	"fmt"
	"log/slog"
)

// START OMIT
func hello() { // HL12
	fmt.Println("Hello world!") // HL
} // HL12
//END OMIT


* Exemplo real
.code hello.go
Display Partial code
.code hello.go /^//START/,/^//END/ HL12
.background images/background.png

* Exemplo real
.play hello.go

* Bullet Points

# Title of Talk

My Name
9 Mar 2020
me@example.com

## Title of Slide or Section (must begin with ##)

Some Text

### Subsection {#anchor}

- bullets
- more bullets
- a bullet continued
  on the next line

#### Sub-subsection

	Preformatted text (code block)
	is indented (by one tab, or four spaces)

Further Text, including command invocations.

## Section 2: Example formatting {#fmt}

// A comment that is completely ignored.
: Speaker notes.

`program`
Markup—_especially italic text_—can easily be overused.
_Why use scoped\_ptr_? Use plain **\*ptr** instead.

Visit [the Go home page](https://golang.org/).