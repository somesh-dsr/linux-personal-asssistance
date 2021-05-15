package processor

import (
	"linux-assistance/handlers"
)

func ProcessRequests(activity string){
	if ref := initiateReferenceFactory(activity); ref != nil {
		go ref.InitiateEvent()
	}else{
		events.ActivityStatus<- struct{}{}
	}
}
func initiateReferenceFactory(activity string)events.EventsActions{
	type Sample struct {
			err error
			result string

	}
	switch activity {
	case "welcome":
			return &events.Welcome{
				VoiceEvents: &events.VoiceEvents{
					EventType: "welcome",
				},
			}
	case "reminder" :
		  return  &events.Reminder{
				VoiceEvents: &events.VoiceEvents{
					EventType: "reminder",
				},
				Notify: make(chan struct{}),
		  }

	default:
		return  &events.CommonEvents{
				VoiceEvents: &events.VoiceEvents{
					EventType: "generic",
				},
				Communicate: make(chan struct{
					Err error
				    Result string
				}),
		}

	}

}