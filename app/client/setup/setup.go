// 12 august 2018

// build OMIT

package setup

import "C"
import (
	"../../models/adobeBackend"
	"../../models/helpers"
	"../../models/photoSession"
	"../../models/preset"
	"../../models/yandexCloud"
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func UISettingsPage() ui.Control {

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	vbox.Append(uiSettingsPathsGroup(), false)
	vbox.Append(uiSettingsAdobeGroup(), false)
	vbox.Append(uiSettingsCloudGroup(), false)
	vbox.Append(uiSettingsPresetGroup(), false)

	return vbox
}

func uiSettingsPathsGroup() ui.Control {
	group := ui.NewGroup("Рабочие папки")
	group.SetMargined(true)
	group.SetChild(ui.NewNonWrappingMultilineEntry())
	entryFormPaths := ui.NewForm()

	entryFormPaths.Append("Источник (папка с исходниками камеры)", helpers.UIGetInputFieldForSetting(photoSession.PathCameraIncomeSetting), false)

	entryFormPaths.Append("Операционная папка (быстрый io для adobe)", helpers.UIGetInputFieldForSetting(photoSession.PathOperationalSetting), false)

	entryFormPaths.Append("Результирующее хранилище (архив папка)", helpers.UIGetInputFieldForSetting(photoSession.PathSessionStorageSetting), false)
	entryFormPaths.Append("Облачное хранилище (папка (корень))", helpers.UIGetInputFieldForSetting(photoSession.PathCloudSetting), false)

	entryFormPaths.Append("Срок хранения (дней)", ui.NewSlider(3, 90), false)
	entryFormPaths.Append("Папка с пресетами", helpers.UIGetInputFieldForSetting(photoSession.PathSessionStorageSetting), false)
	entryFormPaths.Append("Папка логов", helpers.UIGetInputFieldForSetting(photoSession.PathLogsSetting), false)

	entryFormPaths.SetPadded(true)
	group.SetChild(entryFormPaths)
	return group
}

func uiSettingsPresetGroup() ui.Control {
	group := ui.NewGroup("Настройки пресетов")
	group.SetMargined(true)
	group.SetChild(ui.NewNonWrappingMultilineEntry())
	entryFormPaths := ui.NewForm()

	entryFormPaths.Append("Размер BIG файла, px", helpers.UIGetSpinFieldForSetting(preset.PresetBigFileSizePixel), false)
	entryFormPaths.Append("Размер SMALL файла, px", helpers.UIGetSpinFieldForSetting(preset.PresetSmallFileSizePixel), false)

	entryFormPaths.SetPadded(true)
	group.SetChild(entryFormPaths)
	return group
}

func uiSettingsAdobeGroup() ui.Control {
	group := ui.NewGroup("Настройки стыка с ADOBE")
	group.SetMargined(true)
	group.SetChild(ui.NewNonWrappingMultilineEntry())
	entryFormPaths := ui.NewForm()

	//extensionApiEntry := ui.NewEntry()
	//entryFormPaths.Append("Ссылка на API", extensionApiEntry, false)
	//extensionApiEntry.SetText(adobeBackend.AdobeExtensionApiUrlSetting.GetValue())
	//extensionApiEntry.OnChanged(func(entry *ui.Entry) {
	//	adobeBackend.AdobeExtensionApiUrlSetting.UpdateValue(entry.Text())
	//})

	entryFormPaths.Append("Ссылка на API", helpers.UIGetInputFieldForSetting(adobeBackend.AdobeExtensionApiUrlSetting), false)

	//passwordEntry := ui.NewEntry()
	//entryFormPaths.Append("Пароль", passwordEntry, false)
	//passwordEntry.SetText(yandexCloud.YandexCloudPasswordSetting.GetValue())
	//passwordEntry.OnChanged(func(entry *ui.Entry) {
	//	yandexCloud.YandexCloudPasswordSetting.UpdateValue(entry.Text())
	//})

	//	entryFormPaths.Append("Токен", ui.NewEntry(), false)
	entryFormPaths.SetPadded(true)
	group.SetChild(entryFormPaths)
	return group
}

func uiSettingsCloudGroup() ui.Control {
	group := ui.NewGroup("Настройки облака (яндекс.диск)")
	group.SetMargined(true)
	group.SetChild(ui.NewNonWrappingMultilineEntry())
	entryFormPaths := ui.NewForm()

	entryFormPaths.Append("Токен", helpers.UIGetInputFieldForSetting(yandexCloud.YandexCloudTokenSetting), false)

	//loginEntry := ui.NewEntry()
	//entryFormPaths.Append("Логин", loginEntry, false)
	//loginEntry.SetText(yandexCloud.YandexCloudLoginSetting.GetValue())
	//loginEntry.OnChanged(func(entry *ui.Entry) {
	//	yandexCloud.YandexCloudLoginSetting.UpdateValue(entry.Text())
	//})
	//entryFormPaths.Append("Логин", helpers.UIGetInputFieldForSetting(yandexCloud.YandexCloudLoginSetting), false)

	//passwordEntry := ui.NewEntry()
	//entryFormPaths.Append("Пароль", passwordEntry, false)
	//passwordEntry.SetText(yandexCloud.YandexCloudPasswordSetting.GetValue())
	//passwordEntry.OnChanged(func(entry *ui.Entry) {
	//	yandexCloud.YandexCloudPasswordSetting.UpdateValue(entry.Text())
	//})
	//entryFormPaths.Append("Пароль", helpers.UIGetInputFieldForSetting(yandexCloud.YandexCloudPasswordSetting), false)

	//	entryFormPaths.Append("Токен", ui.NewEntry(), false)
	entryFormPaths.SetPadded(true)
	group.SetChild(entryFormPaths)
	return group
}
