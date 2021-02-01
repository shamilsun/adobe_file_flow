package photoSession

import (
	"../adobeBackend"
	"../database"
	"../dotenv"
	"../email"
	"../helpers"
	"../log"
	"../preset"
	"../scene"
	"../yandexCloud"
	"fmt"
	//	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	//"log"
	"path/filepath"
	//	"path/filepath"
	"time"
)
import "../pathwatcher"

var db = database.GetDB()
var fileNumber uint64 = 1

var sessionsList []*IPhotosSession

func (s *IPhotosSession) SetScenario(v aSessionScenario) {
	if v == ESessionScenario.InProgress {
		s.Start()

	}
	if v == ESessionScenario.Finished {
		s.Stop()
	}
}

func (s *IPhotosSession) saveManifest() {
	resultPath := filepath.Join(PathSessionStorageSetting.GetValue(), s.getRelRootSessionPath(false), "manifest.json")
	ioutil.WriteFile(resultPath, []byte(helpers.ToJSONString(s)), 0644)
}

func (s *IPhotosSession) Start() {
	if s.Scenario == ESessionScenario.Stopping {
		s.Scenario = ESessionScenario.AbortStopping
		time.Sleep(2 * time.Second)
	}

	if s.watcher == nil {
		watcher, onCloseWatcher, err := pathwatcher.StartWatcher(PathCameraIncomeSetting.GetValue(), s.newSourceFileAdded)

		if err != nil {
			log.AddLog(err.Error())
			if watcher != nil {
				watcher.Close()
			}
			return
		}

		s.watcher = watcher
		s.onCloseWatcher = onCloseWatcher
		//s.adobeQueue = make(chan adobeBackend.IAdobeJOBRequest, 1000000)
		//s.waitGroup.Add(1)

		s.StartedAt = time.Now()

		log.AddLog("Начата новая сессия")

		// Run 1 worker to handle jobs.
		//go s.worker()

		s.Scenario = ESessionScenario.InProgress
	}

	var arr []uint64 = nil
	for _, v := range s.PresetsList.List {
		arr = append(arr, v.Id)
	}

	s.PresetsList.List = nil
	for _, v := range preset.GetProjectPresets() {
		if helpers.ContainsInt(arr, v.Id) {
			s.PresetsList.List = append(s.PresetsList.List, v)
		}
	}

	log.AddLog("Обновлены пресеты")

	s.Save()

}

func (s *IPhotosSession) Save() {
	log.Println(s)
	if s.Id == 0 {
		_, err := db.Model(s).Insert()
		if err != nil {
			panic(err)
		}
		sessionsList = append(sessionsList, s)
	} else {
		_, err := db.Model(s).WherePK().Update() //.Where("uid = ?0", s.ID).Update()
		if err != nil {
			panic(err)
		}
	}
}

func (s *IPhotosSession) getRelRootSessionPath(forFile bool) string {
	var fileSugar = ""
	var sessionSugar = ""

	if dotenv.GetEnv().Mode == "dev" {
		if forFile {
			fileSugar = "_n" + strconv.Itoa(int(s.incomeFileCount))
		}
		sessionSugar = "_s" + strconv.Itoa(int(s.Id))
	}
	return s.StartedAt.Format("2006-01-02") + sessionSugar + fileSugar + "_SelfPortrait_" + s.Person.LastName + "_" + s.Person.FirstName
}

func (s *IPhotosSession) GetAdobeQueueCount() int {
	return s.CountOfAdobeRequest
}

