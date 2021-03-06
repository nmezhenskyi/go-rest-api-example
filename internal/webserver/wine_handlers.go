package webserver

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/nmezhenskyi/go-rest-api-example/internal/model"
	"github.com/nmezhenskyi/go-rest-api-example/internal/router"
)

func (s *Server) handleWineGetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		s.Logger.Printf("Received %s %q", req.Method, req.URL.Path)

		wineRecords, ok := s.storage.FindAll()
		if !ok {
			sendNotFound(w, "Wine not found")
			return
		}

		sendJSON(w, http.StatusOK, wineRecords)
	}
}

func (s *Server) handleWineGetOne() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		s.Logger.Printf("Received %s %q", req.Method, req.URL.Path)

		wineId := router.URLParam(req, "id")
		wineRecord, ok := s.storage.FindById(wineId)
		if !ok {
			sendNotFound(w, "Wine not found")
			return
		}

		sendJSON(w, http.StatusOK, wineRecord)
	}
}

func (s *Server) handleWineCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		s.Logger.Printf("Received %s %q", req.Method, req.URL.Path)

		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			sendInternalError(w)
			return
		}

		wineRecord := new(model.Wine)
		err = json.Unmarshal(body, wineRecord)
		if err != nil {
			sendInternalError(w)
			return
		}

		_, ok := s.storage.FindById(wineRecord.ID)
		if ok {
			sendBadRequest(w, "Wine with this ID already exists", nil)
			return
		}

		s.storage.Save(wineRecord.ID, wineRecord)
		sendJSON(w, http.StatusCreated, wineRecord)
	}
}

func (s *Server) handleWineUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		s.Logger.Printf("Received %s %q", req.Method, req.URL.Path)

		wineId := router.URLParam(req, "id")
		wineRecord, ok := s.storage.FindById(wineId)
		if !ok {
			sendNotFound(w, "Wine not found")
			return
		}

		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			sendInternalError(w)
			return
		}

		update := new(model.Wine)
		err = json.Unmarshal(body, update)
		if err != nil {
			sendInternalError(w)
			return
		}

		newRecord := wineRecord.(model.Wine)

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
		sendJSON(w, http.StatusOK, newRecord)
	}
}

func (s *Server) handleWineDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		s.Logger.Printf("Received %s %q", req.Method, req.URL.Path)

		wineId := router.URLParam(req, "id")
		_, ok := s.storage.FindById(wineId)
		if !ok {
			sendNotFound(w, "Wine not found")
			return
		}

		s.storage.Remove(wineId)
		w.WriteHeader(http.StatusOK)
	}
}
