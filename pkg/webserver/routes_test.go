package webserver

import (
	"bytes"
	"encoding/json"
	"io"
	"local/winery/pkg/wine"
	"net/http"
	"net/http/httptest"
	"testing"
)

var server *Server

func init() {
	server = NewServer()
}

func TestIndex(t *testing.T) {
	response, err := sendRequest("GET", "/", nil)
	if err != nil {
		t.Error("Failed to send GET / request")
	}
	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestWineGetAll(t *testing.T) {
	err := server.PopulateWithData("../../data/data_sample.json")
	if err != nil {
		t.Error("Failed to populate with data")
	}

	response, err := sendRequest("GET", "/api/wine", nil)
	if err != nil {
		t.Error("Failed to send GET /api/wine request")
	}
	checkResponseCode(t, http.StatusOK, response.Code)

	server.RemoveData()
}

func TestWineGetOne(t *testing.T) {
	err := server.PopulateWithData("../../data/data_sample.json")
	if err != nil {
		t.Error("Failed to populate with data")
	}

	response, err := sendRequest("GET", "/api/wine/1", nil)
	if err != nil {
		t.Error("Failed to send GET /api/wine/1 request")
	}
	checkResponseCode(t, http.StatusOK, response.Code)

	server.RemoveData()
}

func TestWineCreate(t *testing.T) {
	record := wine.Wine{
		ID:       "10",
		Name:     "Example Name",
		Category: "Example Category",
		Label:    "Example Label",
		Volume:   "0.7",
		Region:   "Example Region",
		Producer: "Example Producer",
		Year:     2022,
		Alcohol:  "11.5%",
		Price:    "30.00 CAD",
	}

	byteData, err := json.Marshal(record)
	if err != nil {
		t.Error("Failed to parse JSON")
	}

	response, err := sendRequest("POST", "/api/wine", bytes.NewReader(byteData))
	if err != nil {
		t.Error("Failed to send POST /api/wine request")
	}
	checkResponseCode(t, http.StatusCreated, response.Code)

	server.RemoveData()
}

func TestWineUpdate(t *testing.T) {
	err := server.PopulateWithData("../../data/data_sample.json")
	if err != nil {
		t.Error("Failed to populate with data")
	}

	record := wine.Wine{
		Name:     "Example Name",
		Category: "Example Category",
		Label:    "Example Label",
		Volume:   "0.7",
		Region:   "Example Region",
		Producer: "Example Producer",
		Year:     2022,
		Alcohol:  "11.5%",
		Price:    "20.00",
	}

	byteData, err := json.Marshal(record)
	if err != nil {
		t.Error("Failed to parse JSON")
	}

	response, err := sendRequest("PUT", "/api/wine/1", bytes.NewReader(byteData))
	if err != nil {
		t.Error("Failed to send PUT /api/wine/1 request")
	}
	checkResponseCode(t, http.StatusOK, response.Code)

	server.RemoveData()
}

func TestWineDelete(t *testing.T) {
	err := server.PopulateWithData("../../data/data_sample.json")
	if err != nil {
		t.Error("Failed to populate with data")
	}

	response, err := sendRequest("DELETE", "/api/wine/1", nil)
	if err != nil {
		t.Error("Failed to send DELETE /api/wine/1 request")
	}
	checkResponseCode(t, http.StatusOK, response.Code)

	server.RemoveData()
}

func sendRequest(method, url string, body io.Reader) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	rr := httptest.NewRecorder()
	server.ServeHTTP(rr, req)

	return rr, nil
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d, got %d\n", expected, actual)
	}
}
