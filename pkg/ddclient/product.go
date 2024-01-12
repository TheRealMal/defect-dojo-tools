package ddclient

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type Product struct {
	Name                          string   `json:"name"`
	Description                   string   `json:"description"`
	Tags                          []string `json:"tags,omitempty"`
	BusinessCriticality           string   `json:"business_criticality,omitempty"`
	Platform                      string   `json:"platform,omitempty"`
	Lifecycle                     string   `json:"lifecycle,omitempty"`
	Origin                        string   `json:"origin,omitempty"`
	UserRecords                   int      `json:"user_records,omitempty"`
	Revenue                       string   `json:"revenue,omitempty"`
	ExternalAudience              bool     `json:"external_audience,omitempty"`
	InternetAccessible            bool     `json:"internet_accessible,omitempty"`
	EnableProductTagInheritance   bool     `json:"enable_product_tag_inheritance,omitempty"`
	EnableSimpleRiskAcceptance    bool     `json:"enable_simple_risk_acceptance,omitempty"`
	EnableFullRiskAcceptance      bool     `json:"enable_full_risk_acceptance,omitempty"`
	DisableSlaBreachNotifications bool     `json:"disable_sla_breach_notifications,omitempty"`
	ProductManager                int      `json:"product_manager,omitempty"`
	TechnicalContact              int      `json:"technical_contact,omitempty"`
	TeamManager                   int      `json:"team_manager,omitempty"`
	ProdType                      string   `json:"prod_type"`
	SlaConfiguration              string   `json:"sla_configuration"`
	Regulations                   []int    `json:"regulations,omitempty"`
}

// Creates DefectDojo product with provided data
func (ddClient *Client) CreateProduct(productData Product) ([]byte, error) {
	client := &http.Client{}
	bodyData, _ := json.Marshal(productData)
	req, err := http.NewRequest("POST", ddClient.ApiURL+"/api/v2/products/", bytes.NewBuffer(bodyData))
	if err != nil {
		return []byte{}, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", ddClient.ApiToken)
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return bodyText, nil
}
