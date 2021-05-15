package detect

import (
	"regexp"
)

func GetActivityType(voiceCommand string)([][]string,string){
	activityRegex := make(map[string]string)
	activityRegex[`(?m)(update|upgrade)\s*(my)?\s*(device|laptop|server|ubuntu|linux|rpm|desktop|os)\s*(device|os)?`] = "update"
	activityRegex[`(?m)(please)?(schedule|set)?\s*(a|an|the)?\s*(meeting|reminder)\s*\s*(at|by)\s*(\d?\d:\d?\d?)\s*(pm|am)?$`] = "meeting schedule"
	activityRegex[`(?m)(my)?(user|username)?(is)?([a-z]+)$`] = "user name"

	finalActivity := ""
	regex := ""
	for key,value := range activityRegex {
		if status,_ := regexp.Match(key,[]byte(voiceCommand));status{
			finalActivity = value
			regex = key
			break
		}
	}
	if finalActivity == "meeting schedule" {
		return getDetecteddata(regex,voiceCommand),finalActivity
	}else if finalActivity == "user name" {
		return getDetecteddata(regex,voiceCommand),finalActivity
	}

	return nil,finalActivity

}
func getDetecteddata(regex string,data string)[][]string{
	if ref, err := regexp.Compile(regex); err != nil {
		return nil
	}else {
		return  ref.FindAllStringSubmatch(data,-1)
	}

}