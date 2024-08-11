# basic-crypto-dashboard
A cryptocurrency dashboard that displays real-time data for the top 10 cryptocurrencies, built with React, GraphQL, and Go.

# Crypto Dashboard

A cryptocurrency dashboard designed to provide real-time data for the top 10 cryptocurrencies. This project integrates with the CoinGecko API to fetch up-to-date cryptocurrency information and is built using Go and GraphQL for a robust backend solution.

## Overview

- **Real-Time Data:** Fetches and displays live data for the top 10 cryptocurrencies using the CoinGecko API.
- **Interactive UI:** Features a modern and engaging user interface for tracking cryptocurrency trends.
- **Backend Technology:** Utilizes Go for a high-performance backend and GraphQL for efficient data querying and management.

## API Integration

The dashboard retrieves cryptocurrency data from the [CoinGecko API](https://coingecko.com), which provides comprehensive and real-time information about various cryptocurrencies. This ensures that users have access to the latest market data.

```graphql
{
  coins {
    id
    name
    symbol
    price
    oneHourChange
    twentyFourHourChange
    sevenDayChange
    marketCap
    volume24h
  }
}
```
This query fetches the following details for each cryptocurrency:

- **ID:** Unique identifier for the cryptocurrency.
- **Name:** Full name of the cryptocurrency.
- **Symbol:** Symbol or ticker of the cryptocurrency.
- **Price:** Current price of the cryptocurrency.
- **One-Hour Change:** Percentage change in price over the past hour.
- **24-Hour Change:** Percentage change in price over the past 24 hours.
- **Seven-Day Change:** Percentage change in price over the past seven days.
- **Market Cap:** Total market capitalization of the cryptocurrency.
- **24-Hour Volume:** Total trading volume over the past 24 hours.

## Technologies Used

- **Go:** A powerful programming language used for building backend services, ensuring scalability and performance.
- **GraphQL:** A query language for APIs that allows for efficient and flexible data retrieval from the backend.
- **React:** A JavaScript library for building the user interface, providing a dynamic and responsive experience.

## Purpose

The goal of this project is to offer an up-to-date view of the cryptocurrency market, helping users stay informed about the latest trends and data for major cryptocurrencies.

## Getting Started

To get started, follow the instructions provided in the [Installation](#installation) section of the repository.

## Contributing

We welcome contributions to enhance the dashboard. Please refer to the [Contributing](#contributing) section for guidelines on how to get involved.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.
