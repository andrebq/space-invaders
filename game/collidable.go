package game

type (
	collidables []collidable
)

func (cls *collidables) add(c collidable) {
	*cls = append(*cls, c)
}

func (cls *collidables) remove(c collidable) {
	slice := *cls
	for i, v := range slice {
		if v == c {
			slice[i] = nil
			switch i {
			case len(slice) - 1:
				slice = slice[:len(slice)-1]
			default:
				slice = append(slice[:i], slice[i+1:]...)
			}
		}
	}
	*cls = slice
}
