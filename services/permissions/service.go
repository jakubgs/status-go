package permissions

import (
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/rpc"
)

// NewService initializes service instance.
func NewService(db *Database) *Service {
	return &Service{db: db}
}

type Service struct {
	db *Database
}

// Start a service.
func (s *Service) Start(*p2p.Server) error {
	return nil
}

// Stop a service.
func (s *Service) Stop() error {
	return nil
}

// APIs returns list of available RPC APIs.
func (s *Service) APIs() []rpc.API {
	return []rpc.API{
		{
			Namespace: "permissions",
			Version:   "0.1.0",
			Service:   NewAPI(s.db),
			Public:    true,
		},
	}
}

// Protocols returns list of p2p protocols.
func (s *Service) Protocols() []p2p.Protocol {
	return nil
}
