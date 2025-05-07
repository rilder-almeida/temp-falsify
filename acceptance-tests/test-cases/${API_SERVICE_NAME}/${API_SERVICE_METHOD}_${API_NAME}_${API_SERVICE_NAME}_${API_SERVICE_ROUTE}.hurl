${API_SERVICE_METHOD} http://${API_NAME}:8000/${API_NAME}/${API_SERVICE_NAME}/${API_SERVICE_ROUTE}
{
    "Namespace": "exemplo-namespace",
    "FlowIdentifier": "system_subsystem_flow"
}

HTTP 200
[Asserts]
jsonpath "$.jobid" ==  ""


${API_SERVICE_METHOD} http://${API_NAME}:8000/${API_NAME}/${API_SERVICE_NAME}/${API_SERVICE_ROUTE}
{
    "Namespace": "",
    "FlowIdentifier": "system_subsystem_flow"
}

HTTP 400
[Asserts]
jsonpath "$.error.code" ==  "BAD_REQUEST"
jsonpath "$.error.message" ==  "namespace is empty"


${API_SERVICE_METHOD} http://${API_NAME}:8000/${API_NAME}/${API_SERVICE_NAME}/${API_SERVICE_ROUTE}
{
    "Namespace": "exemplo-namespace",
    "FlowIdentifier": "invalid_format"
}

HTTP 400
[Asserts]
jsonpath "$.error.code" ==  "BAD_REQUEST"
jsonpath "$.error.message" ==  "invalid flow identifier: invalid_format"
