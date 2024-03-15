package handler

import (
	"encoding/json"
	"errors"
	"github.com/Warh40k/vk-intern-filmotecka/internal/domain"
	"github.com/go-playground/validator/v10"
	"log/slog"
	"net/http"
	"strconv"
)

func (h *Handler) ListActors(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")

}

func (h *Handler) CreateActor(w http.ResponseWriter, r *http.Request) {
	const method = "Handlers.Actor.CreateActor"
	log := h.log.With(
		slog.String("method", method),
	)
	var actor domain.Actor
	err := json.NewDecoder(r.Body).Decode(&actor)
	if err != nil {
		newErrResponse(log, w, http.StatusBadRequest, r.RequestURI, "json parse error",
			"Failed to parse json. Please, check your input", err.Error())
		return
	}
	validate := validator.New()
	err = validate.Struct(actor)
	if err != nil {
		var vErr validator.ValidationErrors
		errors.As(err, &vErr)
		newErrResponse(log, w, http.StatusBadRequest, r.RequestURI, "Validation error",
			"Couldn't validate input fields. Please, fix input and try again", vErr.Error())
		return
	}

	actor.Id, err = h.services.CreateActor(actor)
	if err != nil {
		newErrResponse(log, w, http.StatusInternalServerError, r.RequestURI, "Create actor error",
			"Failed to create new actor. Please, try again later", err.Error())
		return
	}

	resp, _ := json.Marshal(actor)
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}

func (h *Handler) DeleteActor(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")

}

func (h *Handler) PatchActor(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) UpdateActor(w http.ResponseWriter, r *http.Request) {
	const method = "Handlers.Actor.CreateActor"
	log := h.log.With(
		slog.String("method", method),
	)

	var actor domain.Actor
	var err error
	actor.Id, err = strconv.Atoi(r.PathValue("actor_id"))
	if err != nil {
		newErrResponse(log, w, http.StatusBadRequest, r.RequestURI, "error getting actor id",
			"failed to get user it. Please, check your input and try again", err.Error())
		return
	}
	err = json.NewDecoder(r.Body).Decode(&actor)
	if err != nil {
		newErrResponse(log, w, http.StatusBadRequest, r.RequestURI, "json parse error",
			"Failed to parse json. Please, check your input", err.Error())
		return
	}
	validate := validator.New()
	err = validate.Struct(actor)
	if err != nil {
		var vErr validator.ValidationErrors
		errors.As(err, &vErr)
		newErrResponse(log, w, http.StatusBadRequest, r.RequestURI, "validation error",
			"Couldn't validate input fields. Please, fix input and try again", vErr.Error())
		return
	}

	err = h.services.PatchActor(actor)
	if err != nil {
		newErrResponse(log, w, http.StatusInternalServerError, r.RequestURI, "error updating actor",
			"Failed to update actor, try again later", err.Error())
		return
	}

	resp, _ := json.Marshal(actor)
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
