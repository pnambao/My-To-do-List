# My-To-do-List
This is a simple to-do list app created as part of a technical assessment.<br>
A user can add items to the list, remove items, edit an item, and view all items on the list at a glance.<br>
Each user's data is kept separate through authentication and database relationships.<br>

## Tech Stack

#### Frontend
Angular: I chose angular because I love how structured it is when organizing components and the overall skeleton of the project. It is also the framework I am most familiar with.

#### Backend
Golang: I chose Golang because it is a new language I started learning, and I thought this would be a nice project to use as a personal test and also to learn more in the process of building the project. Outside of personal reasons, I read that it is fast, lightweight, and well suited for creating REST APIs.

#### Database
PostgreSQL: I went with Postgres because it’s super reliable and handles all the connections between users and tasks without breaking anything.

## Assumtions
I went in with the following assumptions:
- I expect users to register an account first before accessing the internal workings of the app.
- Once a user has an account, I expect them to login before accessing their dashboard.
- On the dashboard, I expect the user's tasks to be listed when they login.
- I expect each task to be linked to one user ID so that each task can only be owned by one user. 
- On the backend, I expect the system to check the user before allowing them to delete, create or edit.

## How To Run

#### Prerequisites
If you don't already have these, download them:
- Go
- Node.js
- Angular CLI
- PostgreSQL

#### Running the project
1. Clone the repository: git clone [repo url]
2. Set up PostgreSQL: Create a PostgreSQL database named `todo_app` and run the SQL script to create the required tables. You might need to update the database connection string in `backend/main.go` with your PostgreSQL credentials for it to connect.
3. Start the Go backend
```bash
cd backend
go mod tidy
go run .
```
The backend API will run at:
```
http://localhost:8080
```
4. Start the Angular frontend
```bash
cd frontend
npm install
ng serve
```
The application will be available at:

```
http://localhost:4200
``` 
5. Try the following things:
- Registering a user 
- logging in with the credentials
- creating, editing, or deleting tasks.

6. To run the unit test
```bash
cd backend
go test
```

## Reflection

I originally aimed to match the dashboard to the Figma wireframe I ideated, but time constraints meant I couldn't fully finish its styling. However, I was able to get the login and register pages very close to my original visual vision, and this is because both pages shared a similar structure. So when I completed the login component, I was able to reuse and adapt that structure for the register component.

Moving forward, I want to continue learning Golang, while continuing to polish my Node.js fundamentals. 

## Early ideation and wireframing done in Figma
<img width="800" alt="Frame 13 (1)" src="https://github.com/user-attachments/assets/30150be1-f487-4d40-b1b5-e5e1feab447f" />
<img width="500" alt="Login" src="https://github.com/user-attachments/assets/5cbc7fe7-f930-46e3-9be8-d8955fd244fd" />
<img width="500" alt="Welcome and registration" src="https://github.com/user-attachments/assets/54453cfe-2051-47c6-bf23-b5053688d5b8" />
<img width="500" alt="Dashboard (2)" src="https://github.com/user-attachments/assets/a3178386-8b38-47f8-8778-7651a87c6dde" />


