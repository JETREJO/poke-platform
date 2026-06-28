package pokemon

/*
 * ------------    ¿ POR QUÉ PUNTEROS ?    --------------
 *
 * Se usan punteros (*int, *string) en algunos campos ya que
 * estas columnas (en la base de datos) pueden aceptar valores nulos,
 * y como aquí se tienen que definir tipos de datos, es necesario
 * usar punteros ya que estos pueden ser NIL en nuestro código y que
 * se mapee como NULL correctamente en el PostgreSQL.
 */

/*
 * ------------    ¿ JSON:"ID" ?    --------------
 *
 * Esto se llama "STRUCT TAG". Le dice al paquete que
 * convierta el Struct a JSON usando ese nombre de propiedad.
 */

type Pokemon struct {
	ID              int     `json:"id"`
	Name            string  `json:"name"`
	PrimaryTypeID   int     `json:"primaryTypeId"`
	SecondaryTypeID *int    `json:"secondaryTypeId"`
	GenerationID    int     `json:"generationId"`
	Sprite          *string `json:"sprite"`
	Shiny           bool    `json:"shiny"`
}
