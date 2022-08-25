package notifier

import (
"log"
"sync"
)


//Notifier notifies about the sleep/wake events
type Notifier struct {
	quit      chan struct{}
	mutex     sync.RWMutex
	isRunning bool
}

//Type determines if it is a sleep or an awake type activity
type Type string

/*
The 2 types of activities
*/
const (
	Sleep Type = "sleep"
	Awake Type = "awake"
)

//Activity struct is used to hold the sleep and awake activity.
type Activity struct {
	Type Type
}

var instance *Notifier
var notifierCh chan *Activity

//GetInstance gets the singleton instance for the notifier
func GetInstance() *Notifier {
	if instance == nil {
		instance = &Notifier{}
	}
	return instance
}

//Start the notifier. It returns an Activity channel
//to listen for machine sleep/wake activities.
func (n *Notifier) Start() chan *Activity {

	n.quit = make(chan struct{})
	notifierCh = make(chan *Activity)

	go func(n *Notifier) {
		n.setIsRunning(true)
		StartNotifier()
	}(n)

	go func(n *Notifier) {
		for {
			select {
			case <-n.quit:
				log.Printf("quitting notifier")
				n.setIsRunning(false)
				StopNotifier()
				return
			}
		}
	}(n)
	return notifierCh
}

//Quit the notifier
func (n *Notifier) Quit() {
	n.quit <- struct{}{}
}

//setIsRunning sets status of notifier
func (n *Notifier) setIsRunning(status bool) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	n.isRunning = status
}

//isStatusRunning checks running status of notifier
func (n *Notifier) isStatusRunning() bool {
	n.mutex.RLock()
	defer n.mutex.RUnlock()
	return n.isRunning
}
