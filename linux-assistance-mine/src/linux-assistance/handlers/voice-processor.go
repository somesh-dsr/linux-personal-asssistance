package events

import (
	"sync"
)

func (vProcessingdata *voiceProcessingData)executeAndFilldata (task interface{}){
	if vProcessingdata.eventType == "welcome" {
		event := task.(*Welcome)
		vProcessingdata.postEventData(event.welcomeMsg)
		event.error =vProcessingdata.error
	}else if vProcessingdata.eventType == "reminder"{
		event := task.(*Reminder)
		vProcessingRef.postEventData(event.userInfo)
		event.error = vProcessingdata.error
	}else if vProcessingdata.eventType == "generic"{
		event := task.(*CommonEvents)
		if event.action == "get Usercommand" {
			vProcessingdata.postEventData("Hi, How can I help you?")
			vProcessingdata.getUserVoiceCommands()
			event.resultData = vProcessingRef.resultData
			event.error = vProcessingRef.error
		}else if event.action == "post data"{
			vProcessingdata.postEventData(event.userInfo)
			event.error = vProcessingdata.error
		}
	}
}
var vProcessingRef *voiceProcessingData
func setAndGetVoiceRef(eventType string)voiceProcessorEvents{
	once := sync.Once{}
	once.Do(func(){
		vProcessingRef = &voiceProcessingData{}
	})
	vProcessingRef.eventType = eventType
	return vProcessingRef
}
