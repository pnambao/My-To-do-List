import { Component, signal } from '@angular/core';
import { TaskService } from '../../services/task';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-tasks',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './tasks.html',
  styleUrl: './tasks.css'
})
export class Tasks {

  tasks = signal<any[]>([]);

  constructor(private taskService: TaskService) {
    console.log("Tasks component created");

    this.taskService.getTasks().subscribe({
      next: (tasks: any) => {
        console.log("Received:", tasks);
        this.tasks.set(tasks);
      },
      error: (err) => {
        console.error("Failed to load tasks:", err);
      }
    });
  }
}