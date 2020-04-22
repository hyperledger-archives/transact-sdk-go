// Copyright 2020 Tyson Foods, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// -----------------------------------------------------------------------------

package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/hyperledger/transact-sdk-go/sabre"
	"github.com/hyperledger/transact-sdk-go/sabre/addressing"
	"github.com/hyperledger/transact-sdk-go/src/protobuf/sabre_pb2"
	"github.com/hyperledger/transact-sdk-go/src/protobuf/transaction_pb2"
	"github.com/hyperledger/transact-sdk-go/transactions"
	t "github.com/hyperledger/transact-sdk-go/transactions"
	"github.com/hyperledger/transact-sdk-go/transactions/signing"
)

const XO_FAMILY_NAME = "xo"
const XO_FAMILY_PREFIX = "5b7349"

func createPayload(xoFamilyVersion, gameName, space string) (sabre.ISabrePayloadBuilder, error) {

	gameAddress := addressing.CalculateDeploymentAdddress(XO_FAMILY_NAME,
		gameName)

	payload := []byte(fmt.Sprintf("%s,take,%s", gameName, space))

	return sabre.NewSabrePayloadBuilder(
		sabre.WithAction(sabre_pb2.SabrePayload_EXECUTE_CONTRACT),
		sabre.WithContractName(XO_FAMILY_NAME),
		sabre.WithContractVersion(xoFamilyVersion),
		sabre.WithInputs([]string{gameAddress}),
		sabre.WithOutputs([]string{gameAddress}),
		sabre.WithExecuteContractPayload(payload),
	)
}

// Build transaction
func createTransaction(sabrePayloadBuilder sabre.ISabrePayloadBuilder,
	signer *signing.Signer) (*transaction_pb2.Transaction, error) {
	txnBuilder, err := t.NewTransactionBuilder()
	if err != nil {
		return nil, err
	}
	sabreTxnBuilder, err := sabre.NewSabreTransactionBuilder(
		sabre.WithPayloadBuilder(sabrePayloadBuilder),
		sabre.WithTransactionBuilder(txnBuilder),
	)

	return sabreTxnBuilder.Build(signer)
}

func createBatch(txns []*transaction_pb2.Transaction, signer *signing.Signer) ([]byte, error) {
	batchBuilder, err := transactions.NewBatchBuilder(
		transactions.WithTransactions(txns),
	)
	return batchBuilder.Build(signer)
}

type Request struct {
	CircuitID      string
	ServiceID      string
	SplinterHost   string
	GameName       string
	Space          string
	UserPrivateKey string
	XOVersion      string
}

func main() {

	var r Request
	flag.StringVar(&r.CircuitID, "circuit", "", "The circuit id")
	flag.StringVar(&r.ServiceID, "service", "", "The service id")
	flag.StringVar(&r.SplinterHost, "host", "localhost:8088", "The FQDN of the Splinter REST endpoint")
	flag.StringVar(&r.GameName, "game", "", "The xo game name")
	flag.StringVar(&r.Space, "space", "", "The space taken in the execute contract")
	flag.StringVar(&r.UserPrivateKey, "key", "", "The signing user's private key")
	flag.StringVar(&r.XOVersion, "xo_version", "0.3.3", "version of the XO contract")
	flag.Parse()

	endpoint := fmt.Sprintf("http://%s/scabbard/%s/%s/batches", r.SplinterHost, r.CircuitID, r.ServiceID)
	log.Println(endpoint)

	signer := signing.NewCryptoFactory(
		signing.CreateContext("secp256k1"),
	).NewSigner(
		signing.NewSecp256k1PrivateKey([]byte(r.UserPrivateKey)),
	)

	payload, err := createPayload(r.XOVersion, r.GameName, r.Space)
	if err != nil {
		log.Fatal(err)
	}
	txn, err := createTransaction(payload, signer)
	if err != nil {
		log.Fatal(err)
	}

	batches, err := createBatch([]*transaction_pb2.Transaction{txn}, signer)
	var b = bytes.NewBuffer(batches)

	resp, err := http.Post(endpoint, "application/octet-stream", b)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	message, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf(string(message))
}
