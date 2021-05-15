package events

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
)

var heapHours, heapMinutes []int
func loadHeap(){
	/*path,err := os.Getwd()
	if err != nil || path == "" {
		path = "/opt/linux-personal-assistance/backup/scheduler"
	}else {
		path += "backup/scheduler"
	}*/
	os.MkdirAll( "/opt/linux-personal-assistance/backup/",0755)
	path := "/opt/linux-personal-assistance/backup/scheduler"
	if _,err := os.Stat(path); err != nil {
		return
	}
	data,err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
    hours,minutes := strings.Split(strings.Split(string(data),"#")[0],"$$"),strings.Split(strings.Split(string(data),"#")[1],"$$")
    for i:=1;i<len(hours);i++ {
		hour ,_ := strconv.Atoi(hours[i])
		min ,_ := strconv.Atoi(minutes[i])
		heapHours = append(heapHours,hour)
		heapMinutes = append(heapMinutes,min)
	}
}
func insertNode(hour,minute int)error{
	loadHeap()
	heapHours,heapMinutes = append(heapHours,hour),append(heapMinutes,minute)
	i := (len(heapHours) - 1)/2
	for ; i > 0 && ((heapHours[i]<heapHours[(i-1)/2]) || (heapHours[i] == heapHours[(i-1)/2])&&((heapMinutes[i] == heapMinutes[(i-1)/2]) || (heapMinutes[i] < heapMinutes[(i-1)/2])));i= (i-1)/2 {
		if heapHours[i] == heapHours[(i-1)/2] && heapMinutes[i] == heapHours[(i-1)/1] {
			heapHours=heapHours[0:i]
			heapMinutes=heapMinutes[0:i]
			return errors.New("Already you have a scheduled meeting on the same time.")
		}else {
			heapHours[i],heapHours[(i-1)/2] = heapHours[(i-1)/2],heapHours[i]
			heapMinutes[i],heapMinutes[(i-1)/2] = heapMinutes[(i-1)/2],heapMinutes[i]
		}
	}

	return nil
}
func deleteNode()(int,int){
	if len(heapHours) == 0 {
		return -1,-1
	}
	loadHeap()
	hour,minute := heapHours[0],heapMinutes[0]
	if len(heapHours) == 1 {
		heapHours,heapMinutes = nil,nil
		return hour,minute
	}
	heapHours[0],heapMinutes[0] = heapHours[len(heapHours)-1],heapMinutes[len(heapMinutes)-1]
	heapHours,heapMinutes = heapHours[0:len(heapHours)-1],heapMinutes[0:len(heapMinutes)-1]
	heapify(0)
	return hour,minute
}
func heapify(index int){
	for i:=index;i < len(heapHours);i++ {
		min := i
		leftNode  := 2*i + 1
		rightNode := 2*i + 2
		if leftNode < len(heapHours)&& ((heapHours[min] > heapHours[leftNode]) || ((heapHours[min] == heapHours[leftNode]) && (heapMinutes[min] > heapMinutes[leftNode]) )) {
			min = leftNode
		}
		if rightNode < len(heapHours) && ((heapHours[min] > heapHours[rightNode]) || ((heapHours[min] == heapHours[rightNode]) && (heapMinutes[min] > heapMinutes[rightNode]) )) {
			min = rightNode
		}
		if min != i {
			heapHours[i],heapHours[min] = heapHours[min], heapHours[i]
			heapMinutes[i],heapMinutes[min] = heapMinutes[min], heapMinutes[i]
		}else {
			break
		}
	}
}
func getMintimer()(int,int){
	if len(heapHours) <=0 {
		return -1,-1
	}
	return heapHours[0],heapMinutes[0]
}
func notoifyUser(status string){
	exec.Command("bash","/opt/linux-personal-assistance/notifiers/notify-user.sh",status).Output()
}
func saveEventsTimers(){
	hourList,minList := "",""
	for index,_ := range heapHours {
		hourList += fmt.Sprintf("%d$$",heapHours[index])
		minList  += fmt.Sprintf("%d$$",heapMinutes[index])
	}
	ioutil.WriteFile("/opt/linux-personal-assistance/backup/scheduler",[]byte(strings.TrimRight(hourList,"$$")+"#"+strings.TrimRight(minList,"$$")),0644)
}
func createCronEntry(hour,min int){
	locker := sync.RWMutex{}
	locker.Lock()
	defer locker.Unlock()
	if hour == -1 || min == -1 {
		return
	}
	cronData,err := ioutil.ReadFile("/etc/crontab")
	if err != nil {
		return
	}
	if !strings.Contains(string(cronData),"/opt/linux-personal-assistance/linux-assistance") {
		ioutil.WriteFile("/etc/crontab",[]byte(string(cronData)+"\n"+fmt.Sprintf("%d %d * * * root /opt/linux-personal-assistance/linux-assistance 'reminder' 'you have a meeting in next few minutes'"+"\n",min,hour)),0644)
	}else {
		cronSlice :=strings.Split(string(cronData),"\n")
		for i := len(cronSlice)-1;i >=0;i-- {
			if strings.Contains(cronSlice[i],"/opt/linux-personal-assistance/linux-assistance"){
				ioutil.WriteFile("/etc/crontab",[]byte(strings.Replace(string(cronData),cronSlice[i],fmt.Sprintf("%d %d * * * root /opt/linux-personal-assistance/linux-assistance 'reminder' 'you have a meeting in next few minutes'",min,hour),-1)),0644)
				break
			}
		}
	}
}
func deleteCronEntry(){
	locker := sync.RWMutex{}
	locker.Lock()
	defer locker.Unlock()
	cronData,err := ioutil.ReadFile("/etc/crontab")
	if err != nil {
		return
	}
	cronSlice :=strings.Split(string(cronData),"\n")
	for i := len(cronSlice)-1;i >=0;i-- {
		if strings.Contains(cronSlice[i],"/opt/linux-personal-assistance/linux-assistance"){
			ioutil.WriteFile("/etc/crontab",[]byte(strings.Replace(string(cronData),cronSlice[i],"",-1)),0644)
			break
		}
	}
}