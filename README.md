# ASCII Art Generator

## Overview
The ASCII Art Generator is a web-based application that allows users to create beautiful ASCII art from input text using different banner styles. Users can view the generated ASCII art directly on the website and also download it as a `.txt` file.

---

## Features
- **Interactive User Interface:** Input text and choose from various banner styles to generate ASCII art.
- **Real-time ASCII Art Display:** View the generated ASCII art instantly in your browser.
- **Download as File:** Save the ASCII art as a `.txt` file for offline use.
- **Custom Error Handling:** Provides descriptive error messages for invalid inputs or server issues.
- **Responsive Design:** Optimized for various screen sizes and devices.

---

## Installation

### Prerequisites
- Go (Golang) installed on your system.
- Access to a terminal or command prompt.

### Steps
1. Clone this repository:
   ```bash
   git clone https://github.com/MrMaRc0s/ascii-art-web-export/
   cd ascii-art-web-export
   ```

2. Run the server:
   ```bash
   go run main.go
   ```

3. Open your browser and navigate to:
   ```
   http://localhost:<PORT>
   ```

---

## Usage

### Web Interface
1. Enter your desired text in the text input box.
2. Select a banner style (Standard, Shadow, or Thinkertoy).
3. Click the **Generate ASCII Art** button to view the result.
4. To download the result, click **Download ASCII Art as txt**.

### API Endpoints

#### Generate ASCII Art
- **Endpoint:** `/ascii-art`
- **Method:** `POST`
- **Parameters:**
  - `text` (string): Input text for ASCII art.
  - `banner` (string): Banner style (`standard`, `shadow`, or `thinkertoy`).
- **Response:** ASCII art as plain text.

#### Download ASCII Art
- **Endpoint:** `/download-ascii-art`
- **Method:** `POST`
- **Parameters:**
  - `text` (string): Input text for ASCII art.
  - `banner` (string): Banner style.
- **Response:** ASCII art served as a `.txt` file.

Example command using `curl`:
```bash
curl -X POST -d "text=Hello&banner=standard" -o output.txt http://localhost:8080/download-ascii-art
```

---

## Error Handling
- **Empty Input Text:** Returns a `400 Bad Request` error.
- **Invalid Banner:** Returns a `400 Bad Request` error if the selected banner is not available.
- **Missing Banner Files:** Returns a `404 Not Found` error if a banner file is missing.
- **Internal Server Error:** Returns a `500 Internal Server Error` for unexpected errors during processing.

---

## Customization
- **Styling:** Modify `index.html` or the CSS styles to change the look and feel of the application.

---

## Acknowledgments
- Inspired by the creativity of ASCII art enthusiasts.
- Special thanks to users for testing and feedback.
