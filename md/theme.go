package main

const (
	TEXT           = "text"
	SB             = "softbreak"
	HB             = "hardbreak"
	EMPH           = "emph"
	STRONG         = "strong"
	DEL            = "del"
	LINK           = "link" // show link addr?
	IMAGE          = "image"
	CODE           = "code"
	DOC            = "doc"
	PARAGRAPH      = "paragraph"
	BLOCKQUOTE     = "blockquote"
	HTMLBLOCK      = "htmlblock"
	HTMLSPAN       = "htmlspan"
	HEADING1       = "heading1"
	HEADING2       = "heading2"
	HEADING3       = "heading3"
	HEADING4       = "heading4"
	HEADING5       = "heading5"
	HORIZONTALRULE = "horizontalrule"
	LIST           = "list"
	LIST_ITEM      = "listitem"
	CODEBLOCK      = "codeblock"
	TABLE          = "table"
	TABLE_HEAD     = "tablehead"
	TABLE_BODY     = "tablebody"
	TABLE_ROW      = "tablerow"
	TABLE_CELL     = "tablecell"
)

type Style struct {
	_italic        bool
	_bold          bool
	_dim           bool
	_uline         bool // underline
	_duline        bool // doubly underline
	_strikethrough bool
	_blink         bool
	_hidden        bool
	_fg            int // color 256, minus for default color
	_bg            int // color 256, minus for default color
}

type md_theme struct {
	styles map[string]*Style
	chars  struct {
		list      rune // unordered list
		task      rune // task
		task_done rune // task done
		image     rune // image
	}
}

// load theme from a json conf
func LoadTheme(p string) *md_theme {
	return nil
}
func _new_style() *Style {
	return &Style{
		_fg: -1,
		_bg: -1,
	}
}

func (s *md_theme) GetStyle(key string) *Style {
	r := s.styles[key]
	if r == nil {
		r = _new_style()
	}
	return r
}

/*
just for test
§№☆★○●◎◇◆□■△▲※→←↑↓〓＃＆＠＼＾＿￣°℃
*/
func TestTheme() *md_theme {
	ret := &md_theme{}
	// chars
	ret.chars.list = '●'
	ret.chars.task = '◇'
	ret.chars.task_done = '◆'
	ret.chars.image = '※'

	// styles
	ret.styles = make(map[string]*Style)
	ret.styles[TEXT] = _new_style().fg(238)
	ret.styles[EMPH] = _new_style().italic(true)
	ret.styles[STRONG] = _new_style().bold(true)
	ret.styles[DEL] = _new_style().strikethrough(true)
	ret.styles[LINK] = _new_style().uline(true).fg(20)
	ret.styles[IMAGE] = _new_style().uline(true).fg(20)
	ret.styles[CODE] = _new_style().bg(236)       // grey bg
	ret.styles[CODEBLOCK] = _new_style().bg(236)  // grey bg
	ret.styles[BLOCKQUOTE] = _new_style().bg(252) // grey bg
	ret.styles[HEADING1] = _new_style().bold(true).fg(232)
	ret.styles[HEADING2] = _new_style().bold(true).fg(233)
	ret.styles[HEADING3] = _new_style().bold(true).fg(234)
	ret.styles[HEADING4] = _new_style().bold(true).fg(235)
	ret.styles[HEADING5] = _new_style().bold(true).fg(236)
	ret.styles[TABLE_HEAD] = _new_style().bold(true)
	return ret
}

func (s *Style) italic(v bool) *Style {
	s._italic = v
	return s
}
func (s *Style) bold(v bool) *Style {
	s._bold = v
	return s
}
func (s *Style) dim(v bool) *Style {
	s._dim = v
	return s
}
func (s *Style) uline(v bool) *Style {
	s._uline = v
	return s
}
func (s *Style) duline(v bool) *Style {
	s._duline = v
	return s
}
func (s *Style) strikethrough(v bool) *Style {
	s._strikethrough = v
	return s
}
func (s *Style) blink(v bool) *Style {
	s._blink = v
	return s
}
func (s *Style) hidden(v bool) *Style {
	s._hidden = v
	return s
}
func (s *Style) fg(v int) *Style {
	s._fg = v
	return s
}
func (s *Style) bg(v int) *Style {
	s._bg = v
	return s
}
