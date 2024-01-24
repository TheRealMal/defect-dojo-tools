package ddclient

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"
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

// Creates DefectDojo product with provided data.
//
// You can skip error handling and check
// if returned id != -1
func (ddClient *Client) CreateProduct(productData Product) (int, error) {
	client := &http.Client{}
	bodyData, _ := json.Marshal(productData)
	req, err := http.NewRequest("POST", ddClient.ApiURL+"/api/v2/products/", bytes.NewBuffer(bodyData))
	if err != nil {
		return -1, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", ddClient.ApiToken)
	resp, err := client.Do(req)
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return -1, err
	}

	var response ResponseOnlyWithID
	err = json.Unmarshal(bodyText, &response)
	if err != nil {
		return -1, err
	}
	return response.ID, nil
}

// Searchs for DefectDojo product by exact name
// returns project id.
//
// You can skip error handling and check
// if returned id != -1
func (ddClient *Client) FindProduct(name string) (int, error) {
	client := &http.Client{}
	requestURL := strings.Builder{}
	requestURL.WriteString(ddClient.ApiURL)
	requestURL.WriteString("/api/v2/products/?name_exact=")
	requestURL.WriteString(url.QueryEscape(name))
	req, err := http.NewRequest("GET", requestURL.String(), &bytes.Buffer{})
	if err != nil {
		return -1, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", ddClient.ApiToken)
	resp, err := client.Do(req)
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return -1, err
	}
	var response FindResponse
	err = json.Unmarshal(bodyText, &response)
	if err != nil {
		return -1, err
	}

	if response.Count == 0 {
		return -1, errors.New("product not found")
	}
	return response.Results[0].ID, nil
}
