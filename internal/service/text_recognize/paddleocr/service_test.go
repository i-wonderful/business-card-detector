package paddleocr

//const BASE_IMG_PATH = "/home/olga/projects/card_detector_imgs"
//
//func TestTextRecognizeService_RecognizeAll(t *testing.T) {
//
//	tests := []struct {
//		name string
//		path string
//		want []model.DetectWorld
//	}{
//		{
//			"16.JPG",
//			BASE_IMG_PATH + "/16.JPG",
//			[]model.DetectWorld{
//				{Text: "HiO"},
//				{Text: "0"},
//				{Text: "shubham.dhamija@deepdivemedia.in"},
//				{Text: "live:.cid.e53090522ec2bf11"},
//				{Text: "Shubham Dhamija"},
//				{Text: "Telegram@dshubham26"},
//				{Text: "WhatsApp:9034901070"},
//				{Text: "Strategy Head"},
//				{Text: "Skype"},
//				{Text: "Eml"},
//			},
//		},
//	}
//
//	// todo
//	service, err := NewService(true, "run.py", "det.onnx", "rec.onnx", "./tmp")
//	assert.NoError(t, err)
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//
//			got, _ := service.RecognizeAll(tt.path)
//
//			for i, world := range tt.want {
//				assert.Equal(t, world.Text, got[i].Text)
//			}
//		})
//	}
//}
