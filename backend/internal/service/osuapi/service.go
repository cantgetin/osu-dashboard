package osuapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getUser(w http.ResponseWriter, r *http.Request) {
	access_token := r.FormValue("access_token")
	url := "https://osu.ppy.sh/api/v2/me/osu"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", access_token))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	apiResponse := make(map[string]interface{})
	err = json.NewDecoder(resp.Body).Decode(&apiResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(resp.StatusCode)
	json.NewEncoder(w).Encode(apiResponse)
}
