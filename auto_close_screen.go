package main

import (
	"fmt"
	"image/color"
    "os/exec"
	"gocv.io/x/gocv"
    "bytes"
    "time"
)


func exec_shell(s string) (string, error){
    cmd := exec.Command("/bin/bash", "-c", s)
    var out bytes.Buffer
    cmd.Stdout = &out
    err := cmd.Run()
    return out.String(), err
}

func main() {
    // set to use a video capture device 0
    deviceID := 0

	// open webcam
	webcam, err := gocv.OpenVideoCapture(deviceID)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer webcam.Close()

	// open display window
	window := gocv.NewWindow("Face Detect")
	defer window.Close()

	// prepare image matrix
	img := gocv.NewMat()
	defer img.Close()

	// color for the rect when faces detected
	blue := color.RGBA{0, 0, 255, 0}

	// load classifier to recognize faces
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()

	if !classifier.Load("/tmp/haarcascade_frontalface_default.xml") {
		fmt.Println("Error reading cascade file: data/haarcascade_frontalface_default.xml")
		return
	}

	fmt.Printf("start reading camera device: %v\n", deviceID)
    var i int
    // var closed bool
	for {
        fmt.Printf("running")
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("cannot read device %v\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}

		// detect faces
		rects := classifier.DetectMultiScale(img)
		// fmt.Printf("found %d faces\n", len(rects))

		// draw a rectangle around each face on the original image
		if len(rects) == 0 {
            i = i + 1
            if i > 5 {
                //if !closed {
                fmt.Printf("close the screen\n")
                exec_shell("pmset displaysleepnow")
                //closed = true
                //}
                if i > 50 {
                    // return
                }
            }
        } else {
            i = 0
            for _, r := range rects {
			    gocv.Rectangle(&img, r, blue, 3)
		    }
        }

        time.Sleep(time.Duration(100)*time.Millisecond)
		// show the image in the window, and wait 1 millisecond
		window.IMShow(img)
		window.WaitKey(1)
	}
}
