# Hayavo Meet Go SDK

Go SDK for integrating **Hayavo Meet** real-time video, audio, and screen-sharing capabilities into your Go application.

## Installation

```bash
go get github.com/hayavo/hayavo-meet-go
```

## Initialize Client

```go
package main

import (
	"context"
	"log"
	"time"

	meet "github.com/hayavo/hayavo-meet-go"
	"github.com/hayavo/hayavo-meet-go/config"
)

func main() {
	cfg := &config.Config{
		APIKey:      "YOUR_API_KEY",
		AppID:       "YOUR_APP_ID",
		AppSecret:   "YOUR_APP_SECRET",

		Platform:   "backend",
		Identifier: "my-application",
		AppName:    "My Application",

		Timeout: 30 * time.Second,

		AllowVideo:  true,
		AllowAudio:  true,
		AllowScreen: true,
	}

	client := meet.New(cfg)

	ctx := context.Background()

	_ = client
	_ = ctx
	log.Println("Hayavo Meet initialized")
}
```

> **Important:** Keep your `APIKey` and `AppSecret` secure. Do not expose them in frontend or public client applications.

---

# Create a Room

Create a new Hayavo Meet room using `client.Rooms.Create()`.

```go
room, err := client.Rooms.Create(
	ctx,
	&models.CreateRoomRequest{
		RoomName: "my-meeting",
	},
)

if err != nil {
	log.Fatal(err)
}

log.Println("Room ID:", room.RoomID)
log.Println("Room Name:", room.RoomName)
```

You can also specify the maximum number of participants:

```go
room, err := client.Rooms.Create(
	ctx,
	&models.CreateRoomRequest{
		RoomName:        "my-meeting",
		MaxParticipants: 10,
	},
)
```

The SDK uses these defaults when they are not provided:

```text
Mode:            hm_default_meet_room
MaxParticipants: 5
```

---

# Generate Host Token

After creating a room, generate a host token using `client.Rooms.HostToken()`.

```go
hostToken, err := client.Rooms.HostToken(
	ctx,
	&models.HostTokenRequest{
		RoomName:    "my-meeting",
		DisplayName: "John",
	},
)

if err != nil {
	log.Fatal(err)
}

log.Println("Token:", hostToken.Token)
log.Println("Identity:", hostToken.Identity)
log.Println("Role:", hostToken.Role)
log.Println("RTC:", hostToken.Data.RTC)
log.Println("Signal:", hostToken.Data.Signal)
```

You can also provide custom attributes:

```go
hostToken, err := client.Rooms.HostToken(
	ctx,
	&models.HostTokenRequest{
		RoomName:    "my-meeting",
		DisplayName: "John",
		Attributes: map[string]string{
			"role":       "host",
			"department": "engineering",
		},
	},
)
```

### Host Token Request

```go
type HostTokenRequest struct {
	RoomName    string
	DisplayName string
	Attributes  map[string]string
}
```

Only `RoomName` is required.

---

# Generate Guest Token

Generate a guest token using `client.Rooms.GuestToken()`.

```go
guestToken, err := client.Rooms.GuestToken(
	ctx,
	&models.GuestTokenRequest{
		RoomName:    "my-meeting",
		AppID:       "YOUR_APP_ID",
		DisplayName: "Guest User",
	},
)

if err != nil {
	log.Fatal(err)
}

log.Println("Token:", guestToken.Token)
log.Println("Identity:", guestToken.Identity)
log.Println("Role:", guestToken.Role)
log.Println("RTC:", guestToken.Data.RTC)
```

You can also provide custom attributes:

```go
guestToken, err := client.Rooms.GuestToken(
	ctx,
	&models.GuestTokenRequest{
		RoomName:    "my-meeting",
		AppID:       "YOUR_APP_ID",
		DisplayName: "Guest User",
		Attributes: map[string]string{
			"role": "guest",
		},
	},
)
```

### Guest Token Request

```go
type GuestTokenRequest struct {
	RoomName    string
	AppID       string
	DisplayName string
	Attributes  map[string]string
}
```

`RoomName` and `AppID` are required.

---

# Complete Example

The complete basic flow is:

```go
package main

import (
	"context"
	"log"
	"time"

	meet "github.com/hayavo/hayavo-meet-go"
	"github.com/hayavo/hayavo-meet-go/config"
	"github.com/hayavo/hayavo-meet-go/models"
)

func main() {
	cfg := &config.Config{
		APIKey:      "YOUR_API_KEY",
		AppID:       "YOUR_APP_ID",
		AppSecret:   "YOUR_APP_SECRET",

		Platform:   "backend",
		Identifier: "my-application",
		AppName:    "My Application",

		Timeout: 30 * time.Second,

		AllowVideo:  true,
		AllowAudio:  true,
		AllowScreen: true,
	}

	client := meet.New(cfg)
	ctx := context.Background()

	// 1. Create room
	room, err := client.Rooms.Create(
		ctx,
		&models.CreateRoomRequest{
			RoomName:        "my-meeting",
			MaxParticipants: 10,
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Room created:", room.RoomName)

	// 2. Generate host token
	hostToken, err := client.Rooms.HostToken(
		ctx,
		&models.HostTokenRequest{
			RoomName:    room.RoomName,
			DisplayName: "Host",
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Host token:", hostToken.Token)

	// 3. Generate guest token
	guestToken, err := client.Rooms.GuestToken(
		ctx,
		&models.GuestTokenRequest{
			RoomName:    room.RoomName,
			AppID:       cfg.AppID,
			DisplayName: "Guest",
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Guest token:", guestToken.Token)

	// RTC endpoint
	log.Println("RTC:", hostToken.Data.RTC)

	// Signal endpoint
	log.Println("Signal:", hostToken.Data.Signal)
}
```

---

# Basic Flow

```text
Create Room
     │
     ▼
Generate Host Token
     │
     ├── Host joins RTC room
     │
     ▼
Generate Guest Token
     │
     └── Guest joins RTC room
```

The SDK provides the generated:

```text
Token
Identity
Role
RTC endpoint
Signal endpoint
```

for connecting your Hayavo Meet client application to the real-time infrastructure.

---

## License

Copyright © Hayavo.
