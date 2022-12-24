package expenses_test

import (
	"bytes"
	"net/http"
	"testing"
)

func TestUpdate(t *testing.T) {
	t.Run("should return statusOk", func(t *testing.T) {
		id := "1"
		body := bytes.NewBufferString(`{ "title": "buy a new phone", "amount": 39000, "note": "buy a new phone", "tags": ["gadget", "shopping"] }`)

		var expenses Expense
		res := request(http.MethodPut, url+"/expenses/"+id, body)
		res.Decode(&expenses)

		if res.StatusCode != http.StatusOK {
			t.Errorf("StatusCode should be: %d, but %d", http.StatusBadRequest, res.StatusCode)
		}
	})
}
