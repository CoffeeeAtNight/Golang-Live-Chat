import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable()
export class ConfigService {
  constructor(private http: HttpClient) { }
  configUrl = 'http://localhost:8080/api/messages';

  getConfig() {
    return this.http.get<Config>(this.configUrl);
  }
}