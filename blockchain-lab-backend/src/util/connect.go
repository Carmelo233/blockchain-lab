/*
Copyright 2021 IBM All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package util

import (
	"blockchain-lab/util/conf"
	"crypto/x509"
	"fmt"
	"os"
	"path"

	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// func NewConnection(conf conf.Connconf) *client.Contract {

// 	// The gRPC client connection should be shared by all Gateway connections to this endpoint
// 	clientConnection := newGrpcConnection(conf)
// 	// defer clientConnection.Close()

// 	id := newIdentity(conf)
// 	sign := newSign(conf)

// 	// Create a Gateway connection for a specific client identity
// 	gw, err := client.Connect(
// 		id,
// 		client.WithSign(sign),
// 		client.WithClientConnection(clientConnection),
// 		// Default timeouts for different gRPC calls
// 		client.WithEvaluateTimeout(5*time.Second),
// 		client.WithEndorseTimeout(15*time.Second),
// 		client.WithSubmitTimeout(5*time.Second),
// 		client.WithCommitStatusTimeout(1*time.Minute),
// 	)
// 	if err != nil {
// 		panic(err)
// 	}
// 	// defer gw.Close()

// 	// Override default values for chaincode and channel name as they may differ in testing contexts.
// 	chaincodeName := "basic"
// 	if ccname := os.Getenv("CHAINCODE_NAME"); ccname != "" {
// 		chaincodeName = ccname
// 	}

// 	channelName := "mychannel"
// 	if cname := os.Getenv("CHANNEL_NAME"); cname != "" {
// 		channelName = cname
// 	}

// 	network := gw.GetNetwork(channelName)
// 	contract := network.GetContract(chaincodeName)

// 	return contract
// }

// newGrpcConnection creates a gRPC connection to the Gateway server.
func NewGrpcConnection(conf conf.Connconf) *grpc.ClientConn {
	certificate, err := LoadCertificate(conf.TlsCertPath)
	if err != nil {
		panic(err)
	}

	certPool := x509.NewCertPool()
	certPool.AddCert(certificate)
	transportCredentials := credentials.NewClientTLSFromCert(certPool, conf.GatewayPeer)

	connection, err := grpc.Dial(conf.PeerEndpoint, grpc.WithTransportCredentials(transportCredentials))
	if err != nil {
		panic(fmt.Errorf("failed to create gRPC connection: %w", err))
	}

	return connection
}

// newIdentity creates a client identity for this Gateway connection using an X.509 certificate.
func NewIdentity(conf conf.Connconf) *identity.X509Identity {
	certificate, err := LoadCertificate(conf.CertPath)
	if err != nil {
		panic(err)
	}

	id, err := identity.NewX509Identity(conf.MspID, certificate)
	if err != nil {
		panic(err)
	}

	return id
}

func LoadCertificate(filename string) (*x509.Certificate, error) {
	certificatePEM, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read certificate file: %w", err)
	}
	return identity.CertificateFromPEM(certificatePEM)
}

// newSign creates a function that generates a digital signature from a message digest using a private key.
func NewSign(conf conf.Connconf) identity.Sign {
	files, err := os.ReadDir(conf.KeyPath)
	if err != nil {
		panic(fmt.Errorf("failed to read private key directory: %w", err))
	}
	privateKeyPEM, err := os.ReadFile(path.Join(conf.KeyPath, files[0].Name()))

	if err != nil {
		panic(fmt.Errorf("failed to read private key file: %w", err))
	}

	privateKey, err := identity.PrivateKeyFromPEM(privateKeyPEM)
	if err != nil {
		panic(err)
	}

	sign, err := identity.NewPrivateKeySign(privateKey)
	if err != nil {
		panic(err)
	}

	return sign
}
