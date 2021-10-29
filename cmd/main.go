package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"request_search_transacciones/models"
)

func main() {

	var dataquery map[string]*string
	dataquery = make(map[string]*string)
	dataquery["acquirer"] = flag.String("a", "", "Aquarier Name eg, genova, getnet, etc")
	dataquery["operation"] = flag.String("o", "", "Type of operation, eg authorization, online_purchase, capture ")
	dataquery["status"] = flag.String("s", "", "Status of Operations, eg contingency, approved, rejected")
	dataquery["offset"] = flag.String("i", "0", "Offset, record number where you want to start")
	dataquery["limit"] = flag.String("l", "30", "Limit, Max records that you want to show")
	dataquery["created_date_from"] = flag.String("f", "1M", "Start time, eg 1m, 1d")     //from
	dataquery["created_date_to"] = flag.String("t", "now", "End time, eg 1m, 1d or now") //to
	dataquery["token"] = flag.String("token", "", "Fury token")
	dataquery["optional"] = flag.String("op", "", "Optional values for add to query")
	flag.Parse()
	//fmt.Println(*dataquery["aquirer"])
	cont := 0
	query := ""

	for i := range dataquery {
		igual := "="
		if len(*dataquery[i]) > 0 {
			if i == "token" {
				continue
			}
			cont++
			if len(query) > 0 && i != "optional" {
				query += "&"
			}
			if i == "optional" {
				i = ""
				igual = ""
			}
			query += i + igual + *dataquery[i]

		}
	}
	if len(*dataquery["token"]) > 0 {
		apicall(query, *dataquery["token"])
	} else {
		fmt.Println("Error! The fury token is required")
		return
	}

}

func apicall(query string, token string) {
	//url := "https://prod_gateway-apitransactions.furyapps.io/gateway/transactions/search?acquirer=genova&operation=authorization&status=contingency&offset=0&limit=1000&creation_date_from=now-1M&creation_date_to=now&options=null" + *acquirer + query
	url := "https://prod_gateway-apitransactions.furyapps.io/gateway/transactions/search?" + query
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("X-auth-token", token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	searchResult := models.SearchResult{}
	if err = json.NewDecoder(res.Body).Decode(&searchResult); err != nil {
		fmt.Println("Error en el formateo de datos")
	} else {
		fmt.Printf("******** SON %v TRANSACCIONES ******* \n", searchResult.Paging.Total)
		for _, valor := range searchResult.Data {
			fmt.Println(valor.ID)
		}

	}

}
