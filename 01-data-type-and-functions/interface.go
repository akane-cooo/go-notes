package ch01

import (
	"fmt"
)

type Notifier interface {
	call()
}

type Navigator interface {
	// interface wont will only have methods
	move(duration int, speed int) int
	rotate(angle int) int
	Notifier
}

type RobotInfo struct {
	Type string
	Name string
}

type WheelRobot struct {
	RobotInfo
}

func (r *WheelRobot) move(duration int, speed int) int {
	return speed * duration
}

func (r *WheelRobot) rotate(angle int) int {
	return angle % 360
}

func (r *WheelRobot) Beep() {
	fmt.Printf("%s, Beep!\n", r.Name)
}

func (r *WheelRobot) call() {
}

type HumanRobot struct {
	RobotInfo
	MaxSpeed int
	MaxAngle int
}

func (r *HumanRobot) move(duration int, speed int) int {
	speed = min(speed, r.MaxSpeed)
	return speed * duration
}

func (r *HumanRobot) rotate(angle int) int {
	angle = min(angle%360, r.MaxAngle)
	return angle
}

func (r *HumanRobot) call() {
}

func moveRobot(name string, navigator Navigator) {
	distance := navigator.move(10, 20)
	fmt.Printf("%s moved %d\n", name, distance)
}

func interfaceBasic() {
	wheelRobot := WheelRobot{RobotInfo{Name: "John", Type: "wheel"}}
	humanRobot := HumanRobot{RobotInfo: RobotInfo{Name: "Tim", Type: "human"}, MaxSpeed: 2, MaxAngle: 30}
	moveRobot(wheelRobot.Name, &wheelRobot)
	moveRobot(humanRobot.Name, &humanRobot)
}

func InterfaceRun() {
	interfaceBasic()
}
