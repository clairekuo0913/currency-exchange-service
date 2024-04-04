# Currency Exchange Service
This service provides an API for converting amounts between different currencies using pre-defined exchange rates.

## Supported Currencies
Note: Currently, this service supports conversions between the following currencies only:

- USD
- JPY
- TWD

Please ensure your requests use these currency codes for the source and target parameters.

## Prerequisites
- Docker
- docker-compose

## Configuration
A `.env` file is required for configuring the service, located in the project root directory. The file should contain the following environment variable:
```
SERVER_PORT=8080
```
This variable specifies the port on which the server will listen. 
If the `SERVER_PORT` variable is not specified, the service defaults to using port 8080.
Ensure the `.env` file is in place and properly configured before building and running the service.

## Building and Running the Service
To build and run the service, use the following command:
```sh
make run
```
If you only want to build the service, the `make build` command will build the Docker image for the currency exchange service.

## Using the API
To convert a currency amount, make a GET request to the `/convert` endpoint with the following query parameters:

- `source`: Source currency code (e.g., "USD")
- `target`: Target currency code (e.g., "TWD")
- `amount`: The amount to convert. This can include commas as thousand separators (e.g., "1,223.29")

Example request:
```
http://localhost:8080/convert?source=USD&target=JPY&amount=1,525
```
or
```
http://localhost:8080/convert?source=USD&target=JPY&amount=1525
```

## API Response
The API response will be a JSON object containing the converted amount, formatted with two decimal places and commas as thousand separators.

Example response for the above request might look like this (assuming the conversion rate is 111.801):
```
{
  "msg": "success",
  "amount": "170496.53"
}
```

## License
This project is licensed under the MIT License - see the LICENSE.md file for details.