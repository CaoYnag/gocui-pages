package main

type style struct {
	italic        bool
	bold          bool
	dim           bool
	uline         bool // underline
	duline        bool // doubly underline
	strikethrough bool
	blink         bool
	hidden        bool
	fg            int // color 256, minus for default color
	bg            int // color 256, minus for default color
}

type md_theme struct {
	styles map[string]*style
	chars  struct {
		list      rune // unordered list
		task      rune // task
		task_done rune /// task done
	}
}

// load theme from a json conf
func LoadTheme(p string) *md_theme {
	return nil
}
func _new_style() *style {
	return &style{
		fg: -1,
		bg: -1,
	}
}

/* just for test */
func TestTheme() *md_theme {
	return nil
}
