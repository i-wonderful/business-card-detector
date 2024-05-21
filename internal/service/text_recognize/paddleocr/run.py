from paddleocr import PaddleOCR, draw_ocr
import sys

ocr = PaddleOCR(use_angle_cls=True, lang='en', show_log=False) # need to run only once to download and load model into memory
img_path = sys.argv[1]  #'/home/olga/projects/card_detector_imgs/16.JPG'
result = ocr.ocr(img_path, cls=True) #det=False detection rec=False means no recognition
for idx in range(len(result)):
    res = result[idx]
    for line in res:
        # print(line[1][0])
        print(line)


# draw result
# from PIL import Image
# result = result[0]
# image = Image.open(img_path).convert('RGB')
# boxes = [line[0] for line in result]
# txts = [line[1][0] for line in result]
# scores = [line[1][1] for line in result]
# im_show = draw_ocr(image, boxes, txts, scores, font_path='./Ubuntu-Th.ttf')   #
#
# im_show = Image.fromarray(im_show)
# im_show.save('result.jpg')