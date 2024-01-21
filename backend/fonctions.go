package backend

func SupprimerAventurierParID(id int, aventuriers *[]Aventurier) bool {
	index := -1
	for i, a := range *aventuriers {
		if a.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		return false
	}

	*aventuriers = append((*aventuriers)[:index], (*aventuriers)[index+1:]...)

	return true
}
