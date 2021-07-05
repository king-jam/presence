package presence

import (
	"sync"
)

type UserStatus struct {
	meetingStatus MeetingStatus
	audioStatus   AudioStatus
	videoStatus   VideoStatus
	sharingStatus SharingStatus

	sync.RWMutex
}

type MeetingStatus int

const (
	MeetingStatusUnknown MeetingStatus = iota
	MeetingStatusBusy
	MeetingStatusAvailable
)

func (m MeetingStatus) String() string {
	return [...]string{"Unknown", "Busy", "Available"}[m]
}

type AudioStatus int

const (
	AudioStatusUnknown AudioStatus = iota
	AudioStatusOn
	AudioStatusOff
)

func (a AudioStatus) String() string {
	return [...]string{"Unknown", "On", "Off"}[a]
}

type VideoStatus int

const (
	VideoStatusUnknown VideoStatus = iota
	VideoStatusOn
	VideoStatusOff
)

func (v VideoStatus) String() string {
	return [...]string{"Unknown", "On", "Off"}[v]
}

type SharingStatus int

const (
	SharingStatusUnknown SharingStatus = iota
	SharingStatusOn
	SharingStatusOff
)

func (s SharingStatus) String() string {
	return [...]string{"Unknown", "Sharing", "NotSharing"}[s]
}

// type Controller struct {
// 	sync.RWMutex
// 	meetingStatus MeetingStatus
// 	audioStatus   AudioStatus
// 	videoStatus   VideoStatus
// 	sharingStatus SharingStatus
// }

// func NewUserStatusController() *Controller {
// 	return &Controller{}
// }

// func (u *Controller) Start() error {
// 	go func(z zoom) {
// 		ticker := time.NewTicker(10 * time.Second)
// 		for {
// 			select {
// 			case <-ticker.C:
// 				processes, err := ps.Processes()
// 				if err != nil {
// 					log.Println("failed to get processes")
// 				}
// 				for _, proc := range processes {
// 					name := proc.Executable()
// 					if name == "aomhost64.exe" || name == "CptHost.exe" {
// 						// if not already sent in a meeting
// 						if !z.isMeeting {
// 							err := z.notifier.JoinedMeeting()
// 							if err != nil {
// 								log.Println("failed to get processes")
// 							}
// 							z.isMeeting = true
// 						}
// 					} else {
// 						if z.isMeeting {
// 							err := z.notifier.LeftMeeting()
// 							if err != nil {
// 								log.Println("failed to get processes")
// 							}
// 						}
// 					}
// 				}
// 			case <-z.quit:
// 				return
// 			}
// 		}
// 	}(*z)
// }

func (u *UserStatus) MeetingStatus() MeetingStatus {
	u.RLock()
	defer u.RUnlock()
	return u.meetingStatus
}

func (u *UserStatus) SetMeetingStatus(ms MeetingStatus) {
	u.Lock()
	defer u.Unlock()
	u.meetingStatus = ms
}

func (u *UserStatus) AudioStatus() AudioStatus {
	u.RLock()
	defer u.RUnlock()
	return u.audioStatus
}

func (u *UserStatus) SetAudioStatus(as AudioStatus) {
	u.Lock()
	defer u.Unlock()
	u.audioStatus = as
}

func (u *UserStatus) VideoStatus() VideoStatus {
	u.RLock()
	defer u.RUnlock()
	return u.videoStatus
}

func (u *UserStatus) SetVideoStatus(vs VideoStatus) {
	u.Lock()
	defer u.Unlock()
	u.videoStatus = vs
}

func (u *UserStatus) SharingStatus() SharingStatus {
	u.Lock()
	defer u.Unlock()
	return u.sharingStatus
}

func (u *UserStatus) SetSharingStatus(ss SharingStatus) {
	u.Lock()
	defer u.Unlock()
	u.sharingStatus = ss
}

// type EventType string

// const (
// 	JoinedMeeting     EventType = "meeting.joined"
// 	LeftMeeting       EventType = "meeting.left"
// 	UnmutedMicrophone EventType = "microphone.unmuted"
// 	MutedMicrophone   EventType = "microphone.muted"
// 	StartedVideo      EventType = "video.started"
// 	StoppedVideo      EventType = "video.stopped"
// 	StartedPresenting EventType = "presentation.started"
// 	StoppedPresenting EventType = "presentation.stopped"
// )

// // Notifier is injected into monitors for their use in notifying
// // the presence detection system. Each monitor will call the below methods
// // upon changes. Calling each method will result in a domain event
// // being published throughout the system.
// type Notifier interface {
// 	JoinedMeeting() error
// 	LeftMeeting() error
// 	MutedMicrophone() error
// 	UnmutedMicrophone() error
// 	StartedVideo() error
// 	StoppedVideo() error
// 	StartedPresenting() error
// 	StoppedPresenting() error
// }
