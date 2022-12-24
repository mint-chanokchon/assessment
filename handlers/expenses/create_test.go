package expenses_test

import (
	"bytes"
	"net/http"
	"os"
	"testing"
)

var url = "http://localhost" + os.Getenv("PORT")

func TestCreate(t *testing.T) {
	statusCodeCases := []StatusCodeCase{
		{Name: "should return StatusCreated",
			Body: `{ "title": "buy a new phone", "amount": 39000, "note": "buy a new phone", "tags": ["gadget", "shopping"] }`,
			Want: http.StatusCreated},
		{Name: "should return BadRequest",
			Body: `{  }`,
			Want: http.StatusBadRequest},
	}

	for _, testCase := range statusCodeCases {
		t.Run(testCase.Name, func(t *testing.T) {
			body := bytes.NewBufferString(testCase.Body)

			var expenses Expense
			res := request(http.MethodPost, url+"/expenses", body)
			res.Decode(&expenses)

			if res.StatusCode != testCase.Want {
				t.Errorf("StatusCode should be: %d, but %d", http.StatusBadRequest, res.StatusCode)
			}
		})
	}
}
