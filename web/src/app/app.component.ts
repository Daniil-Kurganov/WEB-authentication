import {Component} from "@angular/core";
import {FormsModule} from "@angular/forms";
import {HttpClient, HttpClientModule, HttpHeaders} from "@angular/common/http";
import {MatCardModule} from '@angular/material/card';
import {MatInputModule} from '@angular/material/input';
import {MatButtonModule} from '@angular/material/button';
     
@Component({
    selector: "web-authentication",
    standalone: true,
    imports: [FormsModule, HttpClientModule, MatCardModule, MatInputModule, MatButtonModule],
    templateUrl: "app.component.html"
})
export class AppComponent {
  p: number;
  g: number;
  y: number;
  numberSessionRounds = "-";
  numberDoneRounds = "-";
  numberSuccessRounds = "-";

  constructor(private http: HttpClient){}

  sendMessage() {
    const data = {
      p: this.p,
      g: this.g,
      y: this.y
    };
    const headers = new HttpHeaders({ 'Content-Type': 'application/json' });
    this.http.post(`schnorr_auth/init_session`, data, { headers }).subscribe({
      next: (response: string) => {
        console.log(`Number rounds: ${response}`);
        this.numberSessionRounds = response;
      },
      error: error => console.log(error)
    });
  }
}