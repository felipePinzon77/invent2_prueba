// import { Injectable } from '@angular/core';

// @Injectable({
//   providedIn: 'root'
// })
// export class AuthService {

//   constructor() { }
// }
import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  constructor() { }

  login(email: string, password: string): Observable<any> {
    // Aquí iría tu lógica para verificar las credenciales, por ejemplo, una petición HTTP.
    // Simulamos una autenticación exitosa:
    if (email === 'user@example.com' && password === 'password') {
      return of({success: true, token: 'fake-jwt-token'});
    } else {
      return of({success: false});
    }
  }
}
