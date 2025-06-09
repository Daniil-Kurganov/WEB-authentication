import { Component } from "@angular/core";
import { FormsModule } from "@angular/forms";
import { HttpClient, HttpClientModule } from "@angular/common/http";
     
@Component({
    selector: "web-authentication",
    standalone: true,
    imports: [FormsModule, HttpClientModule],
    template: `<div>
      <button (click)="sendMessage()">Send Message</button>
    </div>`
})
export class AppComponent {
  constructor(
     private http: HttpClient
  ){}

  sendMessage() {
    this.http.post(`schnorr_auth/test`, null).subscribe({next: (response: any) => {
      console.log(response);
    },})
  }
}