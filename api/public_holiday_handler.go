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

type PublicHolidayHandler struct {
	service service.PublicHolidayService
}

func NewPublicHolidayHandler(service service.PublicHolidayService) *PublicHolidayHandler {
	return &PublicHolidayHandler{service}
}

func (h *PublicHolidayHandler) SyncAllCountries(w http.ResponseWriter, r *http.Request) {
	err := h.service.SyncAllCountries()
	if err != nil {
		util.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	util.WriteSuccessResponse(w, "Public holidays for all countries synced successfully")
}

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

func (h *PublicHolidayHandler) GetAllPublicHolidays(w http.ResponseWriter, r *http.Request) {
	holidays, err := h.service.GetAllPublicHolidays()
	if err != nil {
		util.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	util.WriteSuccessResponse(w, holidays)
}

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
