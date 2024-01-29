package handler

import (
	"encoding/json"
	"errors"
	"github.com/matnich89/osrs-track-search/internal/client"
	"github.com/matnich89/osrs-track-search/internal/process"
	"net/http"
)

type Handler struct {
	client    client.Client
	processor *process.SkillsProcessor
}

func NewHandler(client client.Client) *Handler {
	return &Handler{client: client, processor: process.NewProcessor()}
}

func (h *Handler) SearchIronman(w http.ResponseWriter, r *http.Request) {
	character, err := getCharacter(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := h.client.GetHighScores(character, client.Ironman)

	if err != nil {
		if errors.Is(err, client.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	stats, err := h.processor.Process(character, resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	b, err := json.Marshal(stats)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(b)

	return
}

func getCharacter(r *http.Request) (string, error) {
	character := r.URL.Query().Get("character")

	if character == "" {
		return "", errors.New("missing character param")
	}

	return character, nil
}
