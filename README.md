# BE-Test
V 1.0.1

BE-Test Currency Converter API written on Go using Gin-Gonic Framework (https://github.com/gin-gonic/gin) and 
[GORM](https://github.com/jinzhu/gorm) as ORM Framework

## Run With Docker
run following command on your terminal, make sure you have installed docker and docker-compose (https://docs.docker.com/compose/install/)
`docker-compose up`

## API

DIAGNOSTIC
http://localhost:8080/api/v1/ping -- to check server status

http://localhost:8080/api/v1/exchange -- to register exchange data 
Payload example : {"from":"USD", "to":"GBP"}
from : string
to : string


http://localhost:8080/api/v1/daily -- to register exchange daily rate 
Payload example : {"date":"2018-07-12","from":"USD","To":"GBP","rate":"0.75709"}
date : string
from : string
to : string

http://localhost:8080/api/v1/last7 -- to see the exchange rate trend from the most recent 7 data points
Payload example : {"from":"USD","to":"GBP","date":"2018-08-01"}
date : string
from : string
to : string


http://localhost:8080/api/v1/tracked -- to see list of exchange rates to be tracked
Payload example : {"Date":"2018-08-01", "Exchanges":[{"From":"USD","To":"GBP"},{"From":"USD","To":"IDR"},{"From":"JPY","To":"IDR"}]}
date : string
Exchanges : Array of Object 
 from : string
 to : string


 http://localhost:8080/api/v1/remove -- to see list of exchange rates to be tracked
Payload example : {""Exchanges":[{"From":"USD","To":"GBP"},{"From":"USD","To":"IDR"},{"From":"JPY","To":"IDR"}]}
Exchanges : Array of Object 
 from : string
 to : string