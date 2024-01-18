package ddclient

type Client struct {
	ApiToken string
	ApiURL   string
}

type ResponseOnlyWithID struct {
	ID int `json:"id"`
}

type DDClient interface {
	CreateProduct(Product) ([]byte, error)
	FindProduct(string) (int, error)
}
