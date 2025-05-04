package structs

type ApiResponse struct {
	Items    []Item `json:"items"`
	NextPage string `json:"next_page"`
}
