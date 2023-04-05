package main

import (
	"fmt"
	"gocv.io/x/gocv"
)

func main() {
	done := make(chan bool)
	go openWindow(0, done)
	<-done
}

func openWindow(deviceID int, done chan bool) {
	// Open the webcam
	webcam, err := gocv.OpenVideoCapture(0)
	if err != nil {
		fmt.Println("Error opening video capture:", err)
		return
	}
	defer webcam.Close()

	// Create a window to display the camera feed
	window := gocv.NewWindow("Kosuha Camera")
	window.ResizeWindow(3840, 2160)
	// open windows in default screen
	window.MoveWindow(0, 0)
	defer window.Close()

	// Create an empty image matrix to store the camera frame
	frame := gocv.NewMat()
	defer frame.Close()

	// Loop until the window is closed
	for window.IsOpen() {
		// Read the camera frame
		if ok := webcam.Read(&frame); !ok {
			fmt.Println("Error reading from the camera")
			break
		}

		// Check if the frame is empty
		if frame.Empty() {
			continue
		}

		// Show the camera frame in the window
		window.IMShow(frame)

		// Break the loop if any key is pressed
		if window.WaitKey(1) >= 0 {
			break
		}
	}
	done <- true
}
