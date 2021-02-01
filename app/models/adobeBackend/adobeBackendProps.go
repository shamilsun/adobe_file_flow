package adobeBackend

import (
	"../settings"
	"net/http"
	"sync"
)

var AdobeQueue chan IAdobeJOBRequest
var waitGroup sync.WaitGroup

type IAdobeJOBRequest struct {
	SourceFile string                                                 `json:"sourceFile"`
	SaveDir    string                                                 `json:"saveDir"`
	LogFolder  string                                                 `json:"logFolder"`
	LogFile    string                                                 `json:"logFile"`
	Presets    []IAdobeJOBPresetsRequest                              `json:"presets"`
	OnDone     func(v *http.Response, adobeRequest *IAdobeJOBRequest) `json:"-"`
	OnError    func()                                                 `json:"-"`
}

type IAdobeJOBPresetsRequest struct {
	PrecetFile string `json:"precetFile"`
	SubFolder  string `json:"subFolder"`
	SaveFormat string `json:"saveFormat"`
	Big        bool   `json:"big"`
	BigSize    int    `json:"bigSize"`
	Small      bool   `json:"small"`
	SmallSize  int    `json:"smallSize"`
	JpgQuality int    `json:"jpgQuality"`
	ActionName string `json:"actionName"`
	ActionSet  string `json:"actionSet"`
}

var AdobeExtensionApiUrlSetting = settings.ISettings{
	PropertyType: settings.EPropertyType.STR,
	PropertyKey:  settings.ADOBE_EXTENSION_API_URL,
}
