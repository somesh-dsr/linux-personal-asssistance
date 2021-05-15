package events

import (
	htgotts "github.com/hegedustibor/htgo-tts"
)

func (vProcessindata *voiceProcessingData)postEventData(data string){
	//os.Setenv("DISPLAY",":0")
	//exec.Command("python3",data).Output()
    //exec.Command("espeak","-s","120",data).Output()
	//fmt.Println("epeak",string(out),err,data)
	speech := htgotts.Speech{Folder: "audio", Language: "en"}
	vProcessindata.error =  speech.Speak(data)
}