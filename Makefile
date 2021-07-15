run:
	# gin --all -i run main.go
	FYNE_FONT=C:\Windows\Fonts\meiryo.ttc || air

mod:
	go mod vendor

build:
	go build -mod=vendor

init:
	go get fyne.io/fyne/v2
	go get fyne.io/fyne/v2/cmd/fyne
	go mod tidy
	go mod vendor

# 日本語対応
setup:
	fyne bundle "C:\Windows\Fonts\meiryo.ttc" > bundle.go 
