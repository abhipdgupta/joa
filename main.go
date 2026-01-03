package main

import (
	"fmt"

	joa "github.com/abhipdgupta/joa/cmd"
)

func main() {
	lex := joa.NewLexer("{\r\n    \"id\": \"usr_102934\ndasdass\",\r\n    \"username\": \"abhi_dev\",\r\n    \"email\": \"abhi@example.com\",\r\n    \"isActive\": true,\r\n    \"age\": 27,\r\n    \"rating\": 4.8,\r\n    \"createdAt\": \"2025-01-12T10:45:30Z\",\r\n    \"lastLogin\": null,\r\n    \"roles\": [\r\n        \"USER\",\r\n        \"ADMIN\"\r\n    ],\r\n    \"profile\": {\r\n        \"firstName\": \"Abhishek\",\r\n        \"lastName\": \"Gupta\",\r\n        \"gender\": \"male\",\r\n        \"dateOfBirth\": \"1998-04-22\",\r\n        \"bio\": \"Full stack developer who likes clean architecture.\",\r\n        \"avatarUrl\": \"https://cdn.example.com/avatars/abhi.png\"\r\n    },\r\n    \"address\": {\r\n        \"street\": \"221B Baker Street\",\r\n        \"city\": \"Bangalore\",\r\n        \"state\": \"Karnataka\",\r\n        \"country\": \"India\",\r\n        \"postalCode\": \"560001\",\r\n        \"geo\": {\r\n            \"latitude\": 12.9716,\r\n            \"longitude\": 77.5946\r\n        }\r\n    },\r\n    \"preferences\": {\r\n        \"language\": \"en\",\r\n        \"timezone\": \"Asia/Kolkata\",\r\n        \"theme\": \"dark\",\r\n        \"notifications\": {\r\n            \"email\": true,\r\n            \"sms\": false,\r\n            \"push\": true\r\n        }\r\n    },\r\n    \"security\": {\r\n        \"emailVerified\": true,\r\n        \"twoFactorEnabled\": false,\r\n        \"loginAttempts\": 1,\r\n        \"lastPasswordChange\": \"2024-11-01T08:12:00Z\"\r\n    },\r\n    \"sessions\": [\r\n        {\r\n            \"sessionId\": \"sess_abc123\",\r\n            \"ipAddress\": \"192.168.1.10\",\r\n            \"device\": \"Chrome on Windows\",\r\n            \"createdAt\": \"2025-01-20T09:00:00Z\",\r\n            \"expiresAt\": \"2025-01-20T21:00:00Z\"\r\n        },\r\n        {\r\n            \"sessionId\": \"sess_xyz789\",\r\n            \"ipAddress\": \"192.168.1.15\",\r\n            \"device\": \"Mobile App\",\r\n            \"createdAt\": \"2025-01-18T07:30:00Z\",\r\n            \"expiresAt\": \"2025-01-18T19:30:00Z\"\r\n        }\r\n    ],\r\n    \"number_in_json\":{\r\n        \"integer\": 42,\r\n        \"negativeInteger\": -7,\r\n        \"float\": 3.14159,\r\n        \"negativeFloat\": -0.00123,\r\n        \"scientificPositive\": 1.2e10,\r\n        \"scientificNegative\": -4.5E-8\r\n    },\r\n    \"metadata\": {\r\n        \"source\": \"web\",\r\n        \"referralCode\": null,\r\n        \"tags\": [\r\n            \"beta-user\",\r\n            \"power-user\"\r\n        ]\r\n    }\r\n}")

	lex.Lex()

	for token := range lex.Tokens() {
		fmt.Println(token)
	}
}
