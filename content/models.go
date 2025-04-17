package content

type Article struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content"`
}

type Blog struct {
	ID        string `json:"id"`
	Headline  string `json:"headline"`
	Writer    string `json:"writer"`
	Body      string `json:"body"`
	Published bool   `json:"published"`
}
