# Request Search Transactions

Client to generate a list of transactions with valid parameters for the endpoint
/transactions/search

# Parametros validos

-  -a 
        Aquarier Name eg, genova, getnet, etc
-  -f 
        Start time, eg 1m, 1d (default "1M")
-  -i 
        Offset, record number where you want to start (default "0")
-  -l 
        Limit, Max records that you want to show (default "30")
-  -o 
        Type of operation, eg authorization, online_purchase, capture 
-  -op 
        Optional values for add to query
-  -s 
        Status of Operations, eg contingency, approved, rejected
-  -t 
        End time, eg 1m, 1d or now (default "now")
-  -token 
        Fury token

## Example
go run main.go -a getnet -o online_purchase -s approved -l 300  -token xxxxxxxxxxxxx > salida.txt
