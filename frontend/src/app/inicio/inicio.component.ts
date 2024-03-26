import { HttpClient } from '@angular/common/http';
import { Component,OnInit } from '@angular/core';
import { error } from 'console';

@Component({
  selector: 'app-inicio',
  standalone: true,
  imports: [],
  templateUrl: './inicio.component.html',
  styleUrl: './inicio.component.css'
})
export class InicioComponent implements OnInit {

  users:any[]=[];
  constructor(private http: HttpClient){

  }
  ngOnInit(): void{
    this.getAllusers();

  }
  getAllusers(){
    this.http.get('').subscribe((res: any)=>{
      this.users = res.data;
    }, error=>{
      alert("Error from API ")
    })
  }
}
