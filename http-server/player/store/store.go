package store

type PlayerStore interface {
	GetPlayerScore(name string) int
}

func GetPlayerScore(name string)(score int){
	if name == "A" {
		return	20
	}

	if name == "B" {
		return 10
	}

	return 
}