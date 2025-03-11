package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"regexp"
)

type MazeRequest struct {
	Maze string `json:"maze"`
}

func main() {
	http.HandleFunc("/process", handleProcess)
	fmt.Println("API server running on :8080")
	http.ListenAndServe(":8080", nil)
}

func handleProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req MazeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	mazePath := "../maps/maze-api.txt"
	if err := ioutil.WriteFile(mazePath, []byte(req.Maze), 0644); err != nil {
		http.Error(w, fmt.Sprintf("Failed to save maze: %v", err), http.StatusInternalServerError)
		return
	}

	cmd := exec.Command("go", "run", "../game/main.go", "-m", mazePath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to execute game: %v\nOutput: %s", err, output), http.StatusInternalServerError)
		return
	}

	intersection := extractIntersection(string(output))
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(intersection))
}

func extractIntersection(output string) string {
	re := regexp.MustCompile(`Оптимальная точка сбора: \{X:(\d+) Y:(\d+)\}`)
	match := re.FindStringSubmatch(output)
	if len(match) == 3 {
		return fmt.Sprintf("Intersection: X:%s, Y:%s", match[1], match[2])
	}
	return "Intersection not found"
}
