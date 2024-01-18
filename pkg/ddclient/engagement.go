package ddclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

type Engagement struct {
	Name                       string   `json:"name"`
	Description                string   `json:"description"`
	Tags                       []string `json:"tags,omitempty"`
	Version                    string   `json:"version,omitempty"`
	FirstContacted             string   `json:"first_contacted,omitempty"`
	TargetStart                string   `json:"target_start,omitempty"`
	TargetEnd                  string   `json:"target_end,omitempty"`
	Reason                     string   `json:"reason,omitempty"`
	Tracker                    string   `json:"tracker,omitempty"`
	TestStrategy               string   `json:"test_strategy,omitempty"`
	ThreatModel                bool     `json:"threat_model,omitempty"`
	ApiTest                    bool     `json:"api_test,omitempty"`
	PenTest                    bool     `json:"pen_test,omitempty"`
	CheckList                  bool     `json:"check_list,omitempty"`
	Status                     string   `json:"status,omitempty"`
	EngagementType             string   `json:"engagement_type,omitempty"`
	BuildID                    string   `json:"build_id,omitempty"`
	CommitHash                 string   `json:"commit_hash,omitempty"`
	BranchTag                  string   `json:"branch_tag,omitempty"`
	SourceCodeManagementURI    string   `json:"source_code_management_uri,omitempty"`
	DeduplicationOnEngagement  bool     `json:"deduplication_on_engagement,omitempty"`
	Lead                       string   `json:"lead,omitempty"`
	Requester                  string   `json:"requester,omitempty"`
	Preset                     string   `json:"preset,omitempty"`
	ReportType                 string   `json:"report_type,omitempty"`
	Product                    string   `json:"product,omitempty"`
	BuildServer                string   `json:"build_server,omitempty"`
	SourceCodeManagementServer string   `json:"source_code_management_server,omitempty"`
	OrchestrationEngine        string   `json:"orchestration_engine,omitempty"`
}

// Creates DefectDojo engagement with provided data.
//
// You can skip error handling and check
// if returned id != -1
func (ddClient *Client) CreateEngagement(engagementData Engagement) (int, error) {
	client := &http.Client{}
	bodyData, _ := json.Marshal(engagementData)
	req, err := http.NewRequest("POST", ddClient.ApiURL+"/api/v2/engagements/", bytes.NewBuffer(bodyData))
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
	fmt.Println(string(bodyText))
	var response ResponseOnlyWithID
	err = json.Unmarshal(bodyText, &response)
	if err != nil {
		return -1, err
	}
	return response.ID, nil
}

func (ddClient *Client) UploadScanReport(engagementID string, format string, filename string, closeOld string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return -1, err
	}
	defer file.Close()

	// Prepare multipart body
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	partFile, _ := writer.CreateFormFile("file", filename)
	io.Copy(partFile, file)
	partActive, _ := writer.CreateFormField("active")
	io.Copy(partActive, strings.NewReader("true"))
	partVerified, _ := writer.CreateFormField("verified")
	io.Copy(partVerified, strings.NewReader("false"))
	partScanType, _ := writer.CreateFormField("scan_type")
	io.Copy(partScanType, strings.NewReader("SARIF"))
	partEngagement, _ := writer.CreateFormField("engagement")
	io.Copy(partEngagement, strings.NewReader(engagementID))
	partCloseOld, _ := writer.CreateFormField("close_old_findings")
	io.Copy(partCloseOld, strings.NewReader(closeOld))
	writer.Close()

	client := &http.Client{}
	req, err := http.NewRequest("POST", ddClient.ApiURL+"/api/v2/import-scan/", body)
	if err != nil {
		return -1, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", writer.FormDataContentType())
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
	return response.TestID, nil
}
