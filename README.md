# YachtSearch

Нужен менеджер зависимостей golang/dep

    brew install dep
    dep init 

в docker-compose заполнить
NAUSYS_USER
NAUSYS_PASSWORD

Сборка и запуск контейнеров (nginx+frontend, postgresql, search-service, update-service)
     
     docker-compose up -d --build

Обновление базы компаний яхт и их моделей

    http://127.0.0.1:8080/update?method=full

Форнтенд доступен на 

    http://127.0.0.1:8080
    
Фронтенд использует запросы к поисковому api:

    http://127.0.0.1:8080/search/{query} - подгрузка через ajax подсказок для поисковой формы
    http://127.0.0.1:8080/info/model/{id} - получение информации по яхтам если была выбрана определенная модель
    http://127.0.0.1:8080/info/builder/{id} - получение информации по яхтам если был выбран определенный производитель
    http://127.0.0.1:8080/info/name/{query} - получение информации через поиск по кнопке
    