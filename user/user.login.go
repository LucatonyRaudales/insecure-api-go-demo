package user
import (
	"log"
	"net/http"
	"time" // For rate limiting

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func LoginUserV2(c *gin.Context) {
    type Credentials struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    var creds Credentials
    if err := c.BindJSON(&creds); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
        return
    }

    // --- 1. Basic Rate Limiting / Brute-Force Protection ---
    // Check if the user is currently locked out
    if loginAttempts[creds.Email] >= maxLoginAttempts {
        if time.Since(lastAttemptTime[creds.Email]) < lockoutDuration {
            c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many login attempts. Please try again later."})
            return
        } else {
            // Reset attempts if lockout duration has passed
            loginAttempts[creds.Email] = 0
        }
    }

    // Find the user by email
    var foundUser *User
    for _, u := range Users {
        if u.Email == creds.Email {
            foundUser = &u
            break
        }
    }

    // Check if user exists and if password matches
    if foundUser == nil || bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(creds.Password)) != nil {
        // --- 2. Increment Login Attempts on Failure ---
        loginAttempts[creds.Email]++
        lastAttemptTime[creds.Email] = time.Now()

        // --- 3. Generic Error Message ---
        // Avoid telling attackers if the user exists or not.
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    // --- 4. Reset Login Attempts on Success ---
    delete(loginAttempts, creds.Email) // Clear attempts on successful login
    delete(lastAttemptTime, creds.Email)

    // --- 5. Return Token (assuming it's a securely generated JWT or session ID) ---
    c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": foundUser.Token})
}



/** Helper **/


var loginAttempts = make(map[string]int)
var lastAttemptTime = make(map[string]time.Time)
const maxLoginAttempts = 5
const lockoutDuration = 5 * time.Minute // Lock out for 5 minutes

func Init() {
    hashPassword := func(password string) (string, error) {
        hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
        if err != nil {
            return "", err
        }
        return string(hashedPassword), nil
    }

    hashedPassAlice, err := hashPassword("password123")
    if err != nil {
        log.Fatalf("Error hashing Alice's password: %v", err)
    }
    hashedPassBob, err := hashPassword("adminpass")
    if err != nil {
        log.Fatalf("Error hashing Bob's password: %v", err)
    }

    Users = []User{
        {ID: 1, Name: "Alice", Email: "alice@example.com", Password: hashedPassAlice, Token: "alice_jwt_token"},
        {ID: 2, Name: "Bob", Email: "bob@example.com", Password: hashedPassBob, Token: "bob_jwt_token"},
    }
    log.Println("Users initialized with hashed passwords.")
}
