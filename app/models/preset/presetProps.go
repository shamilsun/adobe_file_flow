package preset

import (
	"../imageFormat"
	"../settings"
)

type IPreset struct {
	tableName   struct{}                   `pg:"presets"`
	Id          uint64                     `pg:"uid,pk"`
	Filename    string                     `pg:"filename"`
	SubPathName string                     `pg:"sub_path_name"`
	Formats     []imageFormat.AImageFormat `pg:"formats"`
	PresetSizes []APresetSize              `pg:"sizes"`
	ActionName  string                     `pg:"action_name"`
	ActionSet   string                     `pg:"action_set"`
}

type IPresetList struct {
	List []*IPreset
}

var l IPreset = IPreset{Formats: []imageFormat.AImageFormat{
	imageFormat.EImageFormat.JPG,
}}

var PresetBigFileSizePercent settings.ISettings = settings.ISettings{
	PropertyType: settings.EPropertyType.INT,
	PropertyKey:  settings.BIG_FILE_SIZE_PERCENT,
	MinInt:       0,
	MaxInt:       100,
}

var PresetSmallFileSizePercent settings.ISettings = settings.ISettings{
	PropertyType: settings.EPropertyType.INT,
	PropertyKey:  settings.SMALL_FILE_SIZE_PERCENT,
	MinInt:       0,
	MaxInt:       100,
}

var PresetBigFileSizePixel settings.ISettings = settings.ISettings{
	PropertyType: settings.EPropertyType.INT,
	PropertyKey:  settings.BIG_FILE_SIZE_PIXEL,
	MinInt:       300,
	MaxInt:       15000,
}

var PresetSmallFileSizePixel settings.ISettings = settings.ISettings{
	PropertyType: settings.EPropertyType.INT,
	PropertyKey:  settings.SMALL_FILE_SIZE_PIXEL,
	MinInt:       300,
	MaxInt:       15000,
}
