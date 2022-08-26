package susys

import (
	"github.com/therecipe/qt/core"
	_ "github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func NewWindow(btn *widgets.QPushButton) {
	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(320, 320)
	window.SetWindowFlags(core.Qt__CustomizeWindowHint | core.Qt__WindowTitleHint)
	window.SetWindowTitle("スタッフ管理画面")
	//window.SetAttribute(core.Qt__WA_DeleteOnClose, true)

	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(widgets.NewQVBoxLayout())
	window.SetCentralWidget(widget)

	registerButton := widgets.NewQPushButton2("スタッフ登録", nil)
	updateButton := widgets.NewQPushButton2("スタッフ情報更新", nil)
	deleteButton := widgets.NewQPushButton2("スタッフ削除", nil)
	closeButton := widgets.NewQPushButton2("終了", nil)

	widget.Layout().AddWidget(registerButton)
	widget.Layout().AddWidget(updateButton)
	widget.Layout().AddWidget(deleteButton)
	widget.Layout().AddWidget(closeButton)

	registerButton.ConnectClicked(func(bool) {
		widgets.QMessageBox_Warning(nil, "スタッフ登録", "準備中", widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	})

	updateButton.ConnectClicked(func(bool) {
		widgets.QMessageBox_Warning(nil, "スタッフ情報更新", "準備中", widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	})

	deleteButton.ConnectClicked(func(bool) {
		widgets.QMessageBox_Warning(nil, "スタッフ削除", "準備中", widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	})

	closeButton.ConnectClicked(func(bool) {
		btn.SetEnabled(true)
		window.Close()
	}) //ウインドウを閉じるときにボタンを有効化

	window.Show()
}
