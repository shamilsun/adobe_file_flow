package pathwatcher

import (
	"../log"
	"../scene"
	"github.com/fsnotify/fsnotify"
)

func StartWatcher(path string, onNewFile func(path string)) (*fsnotify.Watcher, func(), error) {
	log.AddLog("trying to start watching")
	log.AddLog(path)

	watcher, err := fsnotify.NewWatcher(100000)
	if err != nil {
		log.AddLog("failed")
		return watcher, nil, err
	}
	isRunning := true

	var onClose = func() {
		isRunning = false
		watcher.Close()
	}

	go func() {
		for {
			if scene.IsUserClosedApp() {
				onClose()
				return
			}

			select {

			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Create == fsnotify.Create {
					onNewFile(event.Name)
				}

			case err, ok := <-watcher.Errors:
				if !ok {
					//log.Println("!ok")
					//log.Println(ok)
					//log.Println(err)
					return
				}
				log.Println("error file watcher:")
				log.Println(err.Error())

			default:
				if !isRunning {
					log.Println("close file watcher")
					return
				}
			}

		}
	}()

	err = watcher.Add(path)

	if err != nil {
		return watcher, onClose, err
	}

	return watcher, onClose, err
}
