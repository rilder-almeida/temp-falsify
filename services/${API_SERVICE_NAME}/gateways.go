package ${API_SERVICE_NAME}

// Repository is a ${API_SERVICE_NAME} adapter gateway for service layer 
type Repository interface {
    ReaderAdapter
    WriterAdapter
    EditorAdapter
}

// ReaderAdapter is a ${API_SERVICE_NAME} adapter gateway for service layer 
type ReaderAdapter interface {
    Get(interface{}) (interface{}, error)
    List() ([]interface{}, error)
}

// WriterAdapter is a ${API_SERVICE_NAME} adapter gateway for service layer 
type WriterAdapter interface {
    Create(interface{}) error
}

// EditorAdapter is a ${API_SERVICE_NAME} adapter gateway for service layer 
type EditorAdapter interface {
    Update(interface{}, interface{}) error
    Delete(interface{}) error
}