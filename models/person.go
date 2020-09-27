package models

type Person struct{
	User string `json:"name"`
	Birthday int `json:"birthday"`
	Address string `json:"address"`
	nick string `json:"nick"`
}
