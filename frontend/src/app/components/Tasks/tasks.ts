import { Component, signal } from '@angular/core';
import { TaskService } from '../../services/task';
import { CommonModule } from '@angular/common';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { Router } from '@angular/router';

@Component({
  selector: 'app-tasks',
  standalone: true,
  imports: [CommonModule, ReactiveFormsModule],
  templateUrl: './tasks.html',
  styleUrl: './tasks.css'
})
export class Tasks {

  tasks = signal<any[]>([]);
  editingTaskId: number | null = null;
  username: string = localStorage.getItem('username') ?? 'User';

  createForm: FormGroup;

  constructor(private taskService: TaskService, private fb: FormBuilder, private router: Router) {
    this.createForm = this.fb.group({
      title: ['', Validators.required],
      date: ['', Validators.required]
    });

    this.loadTasks();
  }

  logout() {
    localStorage.removeItem('username');
    localStorage.removeItem('userId');
    this.router.navigate(['/login']);
  }
  userId: string = localStorage.getItem('userId') ?? '';

  loadTasks() {
    this.taskService.getTasks(this.userId).subscribe({
      next: (tasks: any) => this.tasks.set(tasks),
      error: (err) => console.error('Failed to load tasks:', err)
    });
  }

  createTask() {
    if (this.createForm.invalid) return;

    const formValue = this.createForm.value;
    const isoDate = `${formValue.date}T00:00:00Z`;

    if (this.editingTaskId) {
      const existing = this.tasks().find(t => t.id === this.editingTaskId);
      this.taskService.updateTask({
        id: this.editingTaskId,
        title: formValue.title,
        date: isoDate,
        completed: existing?.completed ?? false
      }).subscribe({
        next: () => {
          this.editingTaskId = null;
          this.createForm.reset();
          this.loadTasks();
        },
        error: (err) => console.error('Failed to update task:', err)
      });
    } else {
      this.taskService.createTask({
        user_id: Number(this.userId),
        title: formValue.title,
        date: isoDate
      }).subscribe({
        next: () => {
          this.createForm.reset();
          this.loadTasks();
        },
        error: (err) => console.error('Failed to create task:', err)
      });
    }
  }

  startEdit(task: any) {
    this.editingTaskId = task.id;
    this.createForm.setValue({
      title: task.title,
      date: task.date?.substring(0, 10) ?? ''
    });
  }

  cancelEdit() {
    this.editingTaskId = null;
    this.createForm.reset();
  }

  deleteTask(task: any) {
    this.taskService.deleteTask(task.id).subscribe({
      next: () => this.loadTasks(),
      error: (err) => console.error('Failed to delete task:', err)
    });
  }

  toggleCompleted(task: any) {
    this.taskService.updateTask({
      id: task.id,
      title: task.title,
      date: task.date,
      completed: !task.completed
    }).subscribe({
      next: () => this.loadTasks(),
      error: (err) => console.error('Failed to update task:', err)
    });
  }

}