package transipClient

import "github.com/transip/gotransip/v6/domain"

var _domainRepo *domain.Repository

func getDomainRepo() *domain.Repository {
	if nil == _domainRepo {
		_domainRepo = &domain.Repository{Client: _client}
	}
	return _domainRepo
}
