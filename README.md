# Dory - Real-Time Chat Application

Dory is a basic proof-of-concept (PoC) real-time chat application built with Go (using the Fiber framework) and WebSockets. It includes user authentication via JWT (JSON Web Tokens) and a simple frontend for sending and receiving messages in real time.

## Features
- **Real-Time Messaging**: Uses WebSockets for instant message delivery.
- **JWT Authentication**: Users can log in and receive a JWT for session management.
- **Basic Frontend**: A simple HTML/JavaScript interface for interacting with the chat.
- **Middleware Protection**: Routes like `/chat` are protected by JWT validation middleware.

## Disclaimer
This project is a basic proof of concept (PoC) for a real-time chat application using WebSockets and JWT authentication. It is intended for educational purposes and demonstrates foundational skills in full-stack development.

**This implementation is not production-ready and lacks proper security measures.** For example:
- The JWT cookie is not fully secured (`HTTPOnly` is disabled).
- The WebSocket endpoint is not protected by authentication middleware.
- Input validation and error handling are minimal.
- There is only one password for all users, which is obviously not acceptable by modern security standards, but it was easier for this PoC.

Use this code as a learning resource or a starting point for further development. Do not deploy it in a production environment without significant improvements to security and robustness.

---

## Prerequisites
- Go (version 1.20 or higher)
- Node.js (optional, for frontend development)
- A `.env` file with the following environment variables:
  ```plaintext
  JWT_SECRET=your_jwt_secret_key
  PASSWORD=your_base64_encoded_password
  ```

---

## Setup and Installation

1. **Clone the Repository**
   ```bash
   git clone https://github.com/booIeann/Dory.git
   cd Dory
   ```

2. **Set Up Environment Variables**
   Create a `.env` file in the root directory and add the required variables:
   ```plaintext
   JWT_SECRET=your_jwt_secret_key
   PASSWORD=your_base64_encoded_password
   ```
   To generate a base64-encoded password, you can use:
   ```bash
   echo -n "your_password" | base64
   ```

3. **Run the Application**
   Start the server:
   ```bash
   go run main.go
   ```
   The application will be available at `http://localhost:8080`.

4. **Access the Chat**
   - Open your browser and navigate to `http://localhost:8080`.
   - Log in using the username and password (the password is the decoded value of the `PASSWORD` environment variable).
   - After logging in, you will be redirected to the chat page.

---

## Project Structure

```
dory/
├── public/                  # Frontend files (HTML, CSS, JS)
│   ├── index.html           # Login page
│   └── chat.html            # Chat page
├── internal/
│   ├── api/                 # Backend handlers and middleware
│   │   ├── handlers.go      # Route handlers (login, logout, WebSocket)
│   │   ├── middleware.go    # JWT validation middleware
│   │   ├── responses.go     # Utility functions for API responses
│   │   └── routes.go        # API route definitions
│   └── auth/                # Authentication logic
│       └── auth.go          # JWT token generation and validation
├── cmd/
│   └── main.go              # Entry point for the application
├── README.md                # Project documentation
└── .env.example             # Example environment variables file
```

---

## Usage

1. **Login**
   - Visit `http://localhost:8080`.
   - Enter a username and the correct password (the decoded value of the `PASSWORD` environment variable).
   - Upon successful login, you will be redirected to the chat page.

2. **Chat**
   - On the chat page, you can send and receive messages in real time.
   - Messages are broadcast to all connected users.

3. **Logout**
   - Click the "Logout" button to invalidate your session and return to the login page.

---

## Technologies Used
- **Backend**: Go (Fiber framework)
- **Frontend**: HTML, CSS, JavaScript
- **Real-Time Communication**: WebSockets
- **Authentication**: JWT (JSON Web Tokens)

---

## Contributing
Contributions are welcome! If you find any issues or have suggestions for improvement, please open an issue or submit a pull request.

---

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## Acknowledgments
- [Fiber](https://gofiber.io/) for the web framework.
- [WebSocket](https://developer.mozilla.org/en-US/docs/Web/API/WebSocket) for real-time communication.
- [JWT](https://jwt.io/) for authentication.