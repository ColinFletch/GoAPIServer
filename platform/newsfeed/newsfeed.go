package newsfeed
//Getter gets news items
	Add(ite Item)
}
//Item struct for news items
	Post  string `json:"post"`
t}
//Create new Repo
func New() *Repo {
	return &Repo{
//Add items to a repo
func (r *Repo) Add(item Item) {
	r.Items = append(r.Items, item)
	return r.Items
}	
