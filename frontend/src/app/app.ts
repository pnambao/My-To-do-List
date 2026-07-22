import { Component, signal } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { TaskService } from './services/task';

@Component({
  selector: 'app-root',
  imports: [RouterOutlet],
  templateUrl: './app.html',
  styleUrl: './app.css'
})
export class App {
  protected readonly title = signal('frontend');

  constructor(private taskService: TaskService) {
    this.taskService.getTasks().subscribe({

      next: (tasks:any) => {
        console.log(tasks);
      },
  
      error: (err:any) => {
        console.error(err);
      }
  
    });  
  }
}
