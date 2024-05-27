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