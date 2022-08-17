package weather

type Response struct {
	Coordinate       coordinate       `json:"coord"`
	Weather          []weather        `json:"weather"`
	Base             string           `json:"base"`
	Main             main             `json:"main"`
	Visibility       int              `json:"visibility"`
	Wind             wind             `json:"wind"`
	LocationMetadata locationMetadata `json:"sys"`
	Timezone         int              `json:"timezone"`
	Id               int              `json:"id"`
	Name             string           `json:"name"`
	Cod              int              `json:"cod"`
}

type coordinate struct {
	Longitude float64 `json:"lon"`
	Latitude  float64 `json:"lat"`
}

type weather struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
}

type wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
}

type locationMetadata struct {
	Type    int     `json:"type"`
	Id      int     `json:"id"`
	Message float64 `json:"message"`
	Country string  `json:"country"`
	Sunrise int     `json:"sunrise"`
	Sunset  int     `json:"sunset"`
}
