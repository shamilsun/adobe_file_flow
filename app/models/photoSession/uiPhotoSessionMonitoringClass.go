package photoSession

import (
	"fmt"
	"github.com/andlabs/ui"
	"log"
)
import "../scene"
import "../adobeBackend"
import "time"

var table *ui.Table
var tableModel *ui.TableModel

type IModelHandler struct {
	listOfSessions []*IPhotosSession
}

var modelHandler *IModelHandler
var updateTableTimeout int

func (mh *IModelHandler) CellValue(m *ui.TableModel, row, column int) ui.TableValue {

	session := mh.GetSessionByRowNum(row)
	if session == nil {
		return ui.TableString("")
	}

	//table.AppendTextColumn("#",
	//	0, ui.TableModelColumnNeverEditable, nil)
	//table.AppendTextColumn("Статус",
	//	1, ui.TableModelColumnNeverEditable, nil)
	//table.AppendTextColumn("ФИО",
	//	2, ui.TableModelColumnNeverEditable, nil)
	//table.AppendTextColumn("Дата начала",
	//	3, ui.TableModelColumnNeverEditable, nil)
	//table.AppendTextColumn("Дата завершения",
	//	4, ui.TableModelColumnNeverEditable, nil)
	//table.AppendTextColumn("Осталось обработать файлов",
	//	5, ui.TableModelColumnNeverEditable, nil)
	//table.AppendTextColumn("Прогноз окончания",
	//	6, ui.TableModelColumnNeverEditable, nil)

	if column == 0 {
		return ui.TableString(fmt.Sprintf("%d", session.Id))
	}
	if column == 1 {
		return ui.TableString(session.Scenario)
	}
	if column == 2 {
		return ui.TableString(session.Person.GetFIO())
	}

	if column == 3 {
		return ui.TableString(session.StartedAt.Format("2006-01-02 15:04:05"))
	}

	if column == 4 {
		return ui.TableString(session.StoppedAt.Format("2006-01-02 15:04:05"))
	}

	if column == 5 {
		return ui.TableString(fmt.Sprintf("%d", session.GetAdobeQueueCount()))
	}

	if column == 6 {
		seconds := adobeBackend.GetRequestTimeDelayInSeconds()
		if seconds == 0 || session.GetAdobeQueueCount() == 0 {
			return ui.TableString("-")
		}
		seconds = seconds * session.GetAdobeQueueCount()
		m, _ := time.ParseDuration(fmt.Sprintf("%ds", seconds))
		return ui.TableString(fmt.Sprintf("%.1f минут", m.Minutes()))
	}

	if column == 7 {
		//if session.watcher != nil {
		//	return ui.TableString("+")
		//}
	}

	if column == 8 {
		//if session.watcher != nil {
		//	return ui.TableString("+")
		//}
		//return ui.Table("-")
		return ui.TableString("Завершить")
	}

	return ui.TableString("-")

}

func (mh *IModelHandler) SetCellValue(m *ui.TableModel, row, column int, value ui.TableValue) {
	log.Println(m, row)
	session := mh.GetSessionByRowNum(row)
	if session != nil {
		session.SetScenario(ESessionScenario.Finished)
	}
}

func (mh *IModelHandler) GetSessionByRowNum(row int) *IPhotosSession {
	for i, v := range mh.listOfSessions {
		if i == row {
			return v
		}
	}
	return nil
}

func newModelHandler() *IModelHandler {
	if modelHandler == nil {
		modelHandler = new(IModelHandler)
	}

	return modelHandler
}

func (mh *IModelHandler) ColumnTypes(m *ui.TableModel) []ui.TableValue {
	return []ui.TableValue{}
}

func (mh *IModelHandler) NumRows(m *ui.TableModel) int {
	log.Println("draw lines: ", len(mh.listOfSessions))
	return len(mh.listOfSessions)
}

//var img [2]*ui.Image
func updateTable(listOfSessions []*IPhotosSession) {
	prevLength := len(modelHandler.listOfSessions)
	newLength := len(listOfSessions)

	modelHandler.listOfSessions = listOfSessions

	//было 0 стало 1,
	if newLength > prevLength {
		for i := prevLength; i < newLength; i++ {
			log.Println("insert row", i)
			tableModel.RowInserted(i)
		}
	}

	//было 1 стало 0, 0<1
	if newLength < prevLength {
		for i := prevLength; i > newLength; i-- {
			log.Println("delete row", i)
			//			tableModel.RowDeleted(i)
		}
	}

	for i := 0; i < newLength; i++ {
		//log.Println("changed row", i)
		tableModel.RowChanged(i)
	}

	//table.Hide()
	//table.Show()
	//log.Println("update tables by new list", prevLength, newLength)
	//table.Hide()
	//table.Show()
	//time.Sleep(5*time.Second)
	setUpdateTableTimeout(5)
}

func UIGetPhotoSessionMonitoring() *ui.Group {

	scene.SetSessionStatusPanel(ui.NewGroup("Мониторинг сессий"))
	scene.GetSessionStatusPanel().Show()

	vbox := ui.NewVerticalBox()

	mh := newModelHandler()
	tableModel = ui.NewTableModel(mh)

	table = ui.NewTable(&ui.TableParams{
		Model:                         tableModel,
		RowBackgroundColorModelColumn: 3,
	})

	table.AppendTextColumn("#",
		0, ui.TableModelColumnNeverEditable, nil)
	table.AppendTextColumn("Статус",
		1, ui.TableModelColumnNeverEditable, nil)
	table.AppendTextColumn("ФИО",
		2, ui.TableModelColumnNeverEditable, nil)
	table.AppendTextColumn("Дата начала",
		3, ui.TableModelColumnNeverEditable, nil)
	table.AppendTextColumn("Дата завершения",
		4, ui.TableModelColumnNeverEditable, nil)
	table.AppendTextColumn("Осталось обработать файлов",
		5, ui.TableModelColumnNeverEditable, nil)
	table.AppendTextColumn("Прогноз окончания",
		6, ui.TableModelColumnNeverEditable, nil)
	table.AppendTextColumn("Вотчинг входящих",
		7, ui.TableModelColumnNeverEditable, nil)
	table.AppendButtonColumn("Операция",
		8, ui.TableModelColumnAlwaysEditable)

	//table.AppendImageTextColumn("Column 2",
	//	5,
	//	1, ui.TableModelColumnNeverEditable, &ui.TableTextColumnOptionalParams{
	//		ColorModelColumn:		4,
	//	});
	//table.AppendTextColumn("Editable",
	//	2, ui.TableModelColumnAlwaysEditable, nil)
	//
	//table.AppendCheckboxColumn("Checkboxes",
	//	7, ui.TableModelColumnAlwaysEditable)
	//table.AppendButtonColumn("Buttons",
	//	6, ui.TableModelColumnAlwaysEditable)
	//
	//table.AppendProgressBarColumn("Progress Bar",
	//	8)
	btnUpdate := ui.NewButton("Обновить")
	btnUpdate.OnClicked(func(button *ui.Button) {
		setUpdateTableTimeout(0)
	})
	vbox.Append(btnUpdate, false)
	vbox.Append(table, true)
	scene.GetSessionStatusPanel().SetChild(vbox)

	return scene.GetSessionStatusPanel()
}

func setUpdateTableTimeout(v int) {
	updateTableTimeout = v
}

func getUpdateTableTimeout() int {
	return updateTableTimeout
}

func decUpdateTableTimeout() {
	updateTableTimeout--
}
