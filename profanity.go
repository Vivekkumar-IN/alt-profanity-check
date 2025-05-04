package profanity

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type Output struct {
	Predictions   []int     `json:"predictions"`
	Probabilities []float64 `json:"probabilities"`
}

func Predict(text string) (int, error) {
	cmd := exec.Command("./predict", text)
	out, err := cmd.Output()
	if err != nil {
		return 0, fmt.Errorf("predict error: %w", err)
	}

	var result Output
	if err := json.Unmarshal(out, &result); err != nil {
		return 0, fmt.Errorf("json parse error: %w", err)
	}

	return result.Predictions[0], nil
}

func PredictProb(text string) (float64, error) {
	cmd := exec.Command("./predict", text)
	out, err := cmd.Output()
	if err != nil {
		return 0, fmt.Errorf("predict_prob error: %w", err)
	}

	var result Output
	if err := json.Unmarshal(out, &result); err != nil {
		return 0, fmt.Errorf("json parse error: %w", err)
	}

	return result.Probabilities[0], nil
}
