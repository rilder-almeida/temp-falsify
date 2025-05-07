package noopadapter

import (
	"github.com/rs/zerolog/log"
)

type NoOpAdapter struct{}

func NewNoOpAdapter() *NoOpAdapter {
	return &NoOpAdapter{}
}

func (n *NoOpAdapter) Get(id interface{}) (interface{}, error) {
	log.Info().Interface("id", id).Msg("Get operation called")
	return nil, nil
}

func (n *NoOpAdapter) List() ([]interface{}, error) {
	log.Info().Msg("List operation called")
	return nil, nil
}

func (n *NoOpAdapter) Create(entity interface{}) error {
	log.Info().Interface("entity", entity).Msg("Create operation called")
	return nil
}

func (n *NoOpAdapter) Update(id interface{}, entity interface{}) error {
	log.Info().Interface("id", id).Interface("entity", entity).Msg("Update operation called")
	return nil
}

func (n *NoOpAdapter) Delete(id interface{}) error {
	log.Info().Interface("id", id).Msg("Delete operation called")
	return nil
}
