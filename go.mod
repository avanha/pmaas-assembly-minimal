module github.com/avanha/pmaas-assembly-minimal

go 1.27.1

require github.com/avanha/pmaas-core v0.0.5

require (
	github.com/avanha/pmaas-common v0.0.3 // indirect
	github.com/avanha/pmaas-spi v0.0.8 // indirect
)

replace github.com/avanha/pmaas-core => ../pmaas-core

replace github.com/avanha/pmaas-spi => ../pmaas-spi

replace github.com/avanha/pmaas-common => ../pmaas-common
