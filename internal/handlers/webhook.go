package handlers

import (
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/pmoroney/ambient-weather-prometheus/internal/metrics"
	"github.com/pmoroney/ambient-weather-prometheus/internal/models"
)

func WebhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Printf("Invalid request method: %s", r.Method)
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse query parameters
	query := r.URL.Query()

	// Parse the DateUTC field into a Unix timestamp
	var timestamp float64
	dateUTC := query.Get("dateutc")
	if dateUTC != "" {
		parsedTime, err := time.Parse("2006-01-02 15:04:05", dateUTC)
		if err != nil {
			log.Printf("Invalid dateutc value: %v", err)
		} else {
			timestamp = float64(parsedTime.Unix())
		}
	} else {
		log.Println("Missing dateutc value")
	}

	// Populate the WeatherData model with raw data
	data := models.WeatherData{
		StationID:        query.Get("PASSKEY"),
		StationType:      query.Get("stationtype"),
		DateUTC:          timestamp,
		WindSpeed:        parseFloat(query.Get("windspeedmph"), "wind speed"),
		WindGust:         parseFloat(query.Get("windgustmph"), "wind gust"),
		MaxDailyGust:     parseFloat(query.Get("maxdailygust"), "max daily gust"),
		WindDirection:    parseFloat(query.Get("winddir"), "wind direction"),
		WindDirAvg10m:    parseFloat(query.Get("winddir_avg10m"), "wind dir avg 10m"),
		UVIndex:          parseFloat(query.Get("uv"), "UV index"),
		SolarRadiation:   parseFloat(query.Get("solarradiation"), "solar radiation"),
		HourlyRain:       parseFloat(query.Get("hourlyrainin"), "hourly rain"),
		EventRain:        parseFloat(query.Get("eventrainin"), "event rain"),
		DailyRain:        parseFloat(query.Get("dailyrainin"), "daily rain"),
		WeeklyRain:       parseFloat(query.Get("weeklyrainin"), "weekly rain"),
		MonthlyRain:      parseFloat(query.Get("monthlyrainin"), "monthly rain"),
		YearlyRain:       parseFloat(query.Get("yearlyrainin"), "yearly rain"),
		BatteryRainGauge: parseInt(query.Get("battrain"), "battery rain gauge"),
		BaromRel:         parseFloat(query.Get("baromrelin"), "barometric pressure relative"),
		BaromAbs:         parseFloat(query.Get("baromabsin"), "barometric pressure absolute"),
	}

	// Add outdoor and indoor sensors as extra sensors 0 and 1
	data.Sensors = append(data.Sensors, models.SensorData{
		Temperature: parseFloat(query.Get("tempf"), "temperature"),
		Humidity:    parseFloat(query.Get("humidity"), "humidity"),
		Battery:     parseInt(query.Get("battout"), "battery outdoor"),
	})
	data.Sensors = append(data.Sensors, models.SensorData{
		Temperature: parseFloat(query.Get("tempinf"), "indoor temperature"),
		Humidity:    parseFloat(query.Get("humidityin"), "indoor humidity"),
		Battery:     parseInt(query.Get("battin"), "battery indoor"),
	})

	// Parse additional sensors dynamically
	data.Sensors = append(data.Sensors, parseExtraSensors(query)...)

	// Update Prometheus metrics with raw data
	metrics.WindSpeed.WithLabelValues(data.StationID).Set(data.WindSpeed)
	metrics.WindGust.WithLabelValues(data.StationID).Set(data.WindGust)
	metrics.MaxDailyGust.WithLabelValues(data.StationID).Set(data.MaxDailyGust)
	metrics.WindDirection.WithLabelValues(data.StationID).Set(data.WindDirection)
	metrics.WindDirAvg10m.WithLabelValues(data.StationID).Set(data.WindDirAvg10m)
	metrics.UVIndex.WithLabelValues(data.StationID).Set(data.UVIndex)
	metrics.SolarRadiation.WithLabelValues(data.StationID).Set(data.SolarRadiation)
	metrics.HourlyRain.WithLabelValues(data.StationID).Set(data.HourlyRain)
	metrics.EventRain.WithLabelValues(data.StationID).Set(data.EventRain)
	metrics.DailyRain.WithLabelValues(data.StationID).Set(data.DailyRain)
	metrics.WeeklyRain.WithLabelValues(data.StationID).Set(data.WeeklyRain)
	metrics.MonthlyRain.WithLabelValues(data.StationID).Set(data.MonthlyRain)
	metrics.YearlyRain.WithLabelValues(data.StationID).Set(data.YearlyRain)
	metrics.BatteryRainGauge.WithLabelValues(data.StationID).Set(float64(data.BatteryRainGauge))
	metrics.BaromRel.WithLabelValues(data.StationID).Set(data.BaromRel)
	metrics.BaromAbs.WithLabelValues(data.StationID).Set(data.BaromAbs)

	// Set the timestamp metric
	if timestamp > 0 {
		metrics.Timestamp.WithLabelValues(data.StationID).Set(data.DateUTC)
	}

	// Update Prometheus metrics for extra sensors
	for i, sensor := range data.Sensors {
		sensorID := strconv.Itoa(i) // Sensor IDs start from 0 (outdoor sensor)
		metrics.SensorTemperature.WithLabelValues(data.StationID, sensorID).Set(sensor.Temperature)
		metrics.SensorHumidity.WithLabelValues(data.StationID, sensorID).Set(sensor.Humidity)
		metrics.SensorBattery.WithLabelValues(data.StationID, sensorID).Set(float64(sensor.Battery))
	}

	log.Printf("Processed raw data for station: %s", data.StationID)
	w.WriteHeader(http.StatusNoContent)
}

// Helper function to parse extra sensors dynamically
func parseExtraSensors(query url.Values) []models.SensorData {
	var sensors []models.SensorData
	for i := 2; ; i++ { // Start from sensor 2 (temp2f, humidity2, batt2)
		tempKey := "temp" + strconv.Itoa(i) + "f"
		humidityKey := "humidity" + strconv.Itoa(i)
		batteryKey := "batt" + strconv.Itoa(i)

		tempStr := query.Get(tempKey)
		humidityStr := query.Get(humidityKey)
		batteryStr := query.Get(batteryKey)

		// Break the loop if no more sensors are found
		if tempStr == "" || humidityStr == "" || batteryStr == "" {
			break
		}

		sensor := models.SensorData{
			Temperature: parseFloat(tempStr, tempKey),
			Humidity:    parseFloat(humidityStr, humidityKey),
			Battery:     parseInt(batteryStr, batteryKey),
		}
		sensors = append(sensors, sensor)
	}
	return sensors
}

// Helper function to parse float values
func parseFloat(value string, fieldName string) float64 {
	if value == "" {
		log.Printf("Missing %s value", fieldName)
		return 0
	}
	parsed, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.Printf("Invalid %s value: %v", fieldName, err)
		return 0
	}
	return parsed
}

// Helper function to parse integer values
func parseInt(value string, fieldName string) int {
	if value == "" {
		log.Printf("Missing %s value", fieldName)
		return 0
	}
	parsed, err := strconv.Atoi(value)
	if err != nil {
		log.Printf("Invalid %s value: %v", fieldName, err)
		return 0
	}
	return parsed
}
