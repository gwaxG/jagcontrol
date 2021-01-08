package robot

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
)

type Robot struct {
	robInterface robotInterface
	robState State
	robRos rosMiddleware
}

func (r *Robot) Init(test bool, prms ...float64){
	r.robInterface = robotInterface{}
	r.robInterface.init(test)
	r.robState = State{}
	r.robState.init(prms)
	r.robRos = rosMiddleware{}
	r.robRos.init()
}

func (r *Robot) Close() {
	log.Println("Closing connections!")
	r.robInterface.close()
	r.robRos.close()
}

func  (r *Robot) catchKeyboardEvent(ch chan []byte) {
	b := make([]byte, 1)
	for {
		os.Stdin.Read(b)
		ch <- b
	}
}

func  (r *Robot) updateScreen(msg string){
	//fmt.Print("\033[H\033[2J")
	fmt.Print(r.robState.StringRepr())
	fmt.Println(msg)
}

func  (r *Robot) handleKeyPress(ch chan []byte, done chan<- bool) {
	var left, right, fr, rr, arm1, arm2 int32
	var msg string
	for key := range ch {
		switch key[0] {
		case 'z':
			left, right = r.robState.setVelocity(r.robState.linear + 0.10676, r.robState.angular)
			r.robInterface.writeVel(left, right)
		case 'a':
			left, right = r.robState.setVelocity(0.0, 0.0)
			r.robInterface.writeVel(left, right)
		case 's':
			left, right = r.robState.setVelocity(r.robState.linear - 0.10676, r.robState.angular)
			r.robInterface.writeVel(left, right)
		case 'q':
			left, right = r.robState.setVelocity(r.robState.linear, r.robState.angular - 0.10676)
			r.robInterface.writeVel(left, right)
		case 'd':
			left, right = r.robState.setVelocity(r.robState.linear, r.robState.angular + 0.10676)
			r.robInterface.writeVel(left, right)
		case 'r':
			fr, rr = r.robState.incFlipper(0.157, 0.0)
			r.robInterface.writeFlip(fr, rr)
		case 'f':
			fr, rr = r.robState.incFlipper(-0.157, 0.0)
			r.robInterface.writeFlip(fr, rr)
		case 't':
			fr, rr = r.robState.incFlipper(0.0, 0.157)
			r.robInterface.writeFlip(fr, rr)
		case 'g':
			fr, rr = r.robState.incFlipper(0.0, -0.157)
			r.robInterface.writeFlip(fr, rr)
		case 'y':
			arm1, arm2 = r.robState.incArm(0.157, 0.0)
			r.robInterface.writeArm(arm1, arm2)
		case 'h':
			arm1, arm2 = r.robState.incArm(-0.157, 0.0)
			r.robInterface.writeArm(arm1, arm2)
		case 'u':
			arm1, arm2 = r.robState.incArm(0.0, 0.157)
			r.robInterface.writeArm(arm1, arm2)
		case 'j':
			arm1, arm2 = r.robState.incArm(0.0, -0.157)
			r.robInterface.writeArm(arm1, arm2)
		case 'x':
			r.robInterface.release()
		case 'p':
			exec.Command("stty", "-F", "/dev/tty", "echo").Run()
			done <- true
		}
		r.updateScreen(msg)
	}
}

func (r *Robot) Keyboard(done chan<- bool) {
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	ch := make(chan []byte, 10)
	go r.catchKeyboardEvent(ch)
	go r.handleKeyPress(ch, done)
}

func (r *Robot) handleRPC(msg []byte, res chan []byte, done chan<- bool) {
	var request map[string]interface{}
	json.Unmarshal(msg, &request)
	if _, ok := request["linear"]; ok {
		// -|angular| <- turn left
		left, right := r.robState.setVelocity(request["linear"].(float64), request["angular"].(float64))
		r.robInterface.writeVel(left, right)
	}
	if _, ok := request["front"]; ok {
		fr, rr := r.robState.incFlipper(request["front"].(float64), request["rear"].(float64))
		r.robInterface.writeFlip(fr, rr)
	}
	if _, ok := request["arm1"]; ok {
		arm1, arm2 := r.robState.incArm(request["arm1"].(float64), request["arm2"].(float64))
		r.robInterface.writeFlip(arm1, arm2)
	}
	if _, ok := request["cmd"]; ok {
		if request["cmd"] == 0.0 {
			done <- true
		}
	}
	state := r.robState.getState()
	jsonState, _ := json.Marshal(state)
	res <- jsonState
}

func (r *Robot) handleSensorData(sensorChan <-chan string) {
	for dat := range sensorChan {
		fmt.Println(dat)
	}
}

func (r *Robot) Sensors() {
	sensorData := make(chan string)
	go r.robInterface.readSensors(sensorData)
	go r.handleSensorData(sensorData)
}

func (r *Robot) RPC(done chan<- bool) {
	req := make(chan []byte)
	res := make(chan []byte)
	go r.robRos.consume(req, res)
	for {
		msg := <-req
		go r.handleRPC(msg, res, done)
	}
}


func (r *Robot) Middleware() {

}

