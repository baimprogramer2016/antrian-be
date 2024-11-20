package handlers

import (
	"be-mklinik/f"
	loketrepository "be-mklinik/repositories/loket_repository.go"
	loketservice "be-mklinik/services/loket_service"
	"net/http"

	"gorm.io/gorm"
)

type loketHandler struct {
	loketService loketservice.LoketServiceInterface
}

func NewLoketHandler(db *gorm.DB) *loketHandler {
	respository := loketrepository.NewLoketRepository(db)
	services := loketservice.NewLoketService(respository)
	return &loketHandler{loketService: services}
}

func (s *loketHandler) LoketHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()

	result, err := s.loketService.GetDataLoket()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	f.WriteToJson(w, r, result)
}
