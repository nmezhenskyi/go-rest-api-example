package webserver

import (
	"encoding/json"
	"local/winery/pkg/wine"
	"net/http"
	"os"
)

func (s *Server) PopulateWithData(file string) error {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	var records = []wine.Wine{}
	err = json.Unmarshal(bytes, &records)
	if err != nil {
		return err
	}

	for _, rec := range records {
		s.storage.Save(rec.ID, rec)
	}

	return nil
}

func (s *Server) RemoveData() {
	s.storage.SetEmpty()
}

func sendJSON(w http.ResponseWriter, bytes []byte, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	w.Write(bytes)
}
