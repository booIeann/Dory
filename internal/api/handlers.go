package api

import (
	"dory/internal/auth"
	"encoding/base64"
	"log"
	"os"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

var (
	clients   = make(map[*websocket.Conn]bool)
	clientsMu sync.Mutex
)

func LoginHandler(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	correctPassword := os.Getenv("PASSWORD")
	jwtDecode, err := base64.StdEncoding.DecodeString(correctPassword)
	if err != nil {
		return ErrorResponse(c, fiber.StatusInternalServerError, "Failed to decode password", nil)
	}

	if password != string(jwtDecode) {
		return ErrorResponse(c, fiber.StatusUnauthorized, "Invalid password", nil)
	}

	token, err := auth.GenerateToken(username)
	if err != nil {
		return ErrorResponse(c, fiber.StatusInternalServerError, "Failed to generate token", nil)
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "jwt"
	cookie.Value = token
	cookie.HTTPOnly = false
	cookie.Expires = time.Now().Add(6 * time.Hour)
	cookie.Path = "/"
	c.Cookie(cookie)

	return c.Redirect("/chat", fiber.StatusFound)
}

func FindingNemo(c *fiber.Ctx) error {
	return c.SendFile("/public/chat.html")
}

func LogoutHandler(c *fiber.Ctx) error {
	cookie := new(fiber.Cookie)
	cookie.Name = "jwt"
	cookie.Value = ""
	cookie.Expires = time.Now().Add(-1 * time.Hour)
	cookie.Path = "/"
	c.Cookie(cookie)

	return c.Redirect("/", fiber.StatusFound)
}

func HandleWebSocket(c *websocket.Conn) {
	clientsMu.Lock()
	clients[c] = true
	clientsMu.Unlock()

	defer func() {
		clientsMu.Lock()
		delete(clients, c)
		clientsMu.Unlock()
		c.Close()
	}()

	for {
		messageType, message, err := c.ReadMessage()
		if err != nil {
			log.Println("WebSocket read error:", err)
			break
		}

		clientsMu.Lock()
		for client := range clients {
			if err := client.WriteMessage(messageType, message); err != nil {
				log.Println("WebSocket write error:", err)
				client.Close()
				delete(clients, client)
			}
		}
		clientsMu.Unlock()
	}
}
