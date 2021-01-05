package robot

type Robot struct {
	robInterface robotInterface
	robState State
}

func (r *Robot) Init(prms ...int32){
	r.robInterface = robotInterface{}
	r.robInterface.init()
	r.robState = State{}
	r.robState.init(prms)
}

func (r *Robot) Str() string {
	return "Ok"
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
