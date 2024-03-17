package controllers

import (
	"net/http"
	"github.com/felipePinzon77/invent2_prueba/backend/db"

	"github.com/felipePinzon77/invent2_prueba/backend/models"
)

func Front(w http.ResponseWriter, r *http.Request) {
	conexionEsta, _ := db.ConexionDB()
	// Execute SELECT query
	registros, err := conexionEsta.Query("SELECT ProductoID, Nombre, Descripcion, Precio, UnidadMedida, ProveedorID, LicenciaCertificacionID FROM Productos")
	if err != nil {
		panic(err.Error())
	}
	defer registros.Close()

	// Store query results in a slice of 'southree' structs
	var arregloEmpresa []models.Producto
	for registros.Next() {
		var empresa models.Producto
		if err := registros.Scan(&empresa.ProductoID, &empresa.Nombre, &empresa.Descripcion, &empresa.Precio, &empresa.UnidadMedida, &empresa.ProveedorID, &empresa.LicenciaCertificacionID); err != nil {
			panic(err.Error())
		}
		arregloEmpresa = append(arregloEmpresa, empresa)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Delete(w http.ResponseWriter, r *http.Request) {

	conexionEsta, _ := db.ConexionDB()
	ProductoId := r.URL.Query().Get("id")

	borrarRegistros, err := conexionEsta.Prepare("DELETE FROM Productos WHERE ProductoID = ?")

	//Confirmar que no hayan errores
	if err != nil {
		panic(err.Error())
	}

	//Ejecturar la instruccion anterior (DELETE)
	borrarRegistros.Exec(ProductoId)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Update(w http.ResponseWriter, r *http.Request) {

	conexionEsta, _ := db.ConexionDB()

	if r.Method == "POST" {
		ProductoID := r.FormValue("ProductoID")
		Nombre := r.FormValue("Nombre")
		Descripcion := r.FormValue("Descripcion")
		Precio := r.FormValue("Precio")
		UnidadMedida := r.FormValue("UnidadMedida")
		ProveedorID := r.FormValue("ProveedorID")
		LicenciaCertificacionID := r.FormValue("LicenciaCertificacionID")

		stmt, err := conexionEsta.Prepare("UPDATE Productos SET Nombre = ?, Descripcion = ?, Precio = ?, UnidadMedida = ?, ProveedorID = ?, LicenciaCertificacionID = ? WHERE ProductoID = ?")
		if err != nil {
			http.Error(w, "Error preparing SQL statement", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		// Execute SQL statement
		_, err = stmt.Exec(Nombre, Descripcion, Precio, UnidadMedida, ProveedorID, LicenciaCertificacionID, ProductoID)
		if err != nil {
			http.Error(w, "Error updating product in database", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func Post(w http.ResponseWriter, r *http.Request) {

	conexionEsta, _ := db.ConexionDB()

	if r.Method == "POST" {
		// Parse form values
		ProductoID := r.FormValue("ProductoID")
		Nombre := r.FormValue("Nombre")
		Descripcion := r.FormValue("Descripcion")
		Precio := r.FormValue("Precio")
		UnidadMedida := r.FormValue("UnidadMedida")
		ProveedorID := r.FormValue("ProveedorID")
		LicenciaCertificacionID := r.FormValue("LicenciaCertificacionID")

		// Prepare SQL statement
		stmt, err := conexionEsta.Prepare("INSERT INTO Productos (ProductoID, Nombre, Descripcion, Precio, UnidadMedida, ProveedorID, LicenciaCertificacionID) VALUES (?, ?, ?, ?, ?, ?, ?)")
		if err != nil {
			http.Error(w, "Error preparing SQL statement", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		// Execute SQL statement
		_, err = stmt.Exec(ProductoID, Nombre, Descripcion, Precio, UnidadMedida, ProveedorID, LicenciaCertificacionID)
		if err != nil {
			http.Error(w, "Error updating product in database", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}