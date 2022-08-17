# Weather Forecast Api

This project was aimed for me to practice some new concepts of Golang.

## Setup

Steps needed:

* Duplicate the file `.env.example` and rename it to `.env`
* Create an account on https://openweathermap.org/api
* Copy API (AppId) provided by OpenWeatherMap on https://home.openweathermap.org/api_keys
* Paste the API (AppId) on the `OPEN_WEATHER_API_KEY` key inside the `.env` file.
* Open a terminal on this current directory.
* Execute on a terminal `go run main.go`.

## Usage

This `API` is versioned, therefore throughout this document `:api-version` refers to one of the publicly available versions:

* v1

This `API` exposes the following `endpoints`:

* [`/api/:api-version/weather`](#weather)

## Documentation

### `/weather`

This endpoint retrieves the current weather for the given city.

```http request
GET /api/:api-version/weather/:city
Content-Type: application/json
```

**V1 example response body:**

```json
{
    "coord": {
        "lon": -69.9886,
        "lat": 18.5001
    },
    "weather": [
        {
            "id": 802,
            "main": "Clouds",
            "description": "scattered clouds",
            "icon": "03n"
        }
    ],
    "base": "stations",
    "main": {
        "temp": 28.72,
        "feels_like": 33.97,
        "temp_min": 25.82,
        "temp_max": 28.91,
        "pressure": 1014,
        "humidity": 80
    },
    "visibility": 10000,
    "wind": {
        "speed": 0.45,
        "deg": 84
    },
    "sys": {
        "type": 2,
        "id": 2020859,
        "message": 0,
        "country": "DO",
        "sunrise": 1660731738,
        "sunset": 1660777562
    },
    "timezone": -14400,
    "id": 3492908,
    "name": "Santo Domingo",
    "cod": 200
}
```