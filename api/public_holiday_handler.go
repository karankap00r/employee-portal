package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/karankap00r/employee_portal/dto"
	"github.com/karankap00r/employee_portal/service"
	"github.com/karankap00r/employee_portal/storage/model"
	"github.com/karankap00r/employee_portal/util"
)

// PublicHolidayHandler is the interface for the public holiday handler
type PublicHolidayHandler struct {
	service service.PublicHolidayService
}

// NewPublicHolidayHandler creates a new public holiday handler with the given service
func NewPublicHolidayHandler(service service.PublicHolidayService) *PublicHolidayHandler {
	return &PublicHolidayHandler{service}
}

// SyncAllCountries syncs public holidays for all countries
func (h *PublicHolidayHandler) SyncAllCountries(w http.ResponseWriter, r *http.Request) {
	err := h.service.SyncAllCountries()
	if err != nil {
		util.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	util.WriteSuccessResponse(w, "Public holidays for all countries synced successfully")
}

// SyncCountries syncs public holidays for the specified countries
func (h *PublicHolidayHandler) SyncCountries(w http.ResponseWriter, r *http.Request) {
	var countries []string
	if err := json.NewDecoder(r.Body).Decode(&countries); err != nil {
		util.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	err := h.service.SyncCountries(countries)
	if err != nil {
		util.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	util.WriteSuccessResponse(w, "Public holidays for specified countries synced successfully")
}

// GetAllPublicHolidays gets all public holidays
func (h *PublicHolidayHandler) GetAllPublicHolidays(w http.ResponseWriter, r *http.Request) {
	holidays, err := h.service.GetAllPublicHolidays()
	if err != nil {
		util.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	util.WriteSuccessResponse(w, holidays)
}

// AddPublicHoliday adds a public holiday
func (h *PublicHolidayHandler) AddPublicHoliday(w http.ResponseWriter, r *http.Request) {
	var holiday model.PublicHoliday
	if err := json.NewDecoder(r.Body).Decode(&holiday); err != nil {
		util.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	err := h.service.AddPublicHoliday(holiday)
	if err != nil {
		util.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	util.WriteSuccessResponse(w, "Public holiday added successfully")
}

// UpdatePublicHoliday updates a public holiday
func (h *PublicHolidayHandler) UpdatePublicHolidayStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.WriteErrorResponse(w, http.StatusBadRequest, "Invalid holiday ID")
		return
	}

	var request dto.UpdatePublicHolidayRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		util.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.UpdatePublicHolidayStatus(id, request.Status)
	if err != nil {
		util.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	util.WriteSuccessResponse(w, "Public holiday status updated successfully")
}

// GetPublicHolidaysForNext7Days gets public holidays for the next 7 days
func (h *PublicHolidayHandler) GetPublicHolidaysForNext7Days(w http.ResponseWriter, r *http.Request) {
	country := r.URL.Query().Get("country")
	if country == "" {
		util.WriteErrorResponse(w, http.StatusBadRequest, "Country parameter is required")
		return
	}

	holidays, err := h.service.GetPublicHolidaysForNext7Days(country)
	if err != nil {
		util.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	util.WriteSuccessResponse(w, holidays)
}

// SendPublicHolidayAlert sends a public holiday alert to the specified email
func (h *PublicHolidayHandler) SendPublicHolidayAlert(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Email   string `json:"email"`
		Country string `json:"country"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		util.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if request.Email == "" || request.Country == "" {
		util.WriteErrorResponse(w, http.StatusBadRequest, "Email and Country are required")
		return
	}

	err := h.service.SendPublicHolidayAlert(request.Email, request.Country)
	if err != nil {
		util.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	util.WriteSuccessResponse(w, "Public holiday alert sent successfully")
}
