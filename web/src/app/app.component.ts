import {Component, inject} from "@angular/core";
import {FormBuilder, Validators, FormsModule, ReactiveFormsModule} from '@angular/forms';
import {HttpClient, HttpClientModule, HttpHeaders} from "@angular/common/http";
import {MatCardModule} from '@angular/material/card';
import {MatInputModule} from '@angular/material/input';
import {MatButtonModule} from '@angular/material/button';
import {MatStepperModule} from '@angular/material/stepper';
import {MatFormFieldModule} from '@angular/material/form-field';
     
@Component({
    selector: "web-authentication",
    standalone: true,
    imports: [FormsModule, HttpClientModule, MatCardModule, MatInputModule, MatButtonModule, MatStepperModule, MatFormFieldModule, ReactiveFormsModule],
    templateUrl: "app.component.html"
})
export class AppComponent {
  p: number;
  g: number;
  y: number;
  x: number;
  numberSessionRounds = "-";
  numberDoneRounds = "-";
  numberSuccessRounds = "-";
  private _formBuilder = inject(FormBuilder);
  firstFormGroup = this._formBuilder.group({
    firstCtrl: ['', Validators.required],
  });
  secondFormGroup = this._formBuilder.group({
    secondCtrl: ['', Validators.required],
  });

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