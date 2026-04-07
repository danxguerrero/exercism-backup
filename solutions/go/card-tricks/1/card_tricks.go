package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	return []int{2,6,9}
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
    var value int
	if index < 0 || index > len(slice) - 1 {
        return -1
    }
    for i, v := range slice {
        if i == index {
           value = v 
        }
    }
	return value
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if index < 0 || index > len(slice) - 1 {
        slice = append(slice, value)
    } else {
    	slice[index] = value
    }

    return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
    var newSlice []int
	for _, v := range values {
        newSlice = append(newSlice, v)
    }

    combinedSlice := append(newSlice, slice...)
    return combinedSlice
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
    var newSlice []int
    if index < 0 || index > len(slice) - 1 {
        newSlice = slice
    } else {
		newSlice = append(slice[:index], slice[(index+1):]...)
    }
    return newSlice
}
