package ddclient

type Client struct {
	ApiToken string
	ApiURL   string
}

type DDClient interface {
	CreateProduct(Product) ([]byte, error)
}
