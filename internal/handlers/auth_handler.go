package handlers

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/satyam1560/todo_backend/internal/auth"
	db "github.com/satyam1560/todo_backend/internal/database/generated"
)

type AuthHandler struct {
	Q *db.Queries
}

func (h *AuthHandler) LoginWithFirebase(c *gin.Context) {
	var req struct {
		IDToken string `json:"id_token"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Initialize Firebase
	if err := auth.InitFirebase(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to initialize Firebase"})
		return
	}

	client, err := auth.FirebaseApp.Auth(context.Background())
	if err != nil {
		log.Printf("🔥 Failed to get Firebase Auth client: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get Firebase Auth client"})
		return
	}

	userRecord, err := client.VerifyIDToken(context.Background(), req.IDToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Firebase token"})
		return
	}

	uid := userRecord.UID
	phone := userRecord.Firebase.Identities["phone"].([]interface{})[0].(string)

	log.Println("User ID:", uid)
	log.Println("Phone number:", phone)

	// Get or create user
	ctx := context.Background()
	user, err := h.Q.GetUserByPhone(ctx, phone)
	if err != nil {
		if err == sql.ErrNoRows {
			// Create new user
			userID := uuid.New()
			user, err = h.Q.CreateUser(ctx, db.CreateUserParams{
				ID:    userID,
				Phone: phone,
			})
			if err != nil {
				log.Printf("❌ Failed to create user: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
				return
			}
			log.Println("✅ New user created:", user.ID)
		} else {
			log.Printf("❌ DB error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query user"})
			return
		}
	} else {
		log.Println("🔁 User already exists:", user.ID)
	}

	// Return the user (or later you can attach a JWT here)
	token, err := auth.GenerateJWT(user.ID.String())
if err != nil {
	log.Printf("Failed to create JWT: %v", err)
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
	return
}

c.JSON(http.StatusOK, gin.H{"token": token})
}

