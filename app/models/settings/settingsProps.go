package settings

type APropertyType = string

type lPropertyType struct {
	BOOL APropertyType
	STR  APropertyType
	INT  APropertyType
}

// Enum for public use
var EPropertyType = &lPropertyType{
	BOOL: "BOOL",
	STR:  "STR",
	INT:  "INT",
}

type APropertyKey = string

const (
	CLOUD_YANDEX_LOGIN APropertyKey = "CLOUD_YANDEX_LOGIN"
	CLOUD_YANDEX_PASS  APropertyKey = "CLOUD_YANDEX_PASS"
	CLOUD_YANDEX_TOKEN APropertyKey = "CLOUD_YANDEX_TOKEN"

	ADOBE_EXTENSION_API_URL APropertyKey = "ADOBE_EXTENSION_API_URL"

	PATHS_CAMERA_INCOME       APropertyKey = "PATHS_CAMERA_INCOME"
	PATHS_OPERATIONAL_STORAGE APropertyKey = "PATHS_OPERATIONAL_STORAGE"
	PATHS_SESSIONS_STORAGE    APropertyKey = "PATHS_SESSIONS_STORAGE"
	PATHS_PRESETS             APropertyKey = "PATHS_PRESETS"
	PATHS_LOGS                APropertyKey = "PATHS_LOGS"
	PATHS_CLOUD               APropertyKey = "PATHS_CLOUD"

	BIG_FILE_SIZE_PERCENT   APropertyKey = "BIG_FILE_SIZE_PERCENT"
	SMALL_FILE_SIZE_PERCENT APropertyKey = "SMALL_FILE_SIZE_PERCENT"

	BIG_FILE_SIZE_PIXEL   APropertyKey = "BIG_FILE_SIZE_PIXEL"
	SMALL_FILE_SIZE_PIXEL APropertyKey = "SMALL_FILE_SIZE_PIXEL"

	TIMEOUT_ADOBE_SERVICE APropertyKey = "TIMEOUT_ADOBE_SERVICE"

//	TIMEOUT_SESSION_ATEND  APropertyKey = "SMALL_FILE_SIZE_PERCENT"

)

type ISettings struct {
	tableName     struct{}      `pg:"settings"`
	ID            uint64        `pg:"uid,pk"`
	PropertyKey   APropertyKey  `pg:"prop_key"`
	PropertyType  APropertyType `pg:"prop_type"`
	PropertyValue string        `pg:"prop_value"`
	MinInt        int           `pg:"-"`
	MaxInt        int           `pg:"-"`
}
