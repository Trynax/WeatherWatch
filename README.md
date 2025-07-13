# WeatherWatch 🌦️

A real-time weather monitoring and alerting system built in Go that tracks weather conditions and sends notifications when specified thresholds are exceeded.

## Features ✨

- **Real-time Weather Monitoring**: Fetches current weather data using WeatherAPI
- **Customizable Alert System**: Set thresholds for temperature, humidity, wind speed, and more
- **Multiple Notification Channels**:
  - Console output with colored formatting
  - Email notifications (Gmail SMTP)
  - File logging with detailed reports
- **Scheduled Monitoring**: Configurable monitoring intervals
- **Alert Conditions**:
  - High/Low temperature alerts
  - High humidity alerts
  - High wind speed alerts
  - Extreme weather condition alerts

## Project Structure 📁

```
WeatherWatch/
├── cmd/
│   └── weather-alert/
│       └── main.go              # Application entry point
├── alert/
│   └── alert.go                 # Alert condition logic
├── api/
│   └── weatherapi.go            # WeatherAPI integration
├── notify/
│   ├── console.go               # Console notification system
│   ├── file.go                  # File logging system
│   └── sendEmail.go             # Email notification system
├── scheduler/
│   └── scheduler.go             # Weather monitoring scheduler
├── utils/
│   └── format.go                # Message formatting utilities
├── .env                         # Environment configuration
├── weatherReport.txt            # Generated weather logs
├── go.mod                       # Go module file
├── go.sum                       # Go dependencies
├── Makefile                     # Build automation
└── README.md                    # This file
```

## Prerequisites 📋

- Go 1.19 or higher
- WeatherAPI account (free tier available)
- Gmail account with App Password (for email notifications)

## Installation & Setup 🚀

### 1. Clone the Repository
```bash
git clone https://github.com/yourusername/WeatherWatch.git
cd WeatherWatch
```

### 2. Install Dependencies
```bash
go mod tidy
```

### 3. Environment Configuration

Create a `.env` file in the root directory:

```env
# WeatherAPI Configuration
API_KEY=your_weatherapi_key_here

# Email Configuration (Gmail SMTP)
SENDER_EMAIL="your-email@gmail.com"
SENDER_PASSWORD="your-app-password"
FROM_EMAIL_SMTP="smtp.gmail.com"
SMTP_ADDR="smtp.gmail.com:587"
RECIPIENT_EMAIL="recipient@gmail.com"
```

### 4. Get WeatherAPI Key
1. Visit [WeatherAPI.com](https://www.weatherapi.com/)
2. Sign up for a free account
3. Copy your API key to the `.env` file

### 5. Setup Gmail App Password
1. Enable 2-Factor Authentication on your Gmail account
2. Go to Google Account Settings → Security → App Passwords
3. Generate a new app password
4. Use this password in your `.env` file (not your regular Gmail password)

## Usage 💻

### Basic Usage

Run the weather monitoring system:
```bash
make run
```

Or directly with Go:
```bash
go run ./cmd/weather-alert/main.go
```

### Customization

#### Change Monitoring Location
Edit `cmd/weather-alert/main.go`:
```go
scheduler.CheckWeatherOnSchedule(1, "Lagos")  // Change "Lagos" to your city
```

#### Adjust Monitoring Interval
```go
scheduler.CheckWeatherOnSchedule(5, "Lagos")  // Check every 5 minutes
```

#### Customize Alert Thresholds
Edit `alert/alert.go` to modify temperature, humidity, and wind speed thresholds:

```go
// Example thresholds
const (
    HIGH_TEMP_THRESHOLD = 35.0    // Celsius
    LOW_TEMP_THRESHOLD  = 5.0     // Celsius
    HIGH_HUMIDITY_THRESHOLD = 85  // Percentage
    HIGH_WIND_THRESHOLD = 25.0    // km/h
)
```

## Alert Conditions 🚨

The system monitors and alerts on:

- **High Temperature**: Above configured threshold
- **Low Temperature**: Below configured threshold  
- **High Humidity**: Above configured percentage
- **High Wind Speed**: Above configured speed
- **Extreme Weather**: Severe weather conditions

## Output Examples 📊

### Console Output
```
📍 Lagos, Nigeria
🕓 Local Time: 2025-07-13 09:10
🌡️ Temperature: 25.0°C (Feels like: 27.0°C)
☁️ Condition: Partly cloudy
💧 Humidity: 89%
🌬️ Wind: 20.2 kph (SW)
☁️ Cloud Cover: 50%
```

### Alert Example
```
⚠️ ALERTS Triggered
📍 Lagos, Nigeria | 🕒 2025-07-13 00:00
---------------------------------
🌬️ High Wind Speed: 26.6 kph (threshold: 22.3 kph)
```

## File Structure Details 📋

### Core Components

- **`cmd/weather-alert/main.go`**: Application entry point
- **`scheduler/scheduler.go`**: Main monitoring loop with signal handling
- **`api/weatherapi.go`**: WeatherAPI client and data structures
- **`alert/alert.go`**: Alert condition checking logic
- **`notify/`**: Notification systems (console, email, file)
- **`utils/format.go`**: Message formatting utilities

### Generated Files

- **`weatherReport.txt`**: Continuous log of weather data and alerts
- **Email notifications**: Sent to configured recipient

## Development 🛠️

### Build the Project
```bash
go build ./cmd/weather-alert
```

### Run Tests
```bash
go test ./...
```

### Add New Alert Conditions

1. Modify `alert/alert.go` to add new condition logic
2. Update `utils/format.go` for new message formats
3. Test with your local weather conditions

## Troubleshooting 🔧

### Email Not Sending
- Verify Gmail App Password is correct
- Check if 2FA is enabled on Gmail account
- Ensure firewall allows outbound SMTP connections
- Try different SMTP ports (587 or 465)

### API Errors
- Verify WeatherAPI key is valid and not expired
- Check city name spelling
- Ensure internet connectivity

### High Alert Frequency
- Adjust alert thresholds in `alert/alert.go`
- Increase monitoring interval
- Check if weather conditions are genuinely extreme

## Contributing 🤝

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/new-feature`)
3. Commit changes (`git commit -am 'Add new feature'`)
4. Push to branch (`git push origin feature/new-feature`)
5. Create a Pull Request

## License 📝

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments 🙏

- [WeatherAPI](https://www.weatherapi.com/) for weather data
- [godotenv](https://github.com/joho/godotenv) for environment configuration
- Go community for excellent libraries and documentation

## Support 💬

If you encounter any issues or have questions:

1. Check the troubleshooting section above
2. Review your `.env` configuration
3. Open an issue on GitHub with detailed error messages

---

**Happy Weather Monitoring! 🌦️**
