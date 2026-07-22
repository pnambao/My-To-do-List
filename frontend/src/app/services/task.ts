import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";

@Injectable({
    providedIn: 'root'
  })
  export class TaskService {
  
    private api = 'http://localhost:8080';
  
    constructor(private http: HttpClient) {}
  
    getTasks() {
      return this.http.get(`${this.api}/tasks`);
    }
  
    createTask(task: any) {
      return this.http.post(`${this.api}/tasks`, task);
    }
  
    updateTask(task: any) {
      return this.http.put(`${this.api}/tasks`, task);
    }
  
    deleteTask(id: number) {
      return this.http.delete(`${this.api}/tasks?id=${id}`);
    }
  
  }