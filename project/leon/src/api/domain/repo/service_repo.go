package repo

type ServiceRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ServiceRepoResponse struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Owner string `json:"owner"`
}
