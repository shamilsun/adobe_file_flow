package photoSession

import (
	"../person"
	"../preset"
	"../settings"
	"github.com/fsnotify/fsnotify"
	"time"
)

type IPhotosSession struct {
	tableName              struct{}           `pg:"photo_sessions"`
	Id                     uint64             `pg:"uid,pk"`
	StartedAt              time.Time          `pg:"started_at"`
	StoppedAt              time.Time          `pg:"stopped_at"`
	Person                 person.IPerson     `pg:"person"`
	PresetsList            preset.IPresetList `pg:"presets"`
	Scenario               aSessionScenario   `pg:"scenario"`
	PublicUrl              string             `pg:"public_url"`
	watcher                *fsnotify.Watcher  `pg:"-"`
	CountOfAdobeRequest    int                `pg:"-"`
	cameraFiles            []string           `pg:"-"`
	lastEventNewCameraFile time.Time          `pg:"-"`
	incomeFileCount        uint64             `pg:"-"`

	onCloseWatcher func() `pg:"-"`
}

type aSessionScenario = string

type lSessionScenario struct {
	Pending       aSessionScenario
	InProgress    aSessionScenario
	Stopping      aSessionScenario
	Finished      aSessionScenario
	AbortStopping aSessionScenario
}

// Enum for public use
var ESessionScenario = &lSessionScenario{
	Pending:       "pending",
	InProgress:    "in_progress",
	Stopping:      "stopping",
	Finished:      "finished",
	AbortStopping: "abort_stopping",
}

var PathCameraIncomeSetting settings.ISettings = settings.ISettings{
	PropertyType: settings.EPropertyType.STR,
	PropertyKey:  settings.PATHS_CAMERA_INCOME,
}

var PathOperationalSetting settings.ISettings = settings.ISettings{
	PropertyType: settings.EPropertyType.STR,
	PropertyKey:  settings.PATHS_OPERATIONAL_STORAGE,
}

var PathSessionStorageSetting settings.ISettings = settings.ISettings{
	PropertyType: settings.EPropertyType.STR,
	PropertyKey:  settings.PATHS_SESSIONS_STORAGE,
}

var PathPresetsSetting settings.ISettings = settings.ISettings{
	PropertyType: settings.EPropertyType.STR,
	PropertyKey:  settings.PATHS_PRESETS,
}

var PathLogsSetting settings.ISettings = settings.ISettings{
	PropertyType: settings.EPropertyType.STR,
	PropertyKey:  settings.PATHS_LOGS,
}

var PathCloudSetting settings.ISettings = settings.ISettings{
	PropertyType: settings.EPropertyType.STR,
	PropertyKey:  settings.PATHS_CLOUD,
}
