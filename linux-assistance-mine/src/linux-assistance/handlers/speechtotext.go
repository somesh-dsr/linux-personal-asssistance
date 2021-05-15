package events

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

func (vProcessindata *voiceProcessingData)getUserVoiceCommands(){
	pythonpath := ""
	if path,err := exec.LookPath("python3"); path !="" && err == nil {
		pythonpath = path
	}else if path,err = exec.LookPath("python2"); path !="" && err == nil {
		pythonpath = path
	}else if path,err = exec.LookPath("python"); path !="" && err == nil {
		pythonpath = path
	}
	//cmd := exec.Cmd{Stdout: os.Stdout,
	//	Stderr: os.Stderr}
	/*path,err := os.Getwd()
	if err != nil || path == "" {
		path = "/opt/linux-personal-assistance/command-handlers/speechtotext.py"
	}else {
		path += "/command-handlers/speechtotext.py"
	}*/
	path := "/opt/linux-personal-assistance/command-handlers/speechtotext.py"
	if _,err := os.Stat(path); err != nil {
		fmt.Println("file not found.",err)
		vProcessingRef.error = errors.New("Speech conversion file not found.")
		return
	}
	vCmd,err :=  exec.Command(pythonpath,path).Output()
	fmt.Println(string(vCmd),err)
	vProcessingRef.error = err
	vProcessingRef.resultData = string(vCmd)

}
