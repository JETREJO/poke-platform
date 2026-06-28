package pokemon

// El "package" le indica que "Todo lo que contiene este archivo pertenece al paquete POKEMON"

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// 1. Simepre los asteriscos ( * ) significan "PUNTEROS"
// 2. No estamos copiando la coneción, estamos apuntando a esa misma.
// 3. Ls conexiones suelen manejarse con punteros.

type Repository struct {
	db *pgx.Conn
}

// ------  Esto es un constructor  -----
// 4. En Go los constructores no son como en Java. Aquí hay que definir una función.
// 5. En símbolo ampersand ( & ) devuelve la dirección de memoria de lo que indiquemos.
func NewRepository(db *pgx.Conn) *Repository {
	// Regra la DIRECCIÓN DE MEMORIA del PUNTERO de la conexión
	return &Repository{
		db: db,
	}
}

// Este "func (r *Repository)" es un RECIBIDOR (Receiver)
// - NO es un parámetro, NO es una herencia, NO es una clase.
// - El Receiver es la forma en que GO le "agrega métodos" a un tipo (Repository, en este caso)
//
// 1. La parte "(r *Repository)" quiere decir -> "Esta función pertenece al Tipo REPOSITORY"
// El simil sería a escribir:
//
//	class Repository {
//	    getAll() {
//	    }
//	}
//
// 2. La variable "r" en este contexto es equivalente al "this" en otros lenguajes.
// - En este caso es "r" por "Repository", pero puede nombrarse como se desee.
//
//  3. "GetAll()" Es el nombre de nuestra función y los paréntesis vacíos
//     representan que no recibe parámetros.
//
//  4. "([]Pokemon, error)" quiere decir que nuestra función va a retornar
//     dos valores, el primero es un Slice (arreglo dinámico) del tipo Pokemon
//     y el segundo es un valor de tipo 'error'.
func (r *Repository) GetAll() ([]Pokemon, error) {

	// 1. "query()" es una función propia de la librería PGX.
	// - Se usa para cuando se esperan varias líneas (filas) como resultado.
	//
	// 2. El "context.Background" digamos que es la información que acompaña la operación.
	// - En el backend muchas operaciones pueden cancelarse, tener timeout, expirar, etc,
	//   y el context es justamente toda esa información que "no se ve".
	//
	// 3. "defer" es una palabra reservada de Go.
	// - Quiere decir "Ejecuta esto cuando la función termine".

	rows, err := r.db.Query(
		context.Background(),
		`
		SELECT
			id,
			name,
			primary_type_id,
			secondary_type_id,
			generation_id,
			sprite,
			shiny
		FROM pokemon;
		`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// Aquí declaramos el Slice (la lista de pokemones)
	var pokemonsList []Pokemon

	// 1. ".Next()" es un método de un Slice.
	//    - Devuelve un boolean que indica si después de la fila/posición actual
	//      existe otra posición adelante.
	for rows.Next() {
		var pokemon Pokemon

		// 1. ":=" quiere decir
		//
		// 2. "Scan" toma la fila actual que nos regresó la consulta y asigna los
		//    valores (cada columna) a nuestro Struct automáticamente.
		//
		// 3. Se usa "&" para indicarle al Scan que "Escriba dentro de la variable"
		//    y no que solo la lea y la copie. Es como decirle "No copies el valor,
		//    Ve directamente al lugar donde está guardado y escríbelo ahí."
		//
		// - Este uso de punteros es muy común cuando en Go una función necesita
		//   rellenar datos.
		err := rows.Scan(
			&pokemon.ID,
			&pokemon.Name,
			&pokemon.PrimaryTypeID,
			&pokemon.SecondaryTypeID,
			&pokemon.GenerationID,
			&pokemon.Sprite,
			&pokemon.Shiny,
		)
		if err != nil {
			return nil, err
		}

		// "append()" es el simil a "push()" en javascript.
		pokemonsList = append(pokemonsList, pokemon)
	}

	return pokemonsList, nil
}
