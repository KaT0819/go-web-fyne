package main

import (
	"log"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type MyEntry struct {
	widget.Entry
	entered func(e *MyEntry)
}

// オリジナルの入力フォームを作成
func NewMyEntry(f func(e *MyEntry)) *MyEntry {
	e := &MyEntry{}
	e.ExtendBaseWidget(e)
	e.entered = f
	return e
}

// キーダウンイベント
func (e *MyEntry) KeyDown(key *fyne.KeyEvent) {
	switch key.Name {
	case fyne.KeyReturn, fyne.KeyEnter:
		e.entered(e)
	default:
		e.Entry.KeyDown(key)
	}
}

func main() {
	// アプリケーションの作成
	a := app.New()

	// カスタムテーマを読み込む
	// a.Settings().SetTheme(myTheme{})
	// a.Settings().SetTheme(theme.DarkTheme())

	// ウィンドウの作成
	w := a.NewWindow("Hello")
	w.Resize(fyne.NewSize(400, 600))

	// ウィジェットの設定
	// ラベル
	l := widget.NewLabel("Hello Fyne!")
	l2 := widget.NewLabel("")
	l3 := widget.NewLabel("")

	// 入力フォーム
	e := widget.NewEntry()
	e.SetText("10")

	// チェックボックス
	c := widget.NewCheck("chekk", func(f bool) {
		if f {
			l2.SetText("Checed !")
		} else {
			l2.SetText("not Checked !")
		}
	})
	// チェックボックスの選択状態
	c.SetChecked(false)

	// ラジオボタン
	r := widget.NewRadioGroup(
		[]string{"Go", "Python", "PHP"},
		func(s string) {
			if s == "" {
				l3.SetText("Not Selected")
			} else {
				l3.SetText("Selected: " + s)
			}
		},
	)
	// ラジオボタンの洗濯
	r.SetSelected("Go")

	// スライダー
	s := widget.NewSlider(0.0, 100)
	sv := widget.NewLabel("slider")
	b := widget.NewButton("View Slider Value", func() {
		sv.SetText("Slider Value: " + strconv.Itoa(int(s.Value)))
	})

	// 選択リスト
	lv := widget.NewLabel("Select Listed Value")
	sl := widget.NewSelect([]string{
		"", "Apple", "Orange", "Melon",
	}, func(s string) {
		lv.SetText("Select Listed Value: " + s)
	})
	// リストの選択
	sl.SetSelected("Orange")

	// プログレスバー
	v := 0.
	p := widget.NewProgressBar()
	pb := widget.NewButton("Progress Increment", func() {
		v += 0.1
		if v > 1.0 {
			v = 0.
		}
		p.SetValue(v)
	})

	// グループ（Deprecatedになっており、V2では消えている。Cardの使用を促しているが機能的に違うような気がする。。。）
	card := widget.NewCard("Card", "subtitle", container.NewVBox(p, pb))
	// card.SetContent(pb)

	// ツールバー
	tb := widget.NewToolbar(
		widget.NewToolbarAction(theme.HomeIcon(), func() {
			l.SetText("Select Home Icon")
		}),
		widget.NewToolbarAction(theme.InfoIcon(), func() {
			l.SetText("Select Info Icon")
		}),
	)
	// メニューバー
	mm := fyne.NewMainMenu(
		fyne.NewMenu("File",
			fyne.NewMenuItem("New", func() {
				l.SetText("select 'New' menu item.")
			}),
			fyne.NewMenuItem("Quit", func() {
				a.Quit()
			}),
		),
	)
	w.SetMainMenu(mm)

	// 1タブ　ボーダーレイアウト
	// box1 := container.NewVBox()
	widget.NewButton("Click", nil)

	// フォーム
	fLbl := widget.NewLabel("")
	fName := widget.NewEntry()
	fPass := widget.NewPasswordEntry()
	fMsg := NewMyEntry(func(e *MyEntry) {
		s := e.Text
		e.SetText("")
		fLbl.SetText("you type '" + s + "' .")
	})

	b1 := widget.NewButton("top", nil)
	// b2 := widget.NewButton("bottom", nil)
	b3 := widget.NewButton("left", nil)
	b4 := widget.NewButton("right", nil)
	box2 := container.NewVBox(
		l,
		fLbl,
		widget.NewForm(
			widget.NewFormItem("Name", fName),
			widget.NewFormItem("Pass", fPass),
		),
		widget.NewButton("Form Click", func() {
			fLbl.SetText(fName.Text + " : " + fPass.Text)
		}),
		fMsg,
	)
	// ボーダーレイアウト　ツールバーを下部に設定
	blayout := container.New(
		layout.NewBorderLayout(b1, tb, b3, b4),
		b1, tb, b3, b4, box2,
	)

	de := widget.NewEntry()
	// 2タブ グリッドレイアウト
	glayout := container.New(
		layout.NewGridLayout(3),
		l,
		widget.NewLabel("sample サンプルアプリケーション"),
		e,
		// ボタン
		widget.NewButton("Click", func() {
			n, _ := strconv.Atoi(e.Text)
			log.Println(n)
			// カスタムダイアログ　複数ダイアログを定義した場合書き出しと逆順になる。
			dialog.ShowCustomConfirm(
				"Enter message.",
				"OK",
				"Cancel", de, func(b bool) {
					if b {
						l.SetText("typed: '" + de.Text + "' .")
					} else {
						l.SetText("no message")
					}
				}, w)
			// 確認ダイアログ
			dialog.ShowConfirm("Confirm", "Please check 'YES'!", func(b bool) {
				if b {
					l.SetText("OK, thank you")
				} else {
					l.SetText("oh ...")
				}
			}, w)
			// インフォダイアログ
			dialog.ShowInformation("Alert", "This is sample alert!", w)
			time.Sleep(1 * time.Second) // 1秒待つ
			lv.SetText("Total: " + strconv.Itoa(total(n)))
		}),
	)
	// グリッドレイアウト（縦順）
	grlayout := container.New(
		layout.NewGridLayoutWithRows(3),
		// グループ（カード）
		card,
		c, l2, layout.NewSpacer(),
		r, l3, layout.NewSpacer(),
		sv, s, b, layout.NewSpacer(),
		lv, sl,
	)
	// グリッドレイアウト（幅指定）
	gflayout := container.New(
		layout.NewGridWrapLayout(fyne.NewSize(200, 50)),
		l,
		widget.NewLabel("sample サンプルアプリケーション"),
		e,
		// ボタン
		widget.NewButton("Click", func() {
			n, _ := strconv.Atoi(e.Text)
			log.Println(n)
			lv.SetText("Total: " + strconv.Itoa(total(n)))
		}),
	)

	// コンテンツの作成
	w.SetContent(
		container.NewAppTabs(
			container.NewTabItem("First", blayout),
			container.NewTabItem("grid Layout", glayout),
			container.NewTabItem("grid Layout Row", grlayout),
			container.NewTabItem("wrap grid Layout", gflayout),
			// スクロールコンテナ
			container.NewTabItem("scloll", container.NewVScroll(
				container.NewVBox(
					widget.NewLabel("sample scrolll"),
					widget.NewLabel("sample scrolll"),
					widget.NewLabel("sample scrolll"),
					widget.NewLabel("sample scrolll"),
					widget.NewLabel("sample scrolll"),
					widget.NewLabel("sample scrolll"),
					widget.NewLabel("sample scrolll"),
					widget.NewLabel("sample scrolll"),
					widget.NewLabel("sample scrolll"),
					widget.NewLabel("sample scrolll"),
					widget.NewLabel("sample scrolll"),
					widget.NewLabel("sample scrolll"),
					widget.NewLabel("sample scrolll"),
					widget.NewLabel("sample scrolll"),
					widget.NewLabel("sample scrolll"),
					widget.NewLabel("sample scrolll"),
					widget.NewLabel("sample scrolll"),
					widget.NewLabel("sample scrolll"),
				),
			)),
		),
	)

	// ウィンドウの表示、実行
	w.ShowAndRun()

}

func total(n int) int {
	t := 0

	for i := 1; i <= n; i++ {
		t += i
	}

	return t
}
