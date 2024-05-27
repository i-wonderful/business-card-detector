from rapidocr_onnxruntime import RapidOCR
import sys

model_rec_path = './lib/paddleocr/onnx/en_PP-OCRv4_rec_infer.onnx'
model_det_path = './lib/paddleocr/onnx/en_PP-OCRv3_det_infer.onnx'
engine = RapidOCR(rec_model_path=model_rec_path, det_model_path=model_det_path)

img_path = sys.argv[1]

result, elapse = engine(img_path)

for idx in range(len(result)):
    res = result[idx]
    print(res)


