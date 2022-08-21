package main

import (
	"os"

	"qr/qr"

	"github.com/therecipe/qt/widgets"
)

func main() {

	// QWidgetsをスタートさせるために必要な処理
	app := widgets.NewQApplication(len(os.Args), os.Args)

	// ウィンドウ生成
	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(500, 500)
	window.SetWindowTitle("QR in GoQt")

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
	button.ConnectClicked(func(bool) {
		res := qr.QRScan()
		label.SetText("now: " + res)
	})
	widget.Layout().AddWidget(button)

	// inputに入力された文字からQRコードを生成
	button2 := widgets.NewQPushButton2("Register", nil)
	button2.ConnectClicked(func(bool) {
		err := qr.CreateQR(input.ToPlainText())
		if err != nil {
			widgets.QMessageBox_Critical(nil, "Error", err.Error(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
		} else {
			label.SetText("now:Registered" + input.ToPlainText())
		}
	})
	widget.Layout().AddWidget(button2)

	// ウィンドウ表示
	window.Show()

	// Qtのループを開始，app.Exit()が呼ばれるかユーザによって終了されるまで継続
	app.Exec()
}
