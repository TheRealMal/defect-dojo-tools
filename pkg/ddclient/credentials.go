package ddclient

type Client struct {
	ApiToken string
	ApiURL   string
}

type ResponseOnlyWithID struct {
	ID     int `json:"id"`
	TestID int `json:"test_id"`
}

type FindResponse struct {
	Count   int                  `json:"count"`
	Results []ResponseOnlyWithID `json:"results"`
}

type DDClient interface {
	CreateProduct(Product) ([]byte, error)
	FindProduct(string) (int, error)
	CreateEngagement(Engagement) (int, error)

	UploadScanReport(int, string, string) (int, error)
}
