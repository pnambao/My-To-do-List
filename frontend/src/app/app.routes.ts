import { Routes } from '@angular/router';
import { Login } from './components/Login/login';
import { Register } from './components/Register/register';
import { Tasks } from './components/Tasks/tasks';

export const routes: Routes = [
  {
    path: '',
    redirectTo: 'login',
    pathMatch: 'full'
  },

  {
    path: 'login',
    component: Login
  },

  {
    path: 'register',
    component: Register
  },

  {
    path: 'tasks',
    component: Tasks
  }
];