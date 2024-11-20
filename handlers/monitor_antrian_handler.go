package handlers

import (
	"be-mklinik/f"
	monitorantrianrepository "be-mklinik/repositories/monitor_antrian_repository"
	monitorantrianservice "be-mklinik/services/monitor_antrian_service"
	"net/http"

	"gorm.io/gorm"
)

type pageMonitorAntrianHandler struct {
	monitorAntrianService monitorantrianservice.MonitorAntrianService
}

func NewPageMonitorAntrianHandler(db *gorm.DB) *pageMonitorAntrianHandler {
	repository := monitorantrianrepository.NewMonitorAntrianRepository(db)
	monitorAntrianService := monitorantrianservice.NewMonitorAntrianService(repository)
	return &pageMonitorAntrianHandler{
		monitorAntrianService: monitorAntrianService,
	}
}

func (s *pageMonitorAntrianHandler) MonitorAntrianHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	result, err := s.monitorAntrianService.GetAllMSeqnoAntrianByDay()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	f.WriteToJson(w, r, result)
}
