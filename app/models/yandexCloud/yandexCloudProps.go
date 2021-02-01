package yandexCloud

import "../settings"

var YandexCloudLoginSetting settings.ISettings = settings.ISettings{
	PropertyType: settings.EPropertyType.STR,
	PropertyKey:  settings.CLOUD_YANDEX_LOGIN,
}

var YandexCloudPasswordSetting settings.ISettings = settings.ISettings{
	PropertyType: settings.EPropertyType.STR,
	PropertyKey:  settings.CLOUD_YANDEX_PASS,
}

var YandexCloudTokenSetting settings.ISettings = settings.ISettings{
	PropertyType: settings.EPropertyType.STR,
	PropertyKey:  settings.CLOUD_YANDEX_TOKEN,
}
