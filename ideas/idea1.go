package ideas

type MarkedList map[int]*Item

type BoardV2 struct {
	rows []*Row
	in   chan byte
}

func (b *BoardV2) run() {
	go func() {
		//TODO: impl
	}()
}

func (b *BoardV2) IsWinner() bool {
	//O(log(n))
	return true
}

type Row struct {
	items  []Item
	b      byte
	marked bool
}

type Item struct {
	i      int
	marked bool
}
