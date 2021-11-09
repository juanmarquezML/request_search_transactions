# Request Search Transactions
 Client  to generate lists of ID transactions in search transactions



## Usage
```
    Example usage:
    go run main.go -a getnet -o online_purchase -s approved -l 300  -token xxxxxxxxxxxxx > salida.txt
```
## Obligatory Parameters  
```
    -token          It's fury token  
```

## Optionals Parameters
```
    -a          Acquirer.                                                       Eg. genova, getnet, etc
    -o          Operation Type.                                                 Eg. authorization, online_purchase, capture
    -s          Operation Status.                                               Eg. contingency, approved, rejected
    -f          Date from.                                                      Eg. 2M, 1d (default "1M")         
    -t          Date to.                                                        Eg. 2d, 1M (default "now")
    -i          Offset, record number where you want to start the result        Eg. 10, (default "0")
    -l          Limit, max records that you want to show                        Eg. 20, (default "30")
    -op         Optional values for add to query                                Eg. 2 &regulations=null

```