func (s *IPhotosSession) Stop() {
	go func() {

		s.Scenario = ESessionScenario.Stopping
		s.Save()

		//закрываем вотчер если с него перестали прилетать файлы (а значит их не появлялось)
		for int(time.Now().Sub(s.lastEventNewCameraFile).Seconds()) < 5 {
			time.Sleep(time.Second)
			if scene.IsUserClosedApp() {
				return
			}
		}

		if s.watcher != nil {
			s.onCloseWatcher()
			//s.watcher.Close()
			s.watcher = nil
			s.onCloseWatcher = nil
		}

		for s.GetAdobeQueueCount() > 0 {
			time.Sleep(time.Second)

			if s.Scenario == ESessionScenario.AbortStopping {
				s.Scenario = ESessionScenario.InProgress
				log.Println("Сессия возобновлена")
				return
			}

			if scene.IsUserClosedApp() {
				return
			}
		}

		//for s.GetAdobeQueueCount() > 0 {
		//
		//	time.Sleep(time.Second)
		//	if s.Scenario == ESessionScenario.AbortStopping {
		//		s.Scenario = ESessionScenario.InProgress
		//		log.Println("Сессия возобновлена")
		//		return
		//	}
		//
		//}

		s.Scenario = ESessionScenario.Finished
		//if s.adobeQueue != nil {
		//	close(s.adobeQueue)
		//	s.waitGroup.Wait()
		//}
		s.Save()
		s.saveManifest()
		go s.sendNotify()

		log.AddLog("Сессия остановлена")
	}()
}

func (s *IPhotosSession) sendNotify() {
	if !email.IsEmailValid(s.Person.Email, true) {
		return
	}
	if s.PublicUrl == "" {
		return
	}

	email.SendYandexDiskMailNotify(s.Person.Email, fmt.Sprintf("%s %s", s.Person.FirstName, s.Person.LastName), s.PublicUrl)

}
func (s *IPhotosSession) newSourceFileAdded(f string) {
	if scene.IsUserClosedApp() {
		return
	}

	if s.Scenario != ESessionScenario.InProgress {
		return
	}
	log.AddLog("Получен новый снимок")
	log.AddLog(f)

	fi, err := os.Stat(f)
	if err != nil {
		log.Println(err)
		return
	}
	switch mode := fi.Mode(); {
	case mode.IsDir():
		// do directory stuff
		log.Println("directory")
	case mode.IsRegular():
		// do file stuff
		log.Println("file #", fileNumber)
		fileNumber++

		resultPath := filepath.Join(PathOperationalSetting.GetValue(), s.getRelRootSessionPath(false), s.getRelRootSessionPath(true)+"_"+filepath.Base(f))
		//os.MkdirAll(filepath.Dir(resultPath), os.ModePerm)
		//moveFile(f, resultPath, true)
		s.lastEventNewCameraFile = time.Now()

		err := helpers.MoveFile(f, resultPath, true)
		if err != nil {
			time.Sleep(time.Second)
			s.newSourceFileAdded(f)
			return
		}
		log.Println("moved file")

		s.incomeFileCount++
		s.addToAdobeQueue(resultPath)
	}

}

