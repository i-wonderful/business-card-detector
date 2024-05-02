# Card detector

```
sudo docker-compose up --build
```

Запуск UI на http://127.0.0.1:8080

Rest API. POST http://127.0.0.1:8080/detect 

Пример
```
curl --location 'localhost:8080/detect' \
--form 'image=@"/home/olga/projects/card_detector_imgs/37.JPG"'
```