package response

// Blog struct
type Blog struct {
	ID     uint   `json:"id"`
	Source string `json:"source"`
	Title  string `json:"title"`
	Link   string `json:"link"`
	Date   string `json:"date"`
}

// BlogPaginationResult response model
type BlogPaginationResult struct {
	Total int64  `json:"total"`
	Page  int    `json:"page"`
	Size  int    `json:"size"`
	List  []Blog `json:"list"`
}
