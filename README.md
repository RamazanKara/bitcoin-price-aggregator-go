# Bitcoin Price Aggregator & Next Day Trend Prognosis

## Introduction

This Go application provides historical price data for Bitcoin over the last 7 days and predicts the price for the next day, along with identifying the trend (upwards, downwards, or stable). The application fetches data from the CoinGecko API and uses a linear regression model for the prognosis.
It includes preprocessing steps to prepare the data for more accurate predictions.

## Installation

To set up this project, you need to have Go installed on your system.

1. Clone the repository:

   ```markdown
   git clone [repository-url]
   ```

2. Navigate to the project directory:

   ```markdown
   cd bitcoin-prognosis
   ```

## Usage

Run the application with the following command:

```bash
go run cmd/main/main.go
```

This will output the Bitcoin prices for the last 7 days and the prognosis for the next day, including the price trend.
You are free to also include any of the pkgs into your own application.

## Data Preprocessing

The application includes a preprocessing module that prepares the historical Bitcoin price data for the prognosis. This preprocessing includes handling missing values and removing outliers, providing a cleaner dataset for the prediction model.

## Testing

To run the tests for this application, use the following command in the project directory:

```bash
go test -v ./...
```

The tests cover the functionality of fetching Bitcoin prices and the prognosis algorithm.

## Contributing

Contributions to this project are welcome. Please fork the repository and submit a pull request with your changes.

## License

This project is open source and available under the [MIT License](LICENSE.md).