func (s *IPhotosSession) convertToAdobeRequest(f string) adobeBackend.IAdobeJOBRequest {

	var presets []adobeBackend.IAdobeJOBPresetsRequest

	for _, v := range s.PresetsList.List {
		//PrecetFile string `json:"precetFile"`
		//SubFolder  string `json:"subFolder"`
		//SaveFormat string `json:"saveFormat"`
		//Big        bool   `json:"big"`
		//BigSize    int    `json:"bigSize"`
		//Small      bool   `json:"small"`
		//SmallSize  int    `json:"smallSize"`
		//JpgQuality int    `json:"jpgQuality"`
		//ActionName string `json:"actionName"`
		//ActionSet  string `json:"actionSet"`

		adobePreset := adobeBackend.IAdobeJOBPresetsRequest{
			PrecetFile: v.Filename,
			SubFolder:  v.SubPathName,
			SaveFormat: strings.ToLower(strings.Join(v.Formats, ",")),
			Big:        helpers.ItemExists(v.PresetSizes, preset.EPresetSize.Big),
			BigSize:    preset.PresetBigFileSizePixel.GetIntValue(),
			Small:      helpers.ItemExists(v.PresetSizes, preset.EPresetSize.Small),
			SmallSize:  preset.PresetSmallFileSizePixel.GetIntValue(),
			JpgQuality: 10,
			ActionName: v.ActionName,
			ActionSet:  v.ActionSet,
		}

		presets = append(presets, adobePreset)
	}

	adobeRequest := adobeBackend.IAdobeJOBRequest{
		SourceFile: f,
		SaveDir:    filepath.Dir(f),
		LogFolder:  filepath.Join(PathLogsSetting.GetValue(), "_adobeLogs"),
		LogFile:    filepath.Base(f) + ".log",
		Presets:    presets,
		OnDone: func(v *http.Response, adobeRequest *adobeBackend.IAdobeJOBRequest) {

			s.CountOfAdobeRequest--

			if s.PublicUrl == "" {
				s.PublicUrl = yandexCloud.ShareAndGetPublicURL(s.getRelRootSessionPath(false))
				s.Save()
			}

			//move original file from oper to result path
			//resultPath := filepath.Join(PathSessionStorageSetting.GetValue(), s.getRelRootSessionPath(false))
			//os.MkdirAll(resultPath, os.ModePerm)
			helpers.MoveFile(adobeRequest.SourceFile, filepath.Join(PathSessionStorageSetting.GetValue(), s.getRelRootSessionPath(false), filepath.Base(adobeRequest.SourceFile)), true)

			s.copyToCloudStoragePath(adobeRequest)
			log.Println("done moving file to cloud ... ")
			//moveToResultStoragePath()

			//uiUpdateAdobeQueueLabel(s)
		},
		OnError: func() {
			//log.Println(v)
			log.AddLog("error do request, retry ... ")
		},
	}
	return adobeRequest
}

func (s *IPhotosSession) copyToCloudStoragePath(adobeRequest *adobeBackend.IAdobeJOBRequest) {
	resultPath := filepath.Join(PathCloudSetting.GetValue(), s.getRelRootSessionPath(false))
	os.MkdirAll(resultPath, os.ModePerm)
	//@todo  копировать файлы
	//moveFile(f, resultPath)
	log.AddLog("Копируем результат адобе в облачную папку")
	err := helpers.CopyDir(filepath.Dir(adobeRequest.SourceFile), resultPath, []string{filepath.Dir(adobeRequest.SourceFile)},
		func(f string) {
			//файл результата скопировали в клауд директорию
			//на нужно
			helpers.MoveFile(f,
				strings.Replace(f, filepath.Dir(adobeRequest.SourceFile), filepath.Join(PathSessionStorageSetting.GetValue(), s.getRelRootSessionPath(false)), 1),
				true,
			)
		},
	)
	if err != nil {
		fmt.Println(err)
		log.AddLog("Копирование не удалось")
	}
}

//
//func (s *IPhotosSession) doAdobeRequest(adobeRequest *adobeBackend.IAdobeJOBRequest) {
//
//	copyToCloudStoragePath := func() {
//		resultPath := filepath.Join(PathCloudSetting.GetValue(), s.getRelRootSessionPath(false))
//		os.MkdirAll(resultPath, os.ModePerm)
//		//@todo  копировать файлы
//		//moveFile(f, resultPath)
//		log.AddLog("Копируем результат адобе в облачную папку")
//		err := helpers.CopyDir(filepath.Dir(adobeRequest.SourceFile), resultPath, []string{filepath.Dir(adobeRequest.SourceFile)},
//			func(f string) {
//				//файл результата скопировали в клауд директорию
//				//на нужно
//				helpers.MoveFile(f,
//					strings.Replace(f, filepath.Dir(adobeRequest.SourceFile), filepath.Join(PathSessionStorageSetting.GetValue(), s.getRelRootSessionPath(false)), 1),
//					true,
//				)
//			},
//		)
//		if err != nil {
//			fmt.Println(err)
//			log.AddLog("Копирование не удалось")
//		}
//	}
//
//
//
//	adobeRequest.JobOnAdobe(func(v *http.Response) {
//		if s.PublicUrl == "" {
//			s.PublicUrl = yandexCloud.ShareAndGetPublicURL(s.getRelRootSessionPath(false))
//			s.Save()
//		}
//
//		//move original file from oper to result path
//		//resultPath := filepath.Join(PathSessionStorageSetting.GetValue(), s.getRelRootSessionPath(false))
//		//os.MkdirAll(resultPath, os.ModePerm)
//		helpers.MoveFile(adobeRequest.SourceFile, filepath.Join(PathSessionStorageSetting.GetValue(), s.getRelRootSessionPath(false), filepath.Base(adobeRequest.SourceFile)), true)
//
//		copyToCloudStoragePath()
//		//moveToResultStoragePath()
//
//		//uiUpdateAdobeQueueLabel(s)
//	}, func() {
//		log.AddLog("error do request, retry ... ")
//	})
//
//}

