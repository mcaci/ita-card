package set

func Move(from, to *Cards) {
	to.Add(*from...)
	from.Clear()
}
