# GLUCO-M - Dietary Recommendation System for Diabetes Management

## Overview
GLUCO-M is a dietary recommendation system designed to help diabetes patients manage their food intake efficiently. Developed using **Golang**, **HTML**, and **CSS**, this system provides meal recommendations based on user health data and glucose levels.

## Features
- **User Data Collection** – Gathers height, weight, and dietary habits for personalized suggestions.
- **Glucose Monitoring** – Accepts glucose level inputs from biosensors or manual entry.
- **Nutrition Database** – Uses a pre-defined food database to recommend suitable meal plans.
- **Fast & Efficient** – Processes meal recommendations quickly based on user input.
- **Scalability** – Designed for future integration with biosensors, AI meal predictions, and insulin dosage recommendations.

## How It Works
1. Users input their health details and glucose readings.
2. The system processes the data and fetches relevant dietary recommendations.
3. Personalized meal plans are displayed based on calorie intake and sugar levels.
4. Alerts and suggestions are provided for critical glucose levels.

## Installation
1. **Install Golang** (if not already installed):  
   Download and install Go from [https://golang.org/dl/](https://golang.org/dl/)

2. **Clone the Repository:**  
   ```sh
   git clone https://github.com/yourusername/GLUCO-M.git
   cd GLUCO-M
   ```
3. **Unzip static.tar file**
   Ensure that it is found in the same location as the main.go file. This contains the databases to which you can also add new food items.

3. **Run the Application:**  
   ```sh
   go run main.go
   ```

4. **Access the Web Interface:**  
   Open your browser and go to `http://localhost:8080`

## Future Enhancements
- **API Integration** – Connect with external nutrition APIs for broader food recommendations.
- **Mobile App Support** – Develop a mobile version for real-time monitoring.
- **AI-based Meal Predictions** – Enhance dietary recommendations using machine learning.
- **Insulin Dosage Suggestions** – Provide insulin recommendations based on user data.

## License
This project is open-source under the MIT License.

## Contributors
Benjamin Isaac Thomson
Rameswar Chandran
Karthik S

For contributions or feature requests, feel free to open an issue or submit a pull request!
