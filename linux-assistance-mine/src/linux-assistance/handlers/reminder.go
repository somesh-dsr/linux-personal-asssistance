package events

import (
	"os"
)

func (reminder *Reminder)InitiateEvent(){
	reminder.processReminderEvent()
}
func (reminder *Reminder)processReminderEvent(){

	 ref := setAndGetVoiceRef("reminder")
	 if len(os.Args) > 2 {
	 	reminder.userInfo = os.Args[2]
	 }
	 notoifyUser("meeting reminder")
	 ref.executeAndFilldata(reminder)
	 hour,min := deleteNode()
	 if hour == -1 || min == -1 {
	 	deleteCronEntry()
	 }else {
		 createCronEntry(getMintimer())
	 }
	 ActivityStatus<- struct{}{}
}
