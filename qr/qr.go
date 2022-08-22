package qr

import (
	"fmt"
	"image/png"
	"log"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/mordfustang21/gozbar"

	"gocv.io/x/gocv"
)

func CreateQR(name string) error {
	//QRコードを作成する
	if name == "" {
		return fmt.Errorf("name is empty")
	}

	qr, _ := qr.Encode(name, qr.M, qr.Auto)
	qr, _ = barcode.Scale(qr, 200, 200)

	file, _ := os.Create("qr.png")
	defer file.Close()

	png.Encode(file, qr)
	return nil
}

func QRScan() string {
	webcam, _ := gocv.VideoCaptureDevice(0)
	img := gocv.NewMat()
	var qrinfo string

	for {
		webcam.Read(&img)

		if img.Empty() {
			continue
		}

		qr, _ := img.ToImage()
		qrimg := gozbar.FromImage(qr)
		//ScanしたQRコードを画像に変換し，それを取り込む

		s := gozbar.NewScanner()
		err := s.SetConfig(gozbar.QRCODE, gozbar.CFG_ENABLE, 1)
		if err != nil {
			log.Fatal("error setting config", err)
		}
		//ScanするためのConfig

		res := s.Scan(qrimg)
		//Scan

		if res == nil { //Scanできた場合，結果を表示(resはerror handling!)
			qrimg.First().Each(func(r string) {
				//fmt.Println(r)
				qrinfo = r
			})
			webcam.Close()
			s.Destroy()
			return qrinfo
		}
	}
}
