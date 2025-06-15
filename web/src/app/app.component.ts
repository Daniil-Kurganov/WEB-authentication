import {Component, inject} from "@angular/core";
import {FormBuilder, Validators, FormsModule, ReactiveFormsModule} from '@angular/forms';
import {HttpClient, HttpClientModule, HttpHeaders} from "@angular/common/http";
import {MatCardModule} from '@angular/material/card';
import {MatInputModule} from '@angular/material/input';
import {MatButtonModule} from '@angular/material/button';
import {MatStepperModule} from '@angular/material/stepper';
import {MatFormFieldModule} from '@angular/material/form-field';
import {NgIf} from '@angular/common';
     
@Component({
    selector: "web-authentication",
    standalone: true,
    imports: [
      FormsModule,
      HttpClientModule,
      MatCardModule,
      MatInputModule,
      MatButtonModule,
      MatStepperModule,
      MatFormFieldModule,
      ReactiveFormsModule,
      NgIf
    ],
    templateUrl: "app.component.html",
    styleUrl: "app.component.css"
})
export class AppComponent {
  showResult = false;
  p: number;
  g: number;
  y: number;
  x: number;
  e: number;
  s: number;
  numberSessionRounds = 0;
  numberDoneRounds = 0;
  numberSuccessRounds = 0;
  roundResult = "Определение...";
  roundFinalButtonText = "Начать новый раунд";
  sessionResult: string;
  private _formBuilder = inject(FormBuilder);
  firstFormGroup = this._formBuilder.group({
    firstCtrl: ['', Validators.required],
  });
  secondFormGroup = this._formBuilder.group({
    secondCtrl: ['', Validators.required],
  });

  constructor(private http: HttpClient){}

  initSession() {
    this.showResult = false;
    const data = {
      p: this.p,
      g: this.g,
      y: this.y
    };
    const headers = new HttpHeaders({ 'Content-Type': 'application/json' });
    this.http.post(`schnorr_auth/init_session`, data, { headers }).subscribe({
      next: (response: number) => {
        console.log(`Number rounds: ${response}`);
        this.numberSessionRounds = response;
      },
      error: error => console.log(error)
    });
  }

  firstStep() {
    const data = {
      x: this.x
    }
    const headers = new HttpHeaders({ 'Content-Type': 'application/json' });
    this.http.post(`schnorr_auth/first_step`, data, { headers }).subscribe({
      next: (response: number) => {
        console.log(`Round E parameter: ${response}`);
        this.e = response;
      },
      error: error => console.log(error)
    });
  }

  secondStep() {
    const data = {
      s: this.s
    }
    const headers = new HttpHeaders({ 'Content-Type': 'application/json' });
    this.http.post(`schnorr_auth/final_step`, data, { headers }).subscribe({
      next: (response: boolean) => {
        console.log(`Round result: ${response}`);
        this.numberDoneRounds++;
        if (response) {
          this.numberSuccessRounds++;
          this.roundResult = "Успешно"
        } else {
          this.roundResult = "Не успешно"
        }
      },
      error: error => console.log(error)
    })
    if ((this.numberDoneRounds + 1) === this.numberSessionRounds) {
      this.roundFinalButtonText = "Просмотреть результаты сессии";
    }
  }

  finalRoundStep() {
    if (this.numberDoneRounds === this.numberSessionRounds) {
      this.showResult = true;
      this.http.get(`schnorr_auth/session_result`).subscribe({
        next:(response: boolean) => {
          if (response) {
            this.sessionResult = "Аутентификация пройдена успешно"
          } else {
            this.sessionResult = "Аутентификация провалена"
          }
        },
        error: error => console.log(error)
      });
    }
  }
}