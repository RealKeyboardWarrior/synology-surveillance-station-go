![image](docs/static/header.png)

This is a Go client for interacting with the Synology Surveillance Station API. It allows you to:

✅ Log in to the Surveillance Station  
✅ List available cameras  
✅ Take snapshots from the cameras and save them as JPEG files  


---

## 📚 **Usage**

### 1. Import the package
In your Go file:

```go
import sssg "github.com/RealKeyboardWarrior/synology-surveillance-station-go"
```

---

### 2. Example Code  
Here's a complete example of logging in, listing cameras, and taking a snapshot:

```go
package main

import (
	"fmt"
	"os"
  sssg "github.com/RealKeyboardWarrior/synology-surveillance-station-go"
)

func main() {
	// Create a new client, if you don't have the TLS certificate installed for
  // your Synology NAS then you may see an error:
  // x509: certificate signed by unknown authority
  // An insecure workaround is to switch the last argument to true.
	client := sssg.NewClient("https://your-synology-ip:5001", "your-username", "your-password", false)

	// Log in to the Synology Surveillance Station
	if err := client.Login(); err != nil {
		fmt.Printf("Login failed: %v\n", err)
		return
	}

	// List available cameras
	cameras, err := client.ListCameras()
	if err != nil {
		fmt.Printf("Failed to list cameras: %v\n", err)
		return
	}

	fmt.Println("Available Cameras:")
	for _, cam := range cameras {
		fmt.Printf("ID: %d, Name: %s, IP: %s\n", cam.ID, cam.NewName, cam.IP)
	}

	// Take a snapshot of the first available camera
	if len(cameras) > 0 {
		camera := cameras[0]
		imageData, err := client.TakeSnapshot(camera)
		if err != nil {
			fmt.Printf("Failed to take snapshot: %v\n", err)
			return
		}

		// Save snapshot to a file
		fileName := fmt.Sprintf("%s-snapshot.jpg", camera.NewName)
		if err := os.WriteFile(fileName, imageData, 0644); err != nil {
			fmt.Printf("Failed to save snapshot: %v\n", err)
			return
		}

		fmt.Printf("Snapshot saved as %s\n", fileName)
	}
}
```

---

## 🔍 **API Reference**

### ✅ `NewClient(baseURL, username, password string) *SurveillanceStationClient`
Creates a new client instance.

### ✅ `Login() error`
Logs in to the Surveillance Station.

### ✅ `ListCameras() ([]Camera, error)`
Returns a list of available cameras.

### ✅ `TakeSnapshot(camera Camera) ([]byte, error)`
Takes a snapshot from the specified camera and returns the image data.

---

## 🧪 **Testing**

Run the tests:

```sh
go test -v
```

---

## 🏆 **Contributing**

Feel free to open issues or submit pull requests if you'd like to improve this project!

---

## 📄 **License**

This project is licensed under the [MIT License](LICENSE.md).

---

## ⚠️ **Disclaimer**  
This project is not affiliated with, endorsed by, or sponsored by Synology Inc.  
*Synology* and *Surveillance Station* are trademarks or registered trademarks of **Synology Inc.**  
All product names, logos, and brands are property of their respective owners.  
Use this software at your own risk. The authors are not responsible for any issues that may arise from its use.  
