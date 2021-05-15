package events

import (
	"errors"
	"fmt"
	detect "linux-assistance/detector"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func (allEvents *CommonEvents)InitiateEvent(){
	allEvents.processEvents()
}
func (allEvents *CommonEvents)processEvents(){
	ref := setAndGetVoiceRef("generic")
	allEvents.action = "get Usercommand"
	ref.executeAndFilldata(allEvents)
	if allEvents.error != nil {
		fmt.Println("Faied to fetch data.",allEvents.error)
	}
	data, eventType := detect.GetActivityType(strings.TrimSpace(allEvents.resultData))
	fmt.Println(data,eventType)
	if eventType == "" {
		allEvents.action = "post data"
		allEvents.userInfo = "Sorry I did not get you. We are taking this input for future reference. "
		ref.executeAndFilldata(allEvents)
		ActivityStatus<- struct{}{}
		return
	}
	if eventType == "update" {
		go allEvents.updateDevice()
		finalResult := <-allEvents.Communicate
		allEvents.userInfo = finalResult.Result
		allEvents.action = "post data"
		//fmt.Println(allEvents)
		if finalResult.Err != nil {
			notoifyUser("update-failed")
			ref.executeAndFilldata(allEvents)
		}else {
			notoifyUser("update-success")
			ref.executeAndFilldata(allEvents)
		}
		Update = make(chan struct{})
		Update <- struct{}{}
	}else if eventType == "meeting schedule" {
		var daytype, deviceTimeStamp string
		deviceTimeStamp = data[0][len(data[0])-2]
		if strings.ToLower(data[0][len(data[0])-2]) == "am" ||  strings.ToLower(data[0][len(data[0])-2]) == "pm" {
			deviceTimeStamp = data[0][len(data[0])-3]
			daytype = data[0][len(data[0])-2]
		}
		allEvents.schedule(deviceTimeStamp,daytype)
		allEvents.action = "post data"
		saveEventsTimers()
		ref.executeAndFilldata(allEvents)
		ActivityStatus<- struct{}{}
	}else {
		ActivityStatus <- struct{}{}
	}
}
func (allEvents *CommonEvents)schedule(timer string,dayType string){
	timers := strings.Split(timer,":")
	fmt.Println("scheduler: ",timer)
	if len(timers) < 2{
		allEvents.userInfo = "Failed to shedule the event"
		return
	}
	hours, err := strconv.Atoi(strings.TrimSpace(timers[0]))
	if err != nil {
		allEvents.userInfo = "Failed to shedule the event"
		return
	}
	minutes,err := strconv.Atoi(strings.TrimSpace(timers[1]))
	if err != nil {
		allEvents.userInfo = "Failed to shedule the event"
		return
	}
	fmt.Println("scheduling")
	hour,minute := allEvents.verifyandConvertFomrat(hours,minutes,dayType)
	if hour == -1 {
		allEvents.userInfo = "Meeting time already exceeded."
		return
	}
	_ = insertNode(hour,minute)
	createCronEntry(hour,minute)
	allEvents.userInfo = "Successfully scheduled the event."

}
func (allevents *CommonEvents)verifyandConvertFomrat(hour,minute int, daytype string) (int,int) {
	timer := time.Now()
	if daytype == "am" {

		if hour <= timer.Hour() && minute <= timer.Minute(){
			return -1,-1
		}

		return hour, minute

	}else if daytype == "pm" {
			if hour < 12 {
				hour += 12
			}
			if hour <= timer.Hour() && minute <= timer.Minute() || hour > 24 {
				return -1,-1
			}

	}else if daytype == "" {
		fmt.Println(timer.Hour(),timer.Minute())
		if hour <= (timer.Hour()%12) && minute <= timer.Minute() {
			return -1,-1
		}
		if timer.Hour()>12 && hour < 12 && hour >= (timer.Hour()%12) {
			hour += 12
		}
	}

	if minute == 0 {
		hour -= 1
		minute = 55
	}else {
		minute -= 5
	}
	if minute < timer.Minute() {
		minute = timer.Minute() + 1
	}
	return hour,minute


}
func (allEvents *CommonEvents)updateDevice(){
	executablePath := ""
	if path,err := exec.LookPath("apt-get"); err == nil {
			executablePath = path
	}else if path,err = exec.LookPath("yum"); err == nil {
		executablePath = path
	}else if path,err = exec.LookPath("dnf"); err == nil {
		executablePath = path
	}else if path, err = exec.LookPath("zypper"); err == nil {
		executablePath = path
	}
	if executablePath == "" {
		allEvents.Communicate <- struct {
			Err    error
			Result string
		}{Err: errors.New("Unable to get binary path"), Result: "" }
		return
	}
	_,err := exec.Command(executablePath,"-y","update").CombinedOutput()
	msg := "Successfully Updated your OS"
	fmt.Println(err)
	if err != nil {
		msg = "Failed to update the OS."+"\n"+"Getting: "+err.Error()
	}
	allEvents.Communicate<- struct{
			Err error
			Result string
		}{
			Err: err,
			Result: msg,
	}

}

// su somesh -c 'notify-send "OS updated successfully"'