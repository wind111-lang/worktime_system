package susys

import (
	"github.com/therecipe/qt/widgets"
	_"github.com/therecipe/qt/core"
)

func NewWindow(btn *widgets.QPushButton) {
	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(500, 500)
	window.SetWindowTitle("New window test")
	//window.SetAttribute(core.Qt__WA_DeleteOnClose, true)

	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(widgets.NewQVBoxLayout())
	window.SetCentralWidget(widget)
	// ウインドウ作成処理

	button := widgets.NewQPushButton2("Close", nil)
	widget.Layout().AddWidget(button)

	button.ConnectClicked(func(bool) {
		btn.SetEnabled(true)
		window.Close()
		//押したらウィンドウが閉じられる
	})

	window.Show()
}
