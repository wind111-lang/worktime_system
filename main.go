package main

import (
	"encoding/json"
	"fmt"
	"os"

	"worktime_system/qr"
	"worktime_system/susys"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age,string"`
}

func main() {

	// QWidgetsをスタートさせるために必要な処理
	app := widgets.NewQApplication(len(os.Args), os.Args)

	var nameInput *widgets.QLineEdit
	var ageInput *widgets.QLineEdit

	// ウィンドウ生成
	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(360, 420)
	window.SetWindowTitle("勤怠管理システム")
	window.SetWindowFlags(core.Qt__CustomizeWindowHint | core.Qt__WindowCloseButtonHint | core.Qt__WindowMinimizeButtonHint)
	//最大化ボタン無効化
	//window.SetAttribute(core.Qt__WA_DeleteOnClose, true)

	// ウィジェットを作成し，NewQVBoxLayoutを使ってレイアウトを作成
	// ウィンドウの中央にウィジェットを配置
	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(widgets.NewQVBoxLayout())
	window.SetCentralWidget(widget)

	// 文字を表示させるラベルの生成，文字入力フォームの作成，表の作成
	label := widgets.NewQLabel2("now:", nil, 0)
	widget.Layout().AddWidget(label)

	table := widgets.NewQTableWidget(nil)
	table.SetRowCount(50)
	table.SetColumnCount(3)
	table.SetHorizontalHeaderLabels([]string{"名前", "従業員コード", "出社時刻"})
	table.SetEditTriggers(widgets.QAbstractItemView__NoEditTriggers)
	table.AddScrollBarWidget(widgets.NewQScrollBar2(core.Qt__Vertical, nil), core.Qt__AlignRight)

	//table.SetItem(0, 0, widgets.NewQTableWidgetItem2("test", 0))

	widget.Layout().AddWidget(table)

	nameInput = widgets.NewQLineEdit(nil)
	nameInput.SetPlaceholderText("name")
	widget.Layout().AddWidget(nameInput)

	ageInput = widgets.NewQLineEdit(nil)
	ageInput.SetPlaceholderText("age")
	widget.Layout().AddWidget(ageInput)

	// ボタン生成，QRコード読み取り処理を行う
	button := widgets.NewQPushButton2("QR Scan", nil)
	button2 := widgets.NewQPushButton2("Create QR", nil)
	button3 := widgets.NewQPushButton2("Register", nil)
	button4 := widgets.NewQPushButton2("New Window", nil)

	widget.Layout().AddWidget(button)
	widget.Layout().AddWidget(button2)
	widget.Layout().AddWidget(button3)
	widget.Layout().AddWidget(button4)

	button3.SetEnabled(false)

	button.ConnectClicked(func(bool) {
		res := qr.QRScan()

		var person Person
		err := json.Unmarshal([]byte(res), &person)//jsonを構造体に変換
		if err != nil {
			widgets.QMessageBox_Critical(nil, "Error", err.Error(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
		} else {
			label.SetText("now: " + person.Name)
			fmt.Println(person)
		}
		//label.SetText("now:" + res)
	})

	// inputに入力された文字からQRコードを生成
	button2.ConnectClicked(func(bool) {
		err := qr.CreateQR(nameInput.Text(), ageInput.Text())
		if err != nil {
			widgets.QMessageBox_Critical(nil, "Error", err.Error(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
		} else {
			label.SetText("now:Registered")
			button3.SetEnabled(true)
		}
	})

	button3.ConnectClicked(func(bool) {
		widgets.QMessageBox_Information(nil, "Information", "Clicked", widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
		button3.SetEnabled(false)
		//スイッチ切り替えテスト
	})

	button4.ConnectClicked(func(bool) {
		susys.NewWindow(button4)
		button4.SetEnabled(false)
		//NewWindow表示テスト
	})

	// ウィンドウ表示
	window.Show()

	// Qtのループを開始，app.Exit()が呼ばれるかユーザによって終了されるまで継続
	app.Exec()
}
