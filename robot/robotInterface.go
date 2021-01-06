package robot

import (
	"net"
	"strconv"
)

/*
	This package is mentioned to provide communication with the robot:
		I maintain tcp/ip connection with the robot,
		  transform sending data into a string protocol command,
		  serialize command,
		  send it.
		II receive information about the robot base state.
*/


type robotInterface struct {
	connectionBase net.Conn
	connectionArm net.Conn
}

func (p *robotInterface) init(){
	// 192.168.0.60:10001 base
	// 192.168.0.63:10001 arm
	// connection to base
	conn_base, err := net.Dial("tcp", "localhost:10001")
	if err != nil {
		panic("No connection to the robot base!")
	}
	// connection to arm
	conn_arm, err := net.Dial("tcp", "localhost:10002")
	if err != nil {
		panic("No connection to the robot arm!")
	}
	p.connectionBase = conn_base
	p.connectionArm = conn_arm
}

func (p *robotInterface) close() {
	var err interface{}
	if p.connectionArm != nil {
		err = p.connectionArm.Close()
		if err != nil {
			panic("Can not close the arm connection")
		}
		p.connectionArm = nil
	}
	if p.connectionBase != nil {
		err = p.connectionBase.Close()
		p.connectionBase = nil
		if err != nil {
			panic("Can not close the base connection")
		}
	}
}

func (p *robotInterface) readBattery() {}

func (p *robotInterface) readCurrent() {}

func (p *robotInterface) writeVel(left, right int32) {
	var cmd string = "MMW !M " + strconv.Itoa(int(left)) + " " + strconv.Itoa(int(right)) + "\r\n"
	_, err := p.connectionBase.Write([]byte(cmd))
	if err != nil {
		panic("Can not write velocity.")
	}
	// fmt.Println("Written %d bytes %s", n)
}

func (p *robotInterface) writeFlip(fr, rr int32) {
	var cmd string
	// front
	// "MM2 !PR 1 " + QString::number(leftFrontCmd) + "\r\n"
	// "MM2 !PR 2 " + QString::number(rightFrontCmd) + "\r\n";
	cmd = "MM2 !PR 1 " + strconv.Itoa(int(fr)) + "\r\n"
	_, err := p.connectionBase.Write([]byte(cmd))
	if err != nil {
		panic("Can not write left front flipper action")
	}
	cmd = "MM2 !PR 2 " + strconv.Itoa(int(fr)) + "\r\n"
	_, err = p.connectionBase.Write([]byte(cmd))
	if err != nil {
		panic("Can not write right front flipper action")
	}
	// rear
	// "MM3 !PR 1 " + QString::number(leftRearCmd) + "\r\n";
	// "MM3 !PR 2 " + QString::number(rightRearCmd) + "\r\n"
	cmd = "MM3 !PR 1 " + strconv.Itoa(int(rr)) + "\r\n"
	_, err = p.connectionBase.Write([]byte(cmd))
	if err != nil {
		panic("Can not write left rear flipper action")
	}
	cmd = "MM3 !PR 2 " + strconv.Itoa(int(rr)) + "\r\n"
	_, err = p.connectionBase.Write([]byte(cmd))
	if err != nil {
		panic("Can not write right rear flipper action")
	}
}

func (p *robotInterface) writeArm(arm1, arm2 int32) {
	// '!PR 1 '+str(joint1)+'\r'
	// '!PR 2 '+str(joint1)+'\r'
	var cmd string
	cmd = "!PR 1 " + strconv.Itoa(int(arm1)) + "\r"
	_, err := p.connectionArm.Write([]byte(cmd))
	if err != nil {
		panic("Can not write an action to the first arm joint")
	}
	cmd = "!PR 2 " + strconv.Itoa(int(arm2)) + "\r"
	_, err = p.connectionArm.Write([]byte(cmd))
	if err != nil {
		panic("Can not write an action to the second arm joint")
	}
}

/*
void MainWindow::wheelCmdSend(int cmdValue1,int cmdValue2)
{
	QString strCmd;


	if ((cmdValue1 < -1000) || (cmdValue2 < -1000))
	{
		strCmd = "MMW !EX\r\n";
	}
	else if((cmdValue1 > 1000) || (cmdValue2 > 1000))
	{
		strCmd = "MMW !MG\r\n";
	}
	else
	{
		strCmd = "MMW !M " + QString::number(cmdValue1) + " " + QString::number(cmdValue2) + "\r\n";
	}

	if (tcpRobot != NULL){
		if (tcpRobot->isWritable())
		{
		    tcpRobot->write(strCmd.toUtf8().constData());
		}
	}
}
void MainWindow::flipCmdSend (int leftFrontCmd,int rightFrontCmd,int leftRearCmd,int rightRearCmd)
{
	QString strCmd;
	strCmd = "MM2 !PR 1 " + QString::number(leftFrontCmd) + "\r\n";
	if (tcpRobot != NULL){
		if (tcpRobot->isWritable())
		{
		    tcpRobot->write(strCmd.toUtf8().constData());
		}
	}

	strCmd = "MM2 !PR 2 " + QString::number(rightFrontCmd) + "\r\n";
	if (tcpRobot != NULL){
		if (tcpRobot->isWritable())
		{
		    tcpRobot->write(strCmd.toUtf8().constData());
		}
	}

	strCmd = "MM3 !PR 1 " + QString::number(leftRearCmd) + "\r\n";
	if (tcpRobot != NULL){
		if (tcpRobot->isWritable())
		{
		    tcpRobot->write(strCmd.toUtf8().constData());
		}
	}

	strCmd = "MM3 !PR 2 " + QString::number(rightRearCmd) + "\r\n";
	if (tcpRobot != NULL){
		if (tcpRobot->isWritable())
		{
		    tcpRobot->write(strCmd.toUtf8().constData());
		}
	}
}
*/
