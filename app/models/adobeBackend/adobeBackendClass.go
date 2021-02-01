package adobeBackend

import (
	"encoding/json"
	"github.com/parnurzeal/gorequest"
	"log"
	"time"
)
import "../scene"

var requestTimeDelayInSeconds int

func GetRequestTimeDelayInSeconds() int {
	return requestTimeDelayInSeconds
}

func (req *IAdobeJOBRequest) JobOnAdobe() {

	for {
		if scene.IsUserClosedApp() {
			return
		}

		startAt := time.Now()

		request := gorequest.New()

		res2B, _ := json.Marshal(req)

		log.Println("req adobe:")
		log.Println(string(res2B))

		resp, _, _ := request.Post(AdobeExtensionApiUrlSetting.GetValue()).
			Send(req).
			//		Send(`{"Safari":"5.1.10"}`).
			End()

		if resp != nil && resp.StatusCode == 200 {
			req.OnDone(resp, req)
			requestTimeDelayInSeconds = int(time.Now().Sub(startAt).Seconds())
			return
		}

		req.OnError()
		log.Println("resp", resp)
		//		log.Println(resp)
		//time.Sleep(time.Second)
	}
}

func (req *IAdobeJOBRequest) onFinish() {
	log.Println("adobe finished job")
}

func (req *IAdobeJOBRequest) onFail() {
	log.Println("adobe fail job")
}

func worker() {
	defer waitGroup.Done()
	log.Println("Worker is waiting for jobs")
	for {

		if scene.IsUserClosedApp() {
			return
		}

		select {

		case job, ok := <-AdobeQueue:
			if !ok {
				return
			}
			log.Println("Worker picked job", job)
			//doAdobeRequest(&job)
			job.JobOnAdobe()
		default:
			//if scene.IsUserClosedApp() {
			//	break
			//}
		}
	}
}

func Start() {
	log.Println("start channel to adobe")
	if AdobeQueue == nil {
		AdobeQueue = make(chan IAdobeJOBRequest, 1000000)
		waitGroup.Add(1)

		// Run 1 worker to handle jobs.
		go worker()
	}
}
