import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';


export interface ShortenedURL{
    short_code: string;
    short_url: string;
    original_url: string;
    clicks: number;
    ID?: number;
}

@Injectable({
    providedIn: 'root'
})

export class UrlServive{
    private apiUrl = 'http://localhost:8080'
    
    constructor (private http: HttpClient) {}

    shortenUrl(url: string): Observable<ShortenedURL>{
        const body = {url: url};
        return this.http.post<ShortenedURL>(`${this.apiUrl}/shorten`, body);
}

getUrl(): Observable<ShortenedURL[]>{
    return this.http.get<ShortenedURL[]>(`${this.apiUrl}/list`);

}
}


