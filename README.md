## **Search-Engine-Task**
---

### **Step 1:** Clone the repository :

```
git clone https://github.com/Anvesh965/Search-Engine.git "Search Engine"
```
### **Step 2: Running on Docker** 

(Make sure your docker engine is running)
```
cd "Search Engine"

docker-compose up --build
```

### **Step 3 :** Testing API

- Use any API testing applications such as Postman, Insomnia or ThunderClinet (VS Code extension)

- All available Routes

  1. To see all webpages in MongoDB : **GET** request

     `http://localhost:4000/v1/allpages`

  2. To store webpage in MongoDB : **POST** Request

     `http://localhost:4000/v1/savepage`

     Add this json to request body

     ```json
     {
        "title": "P1",
        "keywords": ["tesla", "ford", "wan"]
     }
     ```

  3. Query : **GET** Request

     `http://localhost:4000/v1/querypages`

     Add this json to request body

     ```json
     {
        "keywords": ["tesla", "ford", "wan"]
     }
     ```

### **Step 4 :** To stop running containers

```
 docker-compose down
```
