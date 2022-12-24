package expenses_test

import (
	"net/http"
	"testing"
)

func TestFindById(t *testing.T) {
	t.Run("should return StatusOk", func(t *testing.T) {
		id := "1"
		var expenses Expense

		res := request(http.MethodGet, url+"/expenses/"+id, nil)
		res.Decode(&expenses)

		if res.StatusCode != http.StatusOK {
			t.Errorf("StatusCode should be: %d, but %d", http.StatusBadRequest, res.StatusCode)
		}
	})
}
