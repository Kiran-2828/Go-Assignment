package model 

type Characters struct {
	Name     string `json:"name"`
	MaxPower int    `json:"max_power"`
}

type Marvel struct {
	Name      string `json:"name"`
	Character []Characters
}

type FinalList struct {
	Name     string `json:"name"`
	MaxPower int    `json:"max_power"`
	Count    int
}
