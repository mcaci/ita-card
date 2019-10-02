package set

// Move moves cards from first to second set in parameters
func Move(from, to *Cards) {
	to.Add(*from...)
	from.Clear()
}
