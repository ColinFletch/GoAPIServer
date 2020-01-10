package newsfeed

//Getter gets news items
type Getter interface {
	GetAll() []Item
}
// Adder adds news items
type Adder interface {
	Add(item Item)
}
//Item struct for news items
type Item struct {
	Title string `json:"title"`,
	Post  string `json:"post"`,
	ID int `json:"id"`,
}
//Repo struct to store items
type Repo struct {
	Items []Item
}
//Create new Repo
func New() *Repo {
	return &Repo{
		Items: []Item{},
	}
}
//Add items to a repo
func (r *Repo) Add(item Item) {
	r.Items = append(r.Items, item)
}
//GetAll items from a repo
func (r *Repo) GetAll() []Item {
	return r.Items
}