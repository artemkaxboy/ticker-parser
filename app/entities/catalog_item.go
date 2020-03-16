package entities

type CatalogItem struct {
	Title   string `json:"title"`
	URL     string `json:"fronturl"`
	Type    string `json:"type"`
	Company struct {
		Name string `json:"title"`
		Logo string `json:"logo_link"`
		URL  string `json:"fronturl"`
	} `json:"company"`
	Price    float64 `json:"price"`
	Currency string  `json:"currency"`
}

type CatalogHTTPData struct {
	ItemsCount int           `json:"currentItemCount"`
	Items      []CatalogItem `json:"items"`
}

func NewCatalogHTTPData(data *[]CatalogItem) *CatalogHTTPData {
	return &CatalogHTTPData{
		ItemsCount: len(*data),
		Items:      *data,
	}
}
