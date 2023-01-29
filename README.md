## **Search-Engine-Task**

<br>

[![codecov](https://codecov.io/gh/Anvesh965/Search-Engine/branch/master/graph/badge.svg?token=TG5HLM0M31)](https://codecov.io/gh/Anvesh965/Search-Engine)

---

### **Step 1:** Clone the repository :

```
git clone https://github.com/Anvesh965/Search-Engine.git "Search Engine"
```

### **Step 2: Running on Docker**

(Make sure your docker engine is running)

```
cd "Search Engine"
make up
```

### **Step 3 :** Testing API

- Use any API testing applications such as Postman, Insomnia or ThunderClinet (VS Code extension)

- All available Routes

  1. To check server status: **GET** Request

     ```
     http://localhost:4000/
     ```

  2. To see all webpages in MongoDB : **GET** Request

     ```
     http://localhost:4000/v1/allpages
     ```

  3. To store webpage in MongoDB : **POST** Request

     ```
     http://localhost:4000/v1/savepage
     ```

     Add this json to request body

     ```json
     {
       "title": "P1",
       "keywords": ["tesla", "ford", "wan"]
     }
     ```

  4. Query : **GET** Request

     ```
     http://localhost:4000/v1/querypages
     ```

     Add this json to request body

     ```json
     {
       "keywords": ["tesla", "ford", "wan"]
     }
     ```

  5. Swagger UI
     ```
     http://localhost:4000/docs/index.html
     ```

### **Step 4 :** To stop running containers

```
 make down
```
