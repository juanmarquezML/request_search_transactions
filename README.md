# Request Search Transactions

Cliente para generar un listado de id de transacciones, segun los parametros que se le establezcan

## Parametros validos

-  [-a] Nombre del Adquirente, ejemplo: genova, getnet, etc
-  [-o] Tipo de Operación, ejemplo: authorization, online_purchase, capture 
-  [-s] Status de la Operación, ejemplo: contingency, approved, rejected
-  [-f] Tiempo desde que se quiere iniciar la busqueda, ejemplo: 2M, 1d (default "1M")
-  [-t] Tiempo hasta, ejemplo: 2d, 1M  (default "now")
-  [-i] Offset, número de registro donde quiere inicia la busqueda
-  [-l] Limit, número maximo de registros que desea mostrar en el resultado (default "30")
-  [-op] Con esta parametro se puede pasar cualquier query string que quiere agregar al request, Ejemplo &regulations=null
-  [-token] Fury token, este parametro es obligatorio

## Ejemplo

go run main.go -a getnet -o online_purchase -s approved -l 300  -token xxxxxxxxxxxxx > salida.txt

El comando anterior está buscando las transacciones para el adquirente Getnet, con operaciones de online_purchase, que tengan estado aprobados y se limitó a un total de 300 registros, y por defecto como no le pasamos opciones se limitó a una búsqueda de un mes atras, ademas para sacar los resultado a una archivo puede agregar la típica opción " > nombre_archivo.txt o la extensión que prefiera"
