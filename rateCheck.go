package easyparcel

import (
	"net/http"
  "io/ioutil"
  "bytes"
  "encoding/json"
  models "github.com/helmiruza/easyparcel-go/models"
)

func GetRates(data models.BulkRateCheck) (models.BulkRateCheckResponse, models.Error) {
  URL += "/?ac=MPRateCheckingBulk"
  requestBody, _ := json.Marshal(data)

  client := &http.Client{}

  req, _ := http.NewRequest("POST", URL, bytes.NewBuffer(requestBody))
  req.Header.Set("Content-type", "application/json")

  resp, _ := client.Do(req)
  body, err := ioutil.ReadAll(resp.Body)

  obj, err1 := response(body, err, "rateCheck")
  objString, _ := json.Marshal(obj)
  k := models.BulkRateCheckResponse{}
  json.Unmarshal(objString, &k)

  return k, err1
}
