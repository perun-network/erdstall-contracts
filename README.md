# Erdstall

A 2nd Layer Plasma based on abstract trusted third parties.

This project was launched at the ETHOnline 2020 hackathon.
The public hackathon code can be found at https://github.com/perun-network/erdstall/

## Solidity Contracts

This repository contains the public Ethereum Solidity contracts of the ongoing
development of Erdstall.

## Description

Erdstall leverages Trusted Execution Environments (TEE) like Intel SGX (or even
MPC committees) to scale Ethereum. Similar to Plasma or Rollups, the system
consists of a smart contract, an untrusted operator running a TEE and a dynamic
group of users. Joining and leaving the system is by a single call to a smart
contract. But once assets are deposited into the system, off-chain transactions
are free and only require the exchange of signatures from users to the operator.
The TEE Enclave receives and verifies those transactions and keeps track of the
system state. The whole system evolves in epochs and at the end of each epoch,
so-called balance proofs are distributed to all users, allowing them to leave
the system at any time. Those proofs are also necessary to give all users the
possibility to exit the system shall the operator decide to cease operating.

The underlying protocols were developed and proven secure by the Chair of
Applied Cryptography research group at Technical University Darmstadt (the same
team behind the Perun generalized state channels). The related paper is
currently in submission at a cryptography conference.

## License

This work is released under the Apache 2.0 license. See LICENSE file for more
details.

_Copyright (C) 2021 - PolyCrypt GmbH._
