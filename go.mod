// Copyright SecureKey Technologies Inc. All Rights Reserved.
//
// SPDX-License-Identifier: Apache-2.0

module github.com/trustbloc/sidetree-fabric

require (
	github.com/Microsoft/hcsshim v0.8.7 // indirect
	github.com/Shopify/sarama v1.22.1 // indirect
	github.com/bluele/gcache v0.0.0-20190301044115-79ae3b2d8680
	github.com/btcsuite/btcutil v0.0.0-20170419141449-a5ecb5d9547a
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/gorilla/mux v1.7.3
	github.com/hyperledger/fabric v2.0.0+incompatible
	github.com/hyperledger/fabric-chaincode-go v0.0.0-20200128192331-2d899240a7ed
	github.com/hyperledger/fabric-protos-go v0.0.0-20200124220212-e9cfc186ba7b
	github.com/hyperledger/fabric-sdk-go v1.0.0-beta1.0.20200203184105-5f7f0b025d89
	github.com/hyperledger/fabric/extensions v0.0.0
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7
	github.com/pkg/errors v0.8.1
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/viper v1.3.2
	github.com/spf13/viper2015 v1.3.2
	github.com/stretchr/testify v1.4.0
	github.com/trustbloc/fabric-peer-ext v0.0.0
	github.com/trustbloc/sidetree-core-go v0.1.2-0.20200207213740-40ada19e1403
)

replace github.com/hyperledger/fabric => github.com/trustbloc/fabric-mod v0.1.2-0.20200211211900-62c249e072e2

replace github.com/hyperledger/fabric/extensions => github.com/trustbloc/fabric-peer-ext/mod/peer v0.0.0-20200211231405-e21ab2bc259f

replace github.com/trustbloc/fabric-peer-ext => github.com/trustbloc/fabric-peer-ext v0.1.2-0.20200211231405-e21ab2bc259f

replace github.com/hyperledger/fabric-protos-go => github.com/trustbloc/fabric-protos-go-ext v0.1.2-0.20200205170340-c69bba6d7b81

replace github.com/spf13/viper2015 => github.com/spf13/viper v0.0.0-20150908122457-1967d93db724

go 1.13
