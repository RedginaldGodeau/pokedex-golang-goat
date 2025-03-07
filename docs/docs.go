// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/pokemon/": {
            "get": {
                "description": "Get all Pokemons",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pokemon"
                ],
                "summary": "Pokemon",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/pokemon.PokemonResponse"
                            }
                        }
                    }
                }
            }
        },
        "/api/pokemon/{id}/": {
            "get": {
                "description": "Get Pokemon by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pokemon"
                ],
                "summary": "Pokemon",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Pokemon ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pokemon.PokemonResponse"
                        }
                    }
                }
            }
        },
        "/api/pokemon/{name}/name/": {
            "get": {
                "description": "Get Pokemon by name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pokemon"
                ],
                "summary": "Pokemon",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Pokemon name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pokemon.PokemonResponse"
                        }
                    }
                }
            }
        },
        "/api/pokemon/{pokedex}/pokedex/": {
            "get": {
                "description": "Get Pokemon by Pokedex ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pokemon"
                ],
                "summary": "Pokemon",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Pokemon pokedex id",
                        "name": "pokedex",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/pokemon.PokemonResponse"
                            }
                        }
                    }
                }
            }
        },
        "/api/pokemon/{search}/search/": {
            "get": {
                "description": "Get Pokemon by name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pokemon"
                ],
                "summary": "Pokemon",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Pokemon name search",
                        "name": "search",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/pokemon.PokemonResponse"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "pokemon.PokemonResponse": {
            "type": "object",
            "properties": {
                "attack": {
                    "type": "integer"
                },
                "defense": {
                    "type": "integer"
                },
                "generation": {
                    "type": "integer"
                },
                "hp": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "legendary": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "pokedex": {
                    "type": "integer"
                },
                "sp_attack": {
                    "type": "integer"
                },
                "sp_defense": {
                    "type": "integer"
                },
                "speed": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                },
                "type_1": {
                    "type": "string"
                },
                "type_2": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Pokedex",
	Description:      "Un Pokédex en golang",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
