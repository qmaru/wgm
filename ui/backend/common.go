package backend

import (
	"strconv"
)

func JsonData(status int, message string, data any) map[string]any {
	return map[string]any{
		"status":  status,
		"message": message,
		"data":    data,
	}
}

func IDtoInt(id string) (int, error) {
	idN, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}
	return idN, nil
}
