package expenses_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"testing"
)

var url = "http://localhost" + os.Getenv("PORT")

type Response struct {
	*http.Response
	err error
}

type Expense struct {
	Id     string   `json:"id"`
	Title  string   `json:"title"`
	Amount int      `json:"amount"`
	Note   string   `json:"note"`
	Tags   []string `json:"tags"`
}

type Case struct {
	Name string
	Body string
	Want int
}

func TestCreate(t *testing.T) {
	testCases := []Case{
		{
			Name: "should return StatusCreated",
			Body: `{ "title": "buy a new phone", "amount": 39000, "note": "buy a new phone", "tags": ["gadget", "shopping"] }`,
			Want: http.StatusCreated},
	}

	for _, testCase := range testCases {
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

func (r *Response) Decode(v interface{}) error {
	if r.err != nil {
		return r.err
	}
	return json.NewDecoder(r.Body).Decode(v)
}

func request(method, url string, body io.Reader) *Response {
	req, _ := http.NewRequest(method, url, body)
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	res, err := client.Do(req)
	return &Response{res, err}
}
