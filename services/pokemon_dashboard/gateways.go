package pokemon_dashboard

// Repository is a pokemon_dashboard adapter gateway for service layer 
type Repository interface {
    ReaderAdapter
    WriterAdapter
    EditorAdapter
}

// ReaderAdapter is a pokemon_dashboard adapter gateway for service layer 
type ReaderAdapter interface {
    Get(interface{}) (interface{}, error)
    List() ([]interface{}, error)
}

// WriterAdapter is a pokemon_dashboard adapter gateway for service layer 
type WriterAdapter interface {
    Create(interface{}) error
}

// EditorAdapter is a pokemon_dashboard adapter gateway for service layer 
type EditorAdapter interface {
    Update(interface{}, interface{}) error
    Delete(interface{}) error
}