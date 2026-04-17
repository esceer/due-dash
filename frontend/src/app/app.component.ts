import { Component, ViewEncapsulation } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { SideBarComponent } from "./core/components/sidebar/sidebar.component";
import { HeaderComponent } from './core/components/header/header.component';
import { TaskListComponent } from './features/tasks/components/task-list/task-list.component';

@Component({
  selector: 'dd-root',
  imports: [RouterOutlet, HeaderComponent, SideBarComponent, TaskListComponent],
  templateUrl: './app.component.html',
  styleUrl: './app.component.css',
  encapsulation: ViewEncapsulation.None
})
export class AppComponent {
  title = 'frontend';
}
