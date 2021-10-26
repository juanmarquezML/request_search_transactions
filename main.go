package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Transaction struct {
	ID                string `json:"id"`
	ProfileID         string `json:"profile_id"`
	Acquirer          string `json:"acquirer"`
	InternalProfileID string `json:"internal_profile_id"`
	PaymentID         string `json:"payment_id"`
	Status            string `json:"status"`
	Operation         string `json:"operation"`
	ProcessingMode    string `json:"processing_mode"`
	Refunds           struct {
		Detail []struct {
			ID     uint64 `json:"id"`
			Status string `json:"status"`
		} `json:"detail"`
	} `json:"refunds"`
	CreationDate     time.Time `json:"creation_date"`
	LastModifiedDate time.Time `json:"last_modified_date"`
}

type Payment struct {
	TransactionID string `json:"transaction_id"`
}

/*type SearchOpts struct {
	AuthToken                   string
	Status                      string
	Operation                   string
	Acquirer                    string
	CaptureExternalResponseCode string
	FromDate                    string
	ToDate                      string
	Offset                      uint64
	Limit                       uint64
	Verbose                     bool
}

type DupResponse struct {
	ID    string `json:"id"`
	Value struct {
		Token string `json:"token"`
	} `json:"value"`
}*/

type SearchResult struct {
	Data   []Transaction `json:"data"`
	Paging struct {
		Total  uint64 `json:"total"`
		Limit  uint64 `json:"limit"`
		Offset string `json:"offset"`
	} `json:"paging"`
}

func main() {

	url := "https://prod_gateway-apitransactions.furyapps.io/gateway/transactions/search?acquirer=genova&operation=authorization&status=contingency&offset=0&limit=1000&creation_date_from=now-2M&creation_date_to=now&options=null"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("X-auth-token", "d2ec822e41a7e3af0661189ba9b229f168016115d997b40de0372f9861f6bae1")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	searchResult := SearchResult{}
	if err = json.NewDecoder(res.Body).Decode(&searchResult); err != nil {
		fmt.Println("Error en el formateo de datos")
	} else {
		fmt.Printf("******** SON %v TRANSACCIONES ******* \n", searchResult.Paging.Total)
		for _, valor := range searchResult.Data {
			fmt.Println(valor.ID)
		}

	}

}
