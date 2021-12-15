# Microservices training Go

The goal of this project is to create 3 microservices, one for clients, one for invoices and one for products in Golang. Each microservice will be in it's own folder with the corresponding name.

## Invoices microservice

This endpoint will manage invoices and will interact with other microservices. It will expose the endpoint in port **5001**
The endpoints to create for this service are:

### Get all invoices

**GET**

**/invoices**

*Response:*
```
[
	{
		"Id",
		"ClientId",
		"ClientName",
		"InvoiceDate",
		"Products": [
			"ProductId",
			"ProductName",
			"Quantity",
			"Cost"
		]
	},
	...
]
```


*Special considerations*

The client name (GET /clients/{id}) and the product name (GET /products/{id}) will come from the corresponding endpoints.

### Get invoice by id

**GET**

**/invoices/{id}**

*Response:*
```
{
	"Id",
	"ClientId",
	"ClientName",
	"InvoiceDate",
	"Products": [
		"ProductId",
		"ProductName",
		"Quantity",
		"Cost"
	]
}
```

*Special considerations*

The client name (GET /clients/{id}) and the product name (GET /products/{id}) will come from the corresponding endpoints.

### Get client's debt

**GET**

**/invoices/client/{id}**

*Response:*
```
{
	"HasDebt": true 
}
```

### Get product usage

**GET**

**/invoices/product/{id}**

*Response:*
```
{
	"IsUsed": true 
}
```

### Create invoice

**POST**

**/invoices**

*Request data*
```
{
	"ClientId",
	"InvoiceDate",
	"Products": [
		"ProductId",
		"Quantity",
		"Cost"
	]
}
```

*Response:*
```
{
	"Id",
	"ClientId",
	"ClientName",
	"InvoiceDate",
	"Products": [
		"ProductId",
		"ProductName",
		"Quantity",
		"Cost"
	]
}
```

*Special considerations*

The id of the invoice will be autoincremental determined by the endpoint or storage
		
## Tips and special considerations
- Use GORM [example usage](https://github.com/mludeiro/go-micro)
- Use in memory database [example usage](https://github.com/mludeiro/go-micro)

## Next steps
- Unit tests
- Docker usage
