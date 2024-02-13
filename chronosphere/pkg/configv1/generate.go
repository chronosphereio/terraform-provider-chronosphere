package configv1

//go:generate swagger generate client --spec=./swagger.json --target=.
//go:generate mockgen -source=./client/collection/collection_client.go -destination=./mocks/collection_client_mock.go -package=mocks ClientService
