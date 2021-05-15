package events

var ActivityStatus,Update chan struct{}
//may be need to remove
type voiceProcessingData struct {
	resultData string
	eventType string
	error error
}



type voiceProcessorEvents interface {
	executeAndFilldata(interface{})
	getUserVoiceCommands()
	postEventData(string2 string)
}
type EventsActions interface {
	InitiateEvent()
	//processEvent()
}
type VoiceEvents struct {
	EventType string
	resultData string
	error error
	status string
}
type Welcome struct {
	*VoiceEvents
	welcomeMsg string
	getuser chan struct{} // pending decision
}
type Reminder struct {
	userInfo string
	*VoiceEvents
	Notify chan struct{}

}
type CommonEvents struct {
	*VoiceEvents
	action string
	Communicate chan struct{
		Err error
		Result string
	}
	userInfo string
}
