package events

import "os"

func (welcome *Welcome)InitiateEvent(){
	welcome.processWelcomeEvent()
}
func (welcome *Welcome)processWelcomeEvent(){
	if len(os.Args) > 2 {
		welcome.welcomeMsg = os.Args[2]
	}
	ref := setAndGetVoiceRef("welcome")
	ref.executeAndFilldata(welcome)
	ActivityStatus<- struct{}{}
}