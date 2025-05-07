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

func Predict(texts ...string) ([]int, error) {
	cmd := exec.Command("./predict", texts...)

	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("predict error: %w", err)
	}

	var result Output
	if err := json.Unmarshal(out, &result); err != nil {
		return nil, fmt.Errorf("json parse error: %w", err)
	}

	return result.Predictions, nil
}

func PredictProb(texts ...string) ([]float64, error) {
	cmd := exec.Command("./predict", texts...)

	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("predict_prob error: %w", err)
	}

	var result Output
	if err := json.Unmarshal(out, &result); err != nil {
		return nil, fmt.Errorf("json parse error: %w", err)
	}

	return result.Probabilities, nil
}
