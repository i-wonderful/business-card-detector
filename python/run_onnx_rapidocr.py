import sys

from rapidocr_onnxruntime import RapidOCR

#
# Detector PaddleOCR onnx
# Arguments:
#   - path to image for recognize
#   - path to detection pnnx model
#   - path to recognition pnnx model
# example:
# model_rec_path = './lib/paddleocr/onnx/en_PP-OCRv4_rec_infer.onnx'
# model_det_path = './lib/paddleocr/onnx/en_PP-OCRv3_det_infer.onnx'

img_path = sys.argv[1]
model_det_path = sys.argv[2]
model_rec_path = sys.argv[3]

rec_image_shape=[3, 30, 320] # [3, 48, 320]

engine = RapidOCR(rec_model_path=model_rec_path, det_model_path=model_det_path,
                  rec_image_shape=rec_image_shape,
                  # det_score_mode='slow',
                  # det_box_thresh=0.2,
                  # rec_batch_num=10,
                  #det_limit_side_len=200,
                  )

box_thresh = 0.4
unclip_ratio = 1.4 #1.5
text_score = 0.88

result, elapse = engine(img_path,
                        det_box_thresh=box_thresh, unclip_ratio=unclip_ratio, text_score=text_score,
                        rec_image_shape=rec_image_shape,
                        #det_donot_use_dilation=True,
                        #drop_score=0.4,
                        use_cls=False,
                        # drop_score=0.3,
                        #crop_word_box=True
                        )

for idx in range(len(result)):
    res = result[idx]
    print(res)
