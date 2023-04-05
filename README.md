# CAM

A simple camera app for macOS that displays the webcam feed in a window. The app runs independently, without the terminal, and closes when the window's close button is clicked.

## Setup

Follow these steps to set up and run the app:

1. Clone the repository:
   https://github.com/koushamad/cam
2. Navigate to the directory containing the Go code and build the app:
    ```bash
    brew install cmake opencv
    cd cam
    go build cam.go
    ln -s cam /usr/local/bin/cam
    ```
3. Run the app:
   ```bash
      cam
      ```
Now, the camera app will open in a window. You can close the app by clicking on the window's close button.

### License

MIT License
