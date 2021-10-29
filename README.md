# Request Search Transactions

Client to generate a list of transactions with valid parameters for the endpoint
/transactions/search

# Parametros validos

  -a string
        Aquarier Name eg, genova, getnet, etc
  -f string
        Start time, eg 1m, 1d (default "1M")
  -i string
        Offset, record number where you want to start (default "0")
  -l string
        Limit, Max records that you want to show (default "30")
  -o string
        Type of operation, eg authorization, online_purchase, capture 
  -op string
        Optional values for add to query
  -s string
        Status of Operations, eg contingency, approved, rejected
  -t string
        End time, eg 1m, 1d or now (default "now")
  -token string
        Fury token

## Example
go run main.go -a getnet -o online_purchase -s approved -l 300  -token xxxxxxxxxxxxx > salida.txt
