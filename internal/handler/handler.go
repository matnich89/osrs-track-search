package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"osrs-track-search/internal/client"
	"osrs-track-search/internal/process"
)

type Handler struct {
	client    client.Client
	processor process.SkillsProcessor
}

func NewHandler(client client.Client) *Handler {
	return &Handler{client: client}
}

func (h *Handler) SearchIronman(w http.ResponseWriter, r *http.Request) {
	character, err := getCharacter(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := h.client.GetHighScores(character, client.Ironman)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
