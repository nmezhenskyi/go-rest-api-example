package webserver

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/nmezhenskyi/go-rest-api-example/pkg/router"
	"github.com/nmezhenskyi/go-rest-api-example/pkg/wine"
)

func (s *Server) handleWineGetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		s.Logger.Printf("Received %s %q", req.Method, req.URL.Path)

		wineRecords, ok := s.storage.FindAll()
		if !ok {
			http.Error(w, "Wine not found", http.StatusNotFound)
			return
		}

		bytes, err := json.Marshal(wineRecords)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		sendJSON(w, bytes, 200)
	}
}

func (s *Server) handleWineGetOne() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		s.Logger.Printf("Received %s %q", req.Method, req.URL.Path)

		wineId := router.URLParam(req, "id")
		wineRecord, ok := s.storage.FindById(wineId)
		if !ok {
			http.Error(w, "Wine not found", http.StatusNotFound)
			return
		}

		bytes, err := json.Marshal(wineRecord)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		sendJSON(w, bytes, 200)
	}
}

func (s *Server) handleWineCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		s.Logger.Printf("Received %s %q", req.Method, req.URL.Path)

		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		wineRecord := new(wine.Wine)
		err = json.Unmarshal(body, wineRecord)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, ok := s.storage.FindById(wineRecord.ID)
		if ok {
			http.Error(w, "Wine with this ID already exists", http.StatusBadRequest)
			return
		}

		s.storage.Save(wineRecord.ID, wineRecord)
		sendJSON(w, body, http.StatusCreated)
	}
}

func (s *Server) handleWineUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		s.Logger.Printf("Received %s %q", req.Method, req.URL.Path)

		wineId := router.URLParam(req, "id")
		wineRecord, ok := s.storage.FindById(wineId)
		if !ok {
			http.Error(w, "Wine not found", http.StatusNotFound)
			return
		}

		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		update := new(wine.Wine)
		err = json.Unmarshal(body, update)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		newRecord := wineRecord.(wine.Wine)

		if update.Name != "" {
			newRecord.Name = update.Name
		}
		if update.Category != "" {
			newRecord.Category = update.Category
		}
		if update.Label != "" {
			newRecord.Label = update.Label
		}
		if update.Volume != "" {
			newRecord.Volume = update.Volume
		}
		if update.Region != "" {
			newRecord.Region = update.Region
		}
		if update.Producer != "" {
			newRecord.Producer = update.Producer
		}
		if update.Year != 0 {
			newRecord.Year = update.Year
		}
		if update.Alcohol != "" {
			newRecord.Alcohol = update.Alcohol
		}
		if update.Price != "" {
			newRecord.Price = update.Price
		}

		s.storage.Save(wineId, newRecord)

		var resData []byte
		resData, err = json.Marshal(newRecord)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		sendJSON(w, resData, http.StatusOK)
	}
}

func (s *Server) handleWineDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		s.Logger.Printf("Received %s %q", req.Method, req.URL.Path)

		wineId := router.URLParam(req, "id")
		_, ok := s.storage.FindById(wineId)
		if !ok {
			http.Error(w, "Wine not found", http.StatusNotFound)
			return
		}

		s.storage.Remove(wineId)
		w.WriteHeader(http.StatusOK)
	}
}