// worker processes jobs.
//func (s *IPhotosSession) worker() {
//	defer s.waitGroup.Done()
//	log.Println("Worker is waiting for jobs")
//	for {
//		select {
//
//		case job, ok := <-s.adobeQueue:
//			if !ok {
//				return
//			}
//			log.Println("Worker picked job", job)
//			s.DoAdobeRequest(&job)
//
//		default:
//			if scene.IsUserClosedApp() {
//				break
//			}
//		}
//	}
//}

func (s *IPhotosSession) addToAdobeQueue(f string) {
	adobeBackend.Start()

	log.AddLog("Добавляем в очередь на запрос на адобне")
	log.AddLog(f)

	adobeBackend.AdobeQueue <- s.convertToAdobeRequest(f)
	s.CountOfAdobeRequest++
	//uiUpdateAdobeQueueLabel(s)
}

//func moveFile(from string, to string, createDir bool) bool {
//	if createDir {
//		os.MkdirAll(filepath.Dir(to), os.ModePerm)
//	}
//	err := os.Rename(from, to)
//	if err != nil {
//		log.Println("error to rename file:", err)
//		return false
//	}
//	return true
//}

func getLastRow() *IPhotosSession {
	//session := new(IPhotosSession)
	//err := db.Model(session).Order("uid DESC").Limit(1).Select()
	//if err != nil {
	//	log.AddLog(err.Error())
	//	panic(err)
	//}
	//log.Println("last session")
	//log.Println(session, session.Person)
	//
	//sessionsList = append(sessionsList, session)
	list := getSessionList()
	//var session *IPhotosSession
	if list == nil {
		session := new(IPhotosSession)
		sessionsList = append(sessionsList, session)
		return session
	}

	return list[0]
}

func getSessionList() []*IPhotosSession {
	if sessionsList == nil {
		err := db.Model(&sessionsList).Order("uid DESC").Select()

		if err != nil {
			panic(err)
		}

		for _, v := range sessionsList {
			if v.Scenario == ESessionScenario.InProgress || v.Scenario == ESessionScenario.Stopping {
				v.resumeSession()

				if v.Scenario == ESessionScenario.Stopping {
					v.Stop()
				}

			}
		}

	}
	return sessionsList
}

func (s *IPhotosSession) resumeSession() {

	src := filepath.Join(PathOperationalSetting.GetValue(), s.getRelRootSessionPath(false))
	entries, err := ioutil.ReadDir(src)
	if err != nil {
		return
	}

	for _, entry := range entries {
		if entry.IsDir() {

		} else {

			// Skip symlinks.
			if entry.Mode()&os.ModeSymlink != 0 {
				continue
			}

			//os.MkdirAll(filepath.Dir(resultPath), os.ModePerm)
			//moveFile(f, resultPath, true)
			srcPath := filepath.Join(src, entry.Name())

			s.lastEventNewCameraFile = time.Now()
			s.incomeFileCount++
			s.addToAdobeQueue(srcPath)

		}
	}
}
