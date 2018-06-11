package render

type (
	renderList []renderComponent
)

func (rl *renderList) add(items ...renderComponent) {
	*rl = append(*rl, items...)
}

func (rl *renderList) remove(items ...renderComponent) {
	for _, c := range items {
		rl.removeOne(c)
	}
}

func (rl *renderList) removeOne(c renderComponent) {
	rlSlice := *rl
	for i, v := range rlSlice {
		if v == c {
			// ensure memory is release, even if
			// we don't reallocate
			rlSlice[i] = nil

			switch {
			case i == len(rlSlice)-1:
				rlSlice = rlSlice[:len(rlSlice)-2]
			default:
				rlSlice = append(rlSlice[:i], rlSlice[i+1:]...)
			}
		}
	}
}
