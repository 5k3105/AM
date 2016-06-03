package main

import (
	"io/ioutil"
	"os"
	"path"
	"strings"

	"local/AM/controlbox"
	"local/AM/gfxcanvas"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

var (
	statusbar     *widgets.QStatusBar
	FsModel       *widgets.QFileSystemModel
	FsTreeView    *widgets.QTreeView
	StackedLayout *widgets.QStackedLayout
)

var (
	textedit *widgets.QTextEdit
	table    *widgets.QTableWidget
	canvas   *gfxcanvas.Canvas
)

// note: watch for -wireless connect- network connection crashing program when set in SLEEP mode

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	// Main Window
	var window = widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("---") 
	//window.SetSurfaceType() opengl

	// Main Widget
	mw := widgets.NewQWidget(nil, 0)

	// Statusbar
	statusbar = widgets.NewQStatusBar(window)
	window.SetStatusBar(statusbar)

	// Stacked Layout
	StackedLayout = widgets.NewQStackedLayout2(mw)

	// Create Canvas, Table, Textedit
	canvas = gfxcanvas.NewCanvas(statusbar)

	table = widgets.NewQTableWidget2(12, 3, nil) //NewTable()
	textedit = widgets.NewQTextEdit(nil)

	StackedLayout.AddWidget(canvas)   // 0
	StackedLayout.AddWidget(table)    // 1
	StackedLayout.AddWidget(textedit) // 2

	// mods
	StackedLayout.SetCurrentIndex(0)

	// Dock Toolbox
	var dControlbox = widgets.NewQDockWidget("Controlbox", window, 0)
	window.AddDockWidget(core.Qt__LeftDockWidgetArea, dControlbox)

	var cb = controlbox.NewControlbox(canvas.QWidget, statusbar)
	dControlbox.SetWidget(cb.QWidget)
	canvas.SetControlbox(cb)

	// Dock FileTree
	var dFsTree = widgets.NewQDockWidget("File System", window, 0)
	window.AddDockWidget(core.Qt__LeftDockWidgetArea, dFsTree)
	var FsTree = NewFsTree()
	dFsTree.SetWidget(FsTree)

	statusbar.ShowMessage(core.QCoreApplication_ApplicationDirPath(), 0)

	// Set Central Widget
	window.SetCentralWidget(mw)

	// Run App
	widgets.QApplication_SetStyle2("fusion")
	window.ShowMaximized()
	widgets.QApplication_Exec()
}

func NewFsTree() *widgets.QTreeView {

	var drive string = "O:\\" 

	FsModel = widgets.NewQFileSystemModel(nil)
	FsModel.SetRootPath(drive)

	FsTreeView = widgets.NewQTreeView(nil)
	FsTreeView.SetModel(FsModel)
	FsTreeView.SetRootIndex(FsModel.Index2(drive, 0))
	FsTreeView.HideColumn(1)
	FsTreeView.HideColumn(2)
	FsTreeView.HideColumn(3)

	//FsTreeView.ConnectCurrentChanged(treeViewCurrentChanged)

	return FsTreeView
}

func treeViewCurrentChanged(current *core.QModelIndex, previous *core.QModelIndex) {
	FileName := FsModel.FileName(current)
	FilePath := FsModel.FilePath(current)

	FsTreeView.ScrollTo(current, 0) // 0 = EnsureVisible, 3 = PositionAtCenter

	//fp := FilePath + FileName
	statusbar.ShowMessage(FilePath, 0)

	s := strings.ToLower(path.Ext(FileName))
	f := strings.TrimLeft(s, ".")

	switch s {
	case ".png", ".jpg", ".tiff":
		canvas.ShowPic(FilePath, f)
		//StackedLayout.SetCurrentIndex(0)

	case ".txt", ".csv", ".kml", ".xml", ".ini", ".log", ".msg":

		var content, err = ioutil.ReadFile(FilePath)
		if err != nil {
			panic(err)
		}

		var font = gui.NewQFont2("courier", 12, 5, false)

		textedit.SetText(string(content))
		textedit.SetFont(font)
		//StackedLayout.SetCurrentIndex(2)

	default:
		table.InsertRow(table.RowCount())
		//StackedLayout.SetCurrentIndex(1)
		canvas.ClearScene()
	}

}

func NewTable() *widgets.QTableWidget {
	table := widgets.NewQTableWidget2(12, 3, nil)

	return table
}

//-------------------------------------------------------------------------Notes
//SelectionModel *core.QItemSelectionModel
//FsTreeView.SetSelectionModel(core.QAbstractItemView.SingleSelection)

//FsTreeView.SelectionModel().CurrentChanged
//FsTreeView.SelectionModel().ConnectCurrentChanged(TreeViewCurrentChanged)

//	var s string = current.Row()
//	statusbar.ShowMessage(s, 0)

//	var Item = widgets.NewQTreeWidgetItem()

//	Item := FsTreeView.CurrentItem()
//	var s string = Item.Text(0)

//	var font = gui.NewQFont2("Meiryo", 20, 2, false)
//	scene.AddText("Hello 世界", font)

//	var color = gui.NewQColor2(255, 0, 0, 255)
//	var pen = gui.NewQPen3(color)

//	scene.AddLine2(0, scene.Height(), scene.Width(), scene.Height(), pen)

//	view.SetScene(scene)
