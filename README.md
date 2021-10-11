#### Get started
To make use of this application, clone this repository and open the prowarehouse-pokemon executable. Alternatieve you can navigate to the directory by command-line shell and than type in `./prowarehouse-pokemon`. 
A Api Endpoint will be created on the port 8080, than using Postman you can send a request on this url `localhost:8080/`. This will return  a list of  pokemon with their name, height, weight, moves and types.

There is also a Api Endpoint on `localhost:8080/resetdb`. This will empty the table data from the pokemon table and refill it with 500 pokemon receive from the pokeapi. (Note: this will take awhile before it's done)

#### Pagination
To make use of pagination, you only need to add the following query params in the url `?page=1&limit=10`.
The `page` param is for the page you want to see and the `limit` param is number of record you want to see in the list.
The default for both parameter are `Page 1` and `Limit 10`.

