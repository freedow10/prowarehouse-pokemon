To make use of this application clone this repository and open the prowarehouse-pokemon executable. Alternatieve you can navigate to the directory by command-line shell and than type in `./prowarehouse-pokemon`. 
A Api Endpoint will be created on the port 8080, than using Postman you can send a request on this url `localhost:8080/`. To receive a list of 500 pokemon with their name, height, weight, moves and types. 

#### Pagination
To make use of pagination, you only need to add the following query params in the url `?page=1&limit=10`
`page` param is for the page you want to see and `limit` is number of record you want to see in the list.
The default for both parameter are `Page 1` and `Limit 10`.