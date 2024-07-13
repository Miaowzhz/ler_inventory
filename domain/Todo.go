package domain

type Todo struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Status   bool   `json:"status"`
	Favorite bool   `json:"favorite"` // 新增字段
}
