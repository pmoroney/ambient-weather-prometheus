package models

type SensorData struct {
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
	Battery     int     `json:"battery"`
}

type WeatherData struct {
	StationID        string       `json:"station_id"`
	StationType      string       `json:"station_type"`
	DateUTC          float64      `json:"date_utc"`
	WindSpeed        float64      `json:"wind_speed"`
	WindGust         float64      `json:"wind_gust"`
	MaxDailyGust     float64      `json:"max_daily_gust"`
	WindDirection    float64      `json:"wind_direction"`
	WindDirAvg10m    float64      `json:"wind_dir_avg_10m"`
	UVIndex          float64      `json:"uv_index"`
	SolarRadiation   float64      `json:"solar_radiation"`
	HourlyRain       float64      `json:"hourly_rain"`
	EventRain        float64      `json:"event_rain"`
	DailyRain        float64      `json:"daily_rain"`
	WeeklyRain       float64      `json:"weekly_rain"`
	MonthlyRain      float64      `json:"monthly_rain"`
	YearlyRain       float64      `json:"yearly_rain"`
	BatteryRainGauge int          `json:"battery_rain_gauge"`
	BaromRel         float64      `json:"barom_rel"`
	BaromAbs         float64      `json:"barom_abs"`
	Sensors          []SensorData `json:"sensors"`
}
