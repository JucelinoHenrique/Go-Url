import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common'; 
import { FormsModule } from '@angular/forms';

import { UrlServive, ShortenedURL } from '../../core/url.service';

@Component({
  selector: 'app-shortener',
  standalone: true,
  imports: [CommonModule, FormsModule],
  templateUrl: './shortener.html',
  styleUrls: ['./shortener.css']
})

export class ShortenerComponent implements OnInit{
  longUrl: string = '';
  shortenedUrl: ShortenedURL | null = null;
  listUrls: ShortenedURL[] = [];

  constructor(private urlService: UrlServive) {}
  ngOnInit(): void {
    this.loadUrls();
  }
  shortenUrl(): void {
    if (!this.longUrl) return;

    this.urlService.shortenUrl(this.longUrl).subscribe({
      next: (data: ShortenedURL) => {
        this.shortenedUrl = data;
        this.longUrl = '';
        this.loadUrls();
      },
      error: (error) => {
        console.error('Error shortening URL:', error);
      }
    });
  }
  loadUrls(): void {
    this.urlService.getUrl().subscribe({
      next: (data: ShortenedURL[]) => {
        this.listUrls = data;
      },
      error: (error) => {
        console.error('Error fetching URLs:', error);
      }
    });
  }
}