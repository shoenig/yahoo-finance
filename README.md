# yahoo-finance

A Go library for interacting with Yahoo Finance.



[nuc10 yahoo-finance] $ curl -s "https://query1.finance.yahoo.com/v8/finance/chart/amd" | jq .
{
  "chart": {
    "result": [
      {
        "meta": {
          "currency": "USD",
          "symbol": "AMD",
          "exchangeName": "NMS",
          "instrumentType": "EQUITY",
          "firstTradeDate": 322151400,
          "regularMarketTime": 1619208002,
          "gmtoffset": -14400,
          "timezone": "EDT",
          "exchangeTimezoneName": "America/New_York",
          "regularMarketPrice": 82.76,
          "chartPreviousClose": 79.06,
          "previousClose": 79.06,
          "scale": 3,
          "priceHint": 2,
          "currentTradingPeriod": {
            "pre": {
              "timezone": "EDT",
              "start": 1619164800,
              "end": 1619184600,
              "gmtoffset": -14400
            },
            "regular": {
              "timezone": "EDT",
              "start": 1619184600,
              "end": 1619208000,
              "gmtoffset": -14400
            },
            "post": {
              "timezone": "EDT",
              "start": 1619208000,
              "end": 1619222400,
              "gmtoffset": -14400
            }
          },
          "tradingPeriods": [
            [
              {
                "timezone": "EDT",
                "start": 1619184600,
                "end": 1619208000,
                "gmtoffset": -14400
              }
            ]
          ],
