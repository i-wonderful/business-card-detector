В папке lib находятся модели PaddleOCR.

./lib/paddleocr/onnx/ - модели формата onnx.
./lib/paddleocr/whl/ - модели формата paddlepaddle.

В данный момент модели paddlepadlle невозможно запустить на виртуаке, 
потому что в серверном процессоре отсутствует поддержка avx команд. Это работает только на новых процах.
Проверить можно так: cat /proc/cpuinfo | grep -i avx.
Если ничего не вывелось, то avx нет и paddlepadlle работать не будет.

Поэтому используется onnx runtime для запуска на хостинге.

Проверить работу onnx 

```bash 
rapidocr_onnxruntime -img ../storage/5.JPG --rec_model_path ./onnx/en_PP-OCRv4_rec_infer.onnx
```
```bash 
rapidocr_onnxruntime -img ./storage/5.JPG --rec_model_path ./lib/paddleocr/onnx/en_PP-OCRv4_rec_infer.onnx
```

Проверить работу PaddleOCR на движке paddlepaddle (должен быть установлен на компе с avx)
```
paddleocr --image_dir /app/storage/<some_img_name> --lang=en  --show_log=False --use_angle_cls=True
```
```bash
paddleocr --image_dir https://marketplace.canva.com/EAFUXb9i_OM/1/0/1600w/canva-green-and-white-modern-business-card-rU-gq1vTReM.jpg --lang=en --use_angle_cls=true
```