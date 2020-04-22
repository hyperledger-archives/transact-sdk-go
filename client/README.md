# Test Client Usage

The test client in this repository targets a running instance of the [Cargill/splinter/examples/gameroom](https://github.com/Cargill/splinter/tree/master/examples/gameroom) example splinter application running some version of the [xo smart contract](https://github.com/hyperledger/sawtooth-sdk-rust/tree/master/examples/xo_rust).

Steps
1. **Clone the Splinter repo** <br> This step assumes `SPLINTER_FOLDER` is set to a directory writable by the git user.
```shell
git clone git@github.com:Cargill/splinter.git $SPLINTER_FOLDER
```

2. **Run the gameroom example**
```shell
cd $SPLINTER_FOLDER
export CARGO_ARGS='-- --features experimental' # enable REST
docker-compose -f examples/gameroom/docker-compose.yaml build # this will take a while
docker-compose -f examples/gameroom/docker-compose.yaml up -d
```

3. **Get circuit and service id** <br> Be sure to find the service id that corresponds to the acme allowed node acme-node-000.
```shell
curl localhost:8088/admin/circuits #8088 is the default port for acme restful service

# > {"data":[{"id":"fcRqy-0YVaN","members":["bubba-node-000","acme-node-000"],"roster":[{"service_id":"gr00","service_type":"scabbard","allowed_nodes":["bubba-node-000"],"arguments":{"admin_keys":"[\"02372b3fe5e225c776d219b2f79ea480e330cd87f2965dc902e555e5e636856c94\"]","peer_services":"[\"gr01\"]"}},{"service_id":"gr01","service_type":"scabbard","allowed_nodes":["acme-node-000"],"arguments":{"admin_keys":"[\"02372b3fe5e225c776d219b2f79ea480e330cd87f2965dc902e555e5e636856c94\"]","peer_services":"[\"gr00\"]"}}],"management_type":"gameroom"}],"paging":{"current":"/admin/circuits?limit=100&offset=0","offset":0,"limit":100,"total":1,"first":"/admin/circuits?limit=100&offset=0","prev":"/admin/circuits?limit=100&offset=0","next":"/admin/circuits?limit=100&offset=0","last":"/admin/circuits?limit=100&offset=0"}}

export CIRCUIT_ID=fcRqy-0YVaN
export ACME_SERVICE_ID=gr01
```

4. **Get user private keys**
```shell
docker exec -it splinterd-node-acme /bin/bash
cat key_registry_shared/alice.priv
# > 5dd1641e1e434387a9b870f47af2eb6a2150209b2712856aa1e7fb78e1426a58
cat key_registry_shared/bob.priv
# > 14feb45a09a82031ea49ff0e50516d645aeceb2f95b5d0e2bfb30d7debceafc1
exit

export ALICE_KEY=5dd1641e1e434387a9b870f47af2eb6a2150209b2712856aa1e7fb78e1426a58
export BOB_KEY=14feb45a09a82031ea49ff0e50516d645aeceb2f95b5d0e2bfb30d7debceafc1
```

5. **Create users in UI** <br> Using the keys you grabbed in step 4, create the users in the UI for each gameroom application:

    - Using the endpoint for acme's UI (i.e. http://localhost:8080/), create the user alice with any email address and password. Input alice's private key.
    - Using the endpoint for bubba's UI (i.e. http://localhost:8081/), create the user bob with any email address and password. Input bob's private key.

6. **Create a gameroom and game in the UI** <br> Using either UI, create a gameroom with the opposite user. For example, if you use Acme's UI, create a gameroom with Bubba. Then, create a new XO game in that gameroom with any name. Remember the game name for use in the next step.

7. **Run the Test Client** <br> This example run assumes the game name is "Test".
```shell
export ACME_HOST='localhost:8088' # the rest client endpoint for acme splitnerd
export GAME_NAME='Test'
export TAKE_SPACE='3'

# Submit a batch list of a single trasnaction that executes the contract
# with a payload of 'Test,take,3' by user alice.
go run client/main.go \
  -host=$ACME_HOST \ # defaults to localhost:8088
  -circuit=$CIRCUIT_ID \
  -service=$ACME_SERVICE_ID \
  -game=$GAME_NAME \
  -space=$TAKE_SPACE \ # defaults to 1
  -key=$ALICE_KEY # Use Alice's key, created in Acme UI, to submit a batch to the ACME_SERVICE_ID
```
