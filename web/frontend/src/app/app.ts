import { Component, signal } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { CommonModule } from '@angular/common';

import { ShortenerComponent } from './components/shortener/shortener';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [CommonModule, RouterOutlet, ShortenerComponent],  
  template: `
  <main>
  <app-shortener></app-shortener>
  </main>
  `,  
  styleUrls: ['./app.css']
})
export class App{
  title = 'Go-Url Frontend';

}
