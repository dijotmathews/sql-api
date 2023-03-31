package api

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sql-api/dto"
)

func NewClient(url string, method string, payload []byte, header http.Header, apiName string, i dto.Identity) ([]byte, error, int) {
	client := &http.Client{}
	if method == "GET" {
		payload = nil
	}
	// start := time.Now()
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	// var response map[string]interface{}

	if err != nil {
		fmt.Println(err)
		return nil, err, 0
	}
	req.Header = header

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err, 0
	}
	defer res.Body.Close()
	// ?=s := time.Since(start)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		// logClientResponse(string(payload), string(err.Error()), url, method, header, res.StatusCode, i, apiName, s.Milliseconds())
		return nil, err, 0
	}
	// logClientResponse(string(payload), string(body), url, method, header, res.StatusCode, i, apiName, s.Milliseconds())
	// json.Unmarshal(body, &response)
	return body, nil, res.StatusCode

}

// func logClientResponse(payload string, response string, url string, method string, header http.Header, statusCode int, i dto.Identity, api string, s int64) {
// 	var resClient string
// 	if reqHeader, err := json.Marshal(header); err != nil {
// 		resClient = "header error"
// 	} else {
// 		resClient = string(reqHeader)
// 	}

// 	db := initializers.GetDB()
// 	//Update Client entry with response
// 	cl := models.ExternalApiLog{
// 		Request:     string(payload),
// 		ServiceName: api,
// 		Response:    string(response),
// 		EndPoint:    url,
// 		Method:      method,
// 		Header:      resClient,
// 		UsersId:     i.UsersId,
// 		TimeTaken:   s,
// 		StatusCode:  statusCode,
// 	}
// 	db.Create(&cl)
// }
