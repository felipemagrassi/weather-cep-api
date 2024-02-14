# Weather Cep API

This API is a simple weather API that returns the temperature of a city by receiving a CEP (Brazilian Zip Code) as a parameter.


## Usage

1. Add your `weatherapi.com` API key to the `.env` file 
```env
WEATHER_API_KEY=your_api_key
```

2. Add the CEP as a parameter in the URL and the API will return the temperature of the city.
```http
curl http://localhost:8080/?cep=10010-000
```

## Demo

* With valid cep
https://weather-cep-api-zj3c47ztra-uc.a.run.app/?cep=20561-250

* With invalid cep
https://weather-cep-api-zj3c47ztra-uc.a.run.app/?cep=

* With not found cep 
https://weather-cep-api-zj3c47ztra-uc.a.run.app/?cep=11111-111
