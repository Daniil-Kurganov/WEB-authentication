import { Component } from "@angular/core";
import { FormsModule } from "@angular/forms";
import { HttpClient, HttpClientModule, HttpHeaders } from "@angular/common/http";
     
@Component({
    selector: "web-authentication",
    standalone: true,
    imports: [FormsModule, HttpClientModule],
    template: `<div>
      <button (click)="sendMessage()">Send Message</button>
    </div>`
})
export class AppComponent {
  numberSessionRounds: number;

  constructor(private http: HttpClient){}

  sendMessage() {
    const data = {
      p: 12,
      g: 14,
      y: 13
    };
    const headers = new HttpHeaders({ 'Content-Type': 'application/json' });
    this.http.post(`schnorr_auth/init_session`, data, { headers }).subscribe({
      next: (response: number) => {
        console.log(response);
        this.numberSessionRounds = response;
      },
      error: error => console.log(error)
    });
  }
}