# Introduction To GraphQL

## Agenda Today

```graphql
    {
        graphql {
            what
            why
            how
        }
    }
```

## What is APIs
An API is a set of programming code that enables data transmission between one software product and another. It also contains the terms of this data exchange. [(altexsoft.com)](https://www.altexsoft.com/blog/engineering/what-is-api-definition-types-specifications-documentation)

## What is GraphQL

> GraphQL is a query language for APIs and a runtime for fulfilling those queries with your existing data. GraphQL provides a complete and understandable description of the data in your API, gives clients the power to ask for exactly what they need and nothing more, makes it easier to evolve APIs over time, and enables powerful developer tools. [(https://graphql.org/)](https://graphql.org/)

### Graphql Type System
graphql schema have many type system, such as scalar types of `String`, `Int`, `Float`, `Boolean`, and `ID`, so you can use these directly in the schema.
By default, every type in graphql schema is nullable, which means you can return or pass a null value.

examples of several types systems for other graphql schema:

- Object types
- Unions
- Enum
- Lists
- Scalars
- etc

### Graphql Operation
If in the REST APIs for each operation you will use different methods, for example there are GET, POST, PATCH, DELETE. In graphql, you will only use 2 operations:
1. Graphql Query
    
    A GraphQL query is used to read or fetch data.

    examples:
    ```graphql
    query GetAllPokemons{
        pokemons{
            id
            name
        }
    }
    ```

    or like this:
    ```graphql
    query GetPokemonWithName{
        pokemon(object: {name: "bulbasaur"}) {
            id
            name
        }
    }
    ```

    GraphQL queries help to reduce over fetching of data. Unlike a Restful API, GraphQL allows a user to restrict fields that should be fetched from the server. This means smaller queries and lesser traffic over the network; which in turn reduces the response time.

2. Graphql Mutation

    > In REST, any request might end up causing some side-effects on the server, but by convention it's suggested that one doesn't use GET requests to modify data. GraphQL is similar - technically any query could be implemented to cause a data write. However, it's useful to establish a convention that any operations that cause writes should be sent explicitly via a mutation.

    which means that all processes of writing or changing data in graphql must be done using a graphql mutation operation.

    example:

    ```graphql
    mutation AddNewPokemon {
        pokemon(object: {
            name: "Pikkachu"
        }) {
            id
            name
        }
    }
    ```

    or like this:

     ```graphql
    mutation AddNewPokemon($name: String!) {
        pokemon(object: {
            name: $name
        }) {
            id
            name
        }
    }
    ```

## Why is GraphQL Compelling ?

### Strongly-typed Schema
which helps determine the data that is available and the form it exists in. This, consequently, makes GraphQL less error-prone, and more validated, and provides auto-completion for supported IDE/editors.

### No Over-Fetching or Under-Fetching
With GraphQL, developers can fetch only what is required. Nothing less, nothing more. This solves the issues that arise due to over-fetching and under-fetching.

Over-fetching happens when the response fetches more than what is required.

Under-fetching, on the other hand, is not fetching adequate data in a single API request.
In this case, you need to make additional API requests to get related or referenced data. For instance, while displaying an individual blog post, you also need to fetch the referenced author’s entry, just so that you can display the author’s name and bio.

### Saves Time and Bandwidth

### Versioning Is Not Required

## How Queries Executed in GraphQL

Each field on each type is backed by a function called the resolver which is provided by the GraphQL server developer. When a field is executed, the corresponding resolver is called to produce the next value.

If a field produces a scalar value like a string or number, then the execution completes. However if a field produces an object value then the query will contain another selection of fields which apply to that object. This continues until scalar values are reached. GraphQL queries always end at scalar values.

examples:
```graphql
type Pokemons {
    id: String
    name: String
    pokemon_type: [PokemonType]
}

type PokemonType {
    name: String
}

query {
    pokemons {
        id
        name
        pokemon_type {
            name
        }
    }
}
```

## How do i run GraphQL server in Production ?

before you want to run graphql in production, you have to make sure of the following:

1. Security
2. Nested Queries


## Demo Time