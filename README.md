# YachtSearch

нужен менеджер зависимостей golang/dep
(brew install dep)

dep init 

в docker-compose заполнить
NAUSYS_USER
NAUSYS_PASSWORD
      
docker-compose up -d --build

http://127.0.0.1:8080/update?method=full - обновление базы компаний яхт и их моделей

http://127.0.0.1:8080/search?query=al - поиск по префиксу названия яхты или ее модели