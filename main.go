package main

import (
	"os"

	"worktime_system/qr"

	"github.com/therecipe/qt/widgets"
)

func main() {

	// QWidgetsをスタートさせるために必要な処理
	app := widgets.NewQApplication(len(os.Args), os.Args)

	// ウィンドウ生成
	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(500, 500)
	window.SetWindowTitle("Worktime System")

	// ウィジェットを作成し，NewQVBoxLayoutを使ってレイアウトを作成
	// ウィンドウの中央にウィジェットを配置
	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(widgets.NewQVBoxLayout())
	window.SetCentralWidget(widget)

	// 文字を表示させるラベルの生成，文字入力フォームの作成
	label := widgets.NewQLabel2("now:", nil, 0)
	widget.Layout().AddWidget(label)
	input := widgets.NewQPlainTextEdit2("", nil)
	widget.Layout().AddWidget(input)

	// ボタン生成，QRコード読み取り処理を行う
	button := widgets.NewQPushButton2("QR Scan", nil)
	button2 := widgets.NewQPushButton2("Create QR", nil)
	button3 := widgets.NewQPushButton2("Register", nil)

	widget.Layout().AddWidget(button)
	widget.Layout().AddWidget(button2)
	widget.Layout().AddWidget(button3)

	button3.SetEnabled(false)

	button.ConnectClicked(func(bool) {
		res := qr.QRScan()
		label.SetText("now: " + res)
	})

	// inputに入力された文字からQRコードを生成
	button2.ConnectClicked(func(bool) {
		err := qr.CreateQR(input.ToPlainText())
		if err != nil {
			widgets.QMessageBox_Critical(nil, "Error", err.Error(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
		} else {
			label.SetText("now:Registered" + input.ToPlainText())
			button3.SetEnabled(true)
		}
	})

	button3.ConnectClicked(func(bool) {
		widgets.QMessageBox_Information(nil, "Information", "Clicked", widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
		button3.SetEnabled(false)
	})

	// ウィンドウ表示
	window.Show()

	// Qtのループを開始，app.Exit()が呼ばれるかユーザによって終了されるまで継続
	app.Exec()
}
