# Currency Exchange Service
This service provides an API for converting amounts between different currencies using pre-defined exchange rates.

## Prerequisites
- Docker
- docker-compose
- A `.env` file placed in the project root directory

## Configuration
The service configuration is managed via a `.env` file located in the project root. It allows for customization of the server port and the path to the currency rates JSON file.

Example `.env` file content:
```
SERVER_PORT=8080
RATE_EXCHANGE_JSON_PATH="./currencies.json"
```
- If the `SERVER_PORT` variable is not specified, the service defaults to using port 8080.
- If the `RATE_EXCHANGE_JSON_PATH` variable is not specified or the path is not found, the service will use default currencies exchange rate that only includes USD, TWD and JPY.

### Customizing Currency Rates
You can modify the `currencies.json` file in the project root to specify your own currency exchange rates. 

Alternatively, you can change the path to the JSON file in the `.env` file using the `RATE_EXCHANGE_JSON_PATH` variable to point to a different file.

## Building and Running the Service
To build and run the service, use the following command:
```sh
make run
```
For building the service without running, use the make build command. This will build the Docker image for the currency exchange service.

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
```json
{
  "msg": "success",
  "amount": "170,496.53"
}
```

## License
This project is licensed under the MIT License - see the LICENSE.md file for details.