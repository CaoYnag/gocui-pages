package main

type style_stack []*Style

// stack: push
func (s style_stack) _push(st *Style) style_stack {
	return append(s, st)
}

// stack: pop
func (s style_stack) _pop() (ss style_stack, st *Style) {
	l := len(s)
	if l == 0 {
		return s, nil
	}
	l = l - 1
	return s[:l], s[l]
}

func (s *TermRender) push_style(st *Style) {
	old := s.curs
	s.curs = st
	s.styles = s.styles._push(st)
	s.update_style(old, s.curs)
}

func (s *TermRender) pop_style() {}

func (s *TermRender) update_style(old *Style, now *Style) {}
