package main

type style_stack []*Style

func (s style_stack) push(st *Style) style_stack {
	return append(s, st)
}
func (s style_stack) pop() (ss style_stack, st *Style) {
	l := len(s)
	if l == 0 {
		return s, nil
	}
	l = l - 1
	return s[:l], s[l]
}
