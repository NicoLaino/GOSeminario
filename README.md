# Seminario GO

Nicolás Laino - TUDAI 2020

The idea of this Final Proyect is to have a CRUD collection.

**Endpoints Implemented:**

**LIST CARS:** GET http://localhost:8080/cars/  
**LIST SPECIFIC CAR :** GET http://localhost:8080/cars/%ID%  (ID for example 4)  
**CREATE NEW CAR:** POST http://localhost:8080/cars/   
(BODY FOR EXAMPLE {  
    "text": "Corolla",  
})
**UPDATE EXISTING CAR:** PATCH http://localhost:8080/cars/%ID%  (ID for example 4)  
(BODY FOR EXAMPLE {  
    "text": "Corolla",      
})  
**DELETE SPECIFIC CAR :** DELETE http://localhost:8080/cars/%ID%  (ID for example 4)  