from rapidocr_onnxruntime import RapidOCR
import sys

#
# Detector PaddleOCR onnx
# Arguments:
#   - path to image for recognize
#   - path to detection pnnx model
#   - path to recognition pnnx model
# example:
#model_rec_path = './lib/paddleocr/onnx/en_PP-OCRv4_rec_infer.onnx'
#model_det_path = './lib/paddleocr/onnx/en_PP-OCRv3_det_infer.onnx'

img_path = sys.argv[1]
model_det_path = sys.argv[2]
model_rec_path = sys.argv[3]

engine = RapidOCR(rec_model_path=model_rec_path, det_model_path=model_det_path)

result, elapse = engine(img_path)

for idx in range(len(result)):
    res = result[idx]
    print(res)


