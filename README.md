# SolarSchedule

## Overview
SolarSchedules is a RESTful API built with Go and the Gin framework, designed to provide sunrise and sunset times based on a given location. It's a simple yet effective tool for anyone needing accurate solar schedule information, from photographers planning the golden hour to astronomers tracking the daylight hours.

## Features
- Retrieve sunrise and sunset times by city, state, or ZIP code.
- Easy-to-use RESTful API with JSON responses.
- Fast and efficient, built with Go and Gin.

## Getting Started

### Prerequisites
- Go 1.15 or later.

### Installation
1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/SolarSchedules.git
2. Navigate to the project directory:
   ```sh
   cd SolarSchedules
   ```
3. Install the dependencies:
   ```sh
   go mod tidy
   ```

### Running the Application
Run the application using:
```sh
go run cmd/solarschedules/main.go
```
The server will start, typically listening on port 8080.

### Usage
Make GET requests to `/times` with the required query parameters. For example:
```
GET http://localhost:8080/times?city=NewYork&state=NY
```

## API Documentation
### Endpoints
#### `/times`
- **Method**: GET
- **Parameters**:
    - `city`: Name of the city (optional if ZIP code is provided).
    - `state`: Name of the state (optional).
    - `zip`: ZIP code (optional if city and state are provided).
- **Response**: JSON object containing sunrise and sunset times.

### Example Response
```json
{
  "sunrise": "06:24:00",
  "sunset": "18:46:00"
}
```

## Contributing
Contributions to SolarSchedule are welcome. Please follow these steps:
1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes and commit (`git commit -am 'Add some feature'`).
4. Push to the branch (`git push origin feature-branch`).
5. Create a new Pull Request.

## License
This project is licensed under the [MIT License](LICENSE).
