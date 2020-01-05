package newsfeed
//Getter interface to program against
type Getter interface {
	GetAll() []Item
}
//Adder Interface to program against
type Adder interface {
	Add(item Item)
}
//Item is for our News items
type Item struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}
//Repo for holding news items
type Repo struct {
	Items []Item
}
//New Repo (empty) of news items
func New() *Repo {
	return &Repo{
		Items: []Item{},
	}
}
//Add impl
func (r *Repo) Add(item Item) {
	r.Items = append(r.Items, item)
}
//GetAll impl
func (r *Repo) GetAll() []Item {
	return r.Items
}
