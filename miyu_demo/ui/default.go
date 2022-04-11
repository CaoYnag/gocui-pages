package ui

var (
	_cur CUI
	_nxt CUI
)

func Run(start CUI) error {
	_nxt = start
	for _nxt != nil {
		_cur, _nxt = _nxt, nil
		if e := _cur.Init(); e != nil {
			return e
		}
		if e := _cur.Run(); e != nil {
			return e
		}
		_cur.Release()
	}
	return nil
}

func _jump_to(nxt CUI) {
	_nxt = nxt
}
