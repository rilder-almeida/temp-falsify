POST http://pokedex_verse:8000/pokedex_verse/pokemon_dashboard/fetch_pokemon
{
    "Namespace": "exemplo-namespace",
    "FlowIdentifier": "system_subsystem_flow"
}

HTTP 200
[Asserts]
jsonpath "$.jobid" ==  ""


POST http://pokedex_verse:8000/pokedex_verse/pokemon_dashboard/fetch_pokemon
{
    "Namespace": "",
    "FlowIdentifier": "system_subsystem_flow"
}

HTTP 400
[Asserts]
jsonpath "$.error.code" ==  "BAD_REQUEST"
jsonpath "$.error.message" ==  "namespace is empty"


POST http://pokedex_verse:8000/pokedex_verse/pokemon_dashboard/fetch_pokemon
{
    "Namespace": "exemplo-namespace",
    "FlowIdentifier": "invalid_format"
}

HTTP 400
[Asserts]
jsonpath "$.error.code" ==  "BAD_REQUEST"
jsonpath "$.error.message" ==  "invalid flow identifier: invalid_format"
