package plans

type User struct {
	UserId  string `csv:"uid" json:"UserId"`
	Name    string `csv:"name" json:"Name"`
	Address string `csv:"address" json:"Address"`
	ZipCode string `json:"-"`
}
