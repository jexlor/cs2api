package api

// structure of skins
type Skin struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Rarity     string `json:"rarity"`
	Collection string `json:"collection"`
	Quality    string `json:"quality"`
	Price      string `json:"price"` //type string in order to handle currency symbols
	Url        string `json:"url"`
}

type Col struct {
	Collection string `json:"collection"`
}
