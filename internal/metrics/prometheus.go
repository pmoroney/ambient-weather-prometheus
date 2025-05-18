package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	WindSpeed = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "weather_wind_speed_mph",
			Help: "Wind speed in MPH",
		},
		[]string{"station_id"},
	)

	WindGust = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "weather_wind_gust_mph",
			Help: "Wind gust in MPH",
		},
		[]string{"station_id"},
	)

	MaxDailyGust = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "weather_max_daily_gust_mph",
			Help: "Maximum daily wind gust in MPH",
		},
		[]string{"station_id"},
	)

	WindDirection = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "weather_wind_direction_degrees",
			Help: "Wind direction in degrees",
		},
		[]string{"station_id"},
	)

	UVIndex = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "weather_uv_index",
			Help: "UV index",
		},
		[]string{"station_id"},
	)

	SolarRadiation = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "weather_solar_radiation",
			Help: "Solar radiation in W/m²",
		},
		[]string{"station_id"},
	)

	HourlyRain = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "weather_hourly_rain_inches",
			Help: "Hourly rain in inches",
		},
		[]string{"station_id"},
	)

	DailyRain = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "weather_daily_rain_inches",
			Help: "Daily rain in inches",
		},
		[]string{"station_id"},
	)

	BaromRel = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "weather_barometric_pressure_relative_inhg",
			Help: "Relative barometric pressure in inHg",
		},
		[]string{"station_id"},
	)

	BaromAbs = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "weather_barometric_pressure_absolute_inhg",
			Help: "Absolute barometric pressure in inHg",
		},
		[]string{"station_id"},
	)

	SensorTemperature = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "weather_sensor_temperature_fahrenheit",
			Help: "Temperature of external sensors in Fahrenheit",
		},
		[]string{"station_id", "sensor_id"},
	)

	SensorHumidity = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "weather_sensor_humidity_percentage",
			Help: "Humidity of extra sensors in percentage",
		},
		[]string{"station_id", "sensor_id"},
	)

	SensorBattery = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "weather_sensor_battery_status",
			Help: "Battery status of extra sensors",
		},
		[]string{"station_id", "sensor_id"},
	)

	WindDirAvg10m = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "weather_wind_direction_avg_10m_degrees",
			Help: "Average wind direction over 10 minutes in degrees",
		},
		[]string{"station_id"},
	)

	EventRain = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "weather_event_rain_inches",
			Help: "Event rain in inches",
		},
		[]string{"station_id"},
	)

	WeeklyRain = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "weather_weekly_rain_inches",
			Help: "Weekly rain in inches",
		},
		[]string{"station_id"},
	)

	MonthlyRain = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "weather_monthly_rain_inches",
			Help: "Monthly rain in inches",
		},
		[]string{"station_id"},
	)

	YearlyRain = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "weather_yearly_rain_inches",
			Help: "Yearly rain in inches",
		},
		[]string{"station_id"},
	)

	BatteryRainGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "weather_battery_rain_gauge_status",
			Help: "Rain gauge battery status",
		},
		[]string{"station_id"},
	)

	Timestamp = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "weather_data_timestamp",
			Help: "Timestamp of the weather data in Unix time",
		},
		[]string{"station_id"},
	)
)

func RegisterMetrics() {
	prometheus.MustRegister(WindSpeed)
	prometheus.MustRegister(WindGust)
	prometheus.MustRegister(MaxDailyGust)
	prometheus.MustRegister(WindDirection) // Register the new WindDirection metric
	prometheus.MustRegister(UVIndex)
	prometheus.MustRegister(SolarRadiation)
	prometheus.MustRegister(HourlyRain)
	prometheus.MustRegister(DailyRain)
	prometheus.MustRegister(BaromRel)
	prometheus.MustRegister(BaromAbs)
	prometheus.MustRegister(SensorTemperature)
	prometheus.MustRegister(SensorHumidity)
	prometheus.MustRegister(SensorBattery)
	prometheus.MustRegister(WindDirAvg10m)
	prometheus.MustRegister(EventRain)
	prometheus.MustRegister(WeeklyRain)
	prometheus.MustRegister(MonthlyRain)
	prometheus.MustRegister(YearlyRain)
	prometheus.MustRegister(BatteryRainGauge)
	prometheus.MustRegister(Timestamp)
}

func ExposeMetrics(w http.ResponseWriter, r *http.Request) {
	promhttp.Handler().ServeHTTP(w, r)
}
