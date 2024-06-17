# `cosmoavs/statesync`

- A cosmos-sdk module which helps to retrive L1 on-chain state in the context of cosmos app.
- It facilites querying of operator set with stake for given block and quorum to be used for consensus and block advancement in cosmos chain.
- Implementation uses `OperatorStateRetriever` contract via eigensdk

This repository setup can also be used as template to boostrap other modules that require interacting with eigenlayer L1 contracts

## Future Improvements

- Add more methods to module to allow querying more complex logic
- Add simulation and more test cases
- Add support for ingesting event logs and checkpoint operator set in KVStore
- Integrate with other standard cosmos modules such as `x/epoch` and `x/evidence`
- Create modules which writes updates to cometbft's `EndBlock` validator set by integrating `x/staking` and `x/slashing` standard modules with AVS events
- Add support for BLS signature verification for CometBFT node

## Client

```proto
service Query {
  rpc OperatorState(QueryOperatorStateRequest)
      returns (QueryOperatorStateResponse) {
    option (cosmos.query.v1.module_query_safe) = true;
    option (google.api.http).get = "/cosmoavs/statesync/{block}/{quorum}";
  }
}
```

Supports default cosmos chain GRPC, REST and native method calling via keeper

## Params

```proto
message GenesisState {
  uint64 l1_chain_id = 1;
  uint64 l1_start_block = 2;

  string registry_coordinator = 3;
  string operator_state_retriever = 4;
}
```

## Tests

Basic test cases for genesis and state querying using `eigenda` mainnet deployment

```bash
cd ./x/statesync/keeper && RPC_URL="https://rpc.ankr.com/eth" go test
```
