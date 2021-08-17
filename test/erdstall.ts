// SPDX-License-Identifier: Apache-2.0

import { ethers } from "hardhat";
import {
  Signer,
  constants, utils,
} from "ethers";
import chai from "chai";
import { solidity } from "ethereum-waffle";

import {
  PerunToken__factory, PerunToken,
  Erdstall__factory, Erdstall,
  ETHHolder__factory, ETHHolder,
  ERC20Holder__factory, ERC20Holder,
} from "../typechain";

chai.use(solidity);
const { expect } = chai;

describe("Erdstall", () => {
  const TEE = 0, OP = 1, ALICE = 2, BOB = 3;

  let prn: PerunToken;
  let erdstall: Erdstall; // bound to Operator=Owner
  let ethHolder: ETHHolder; // bound to Operator=Owner
  let erc20Holder: ERC20Holder; // bound to Operator=Owner
  let accounts: Signer[];

  beforeEach(async () => {
    accounts = await ethers.getSigners();
  });

  it("deploy the PerunToken", async () => {
    const prnFactory = (await ethers.getContractFactory("PerunToken", accounts[OP])) as PerunToken__factory;
    prn = await prnFactory.deploy(
      [await accounts[ALICE].getAddress(), await accounts[BOB].getAddress()],
      utils.parseEther("10000"),
    );
    await prn.deployed();
  });

  it("deploy Erdstall, TokenHolders, register them", async () => {
    //const Erdstall = await ethers.getContractFactory("Erdstall");
    //const erdstall = await Erdstall.deploy("")

    const erdstallFactory = (await ethers.getContractFactory("Erdstall", accounts[OP])) as Erdstall__factory;
    erdstall = await erdstallFactory.deploy(await accounts[TEE].getAddress(), 12);
    await erdstall.deployed();

    const ethHolderFactory = (await ethers.getContractFactory("ETHHolder", accounts[OP])) as ETHHolder__factory;
    ethHolder = await ethHolderFactory.deploy(erdstall.address);
    await ethHolder.deployed();

    const erc20HolderFactory = (await ethers.getContractFactory("ERC20Holder", accounts[OP])) as ERC20Holder__factory;
    erc20Holder = await erc20HolderFactory.deploy(erdstall.address);
    await erc20Holder.deployed();

    await expect(erdstall.registerTokenType(ethHolder.address, "ETH"))
      .to.emit(erdstall, "TokenTypeRegistered").withArgs("ETH", ethHolder.address);
    await expect(erdstall.registerToken(constants.AddressZero, ethHolder.address))
      .to.emit(erdstall, "TokenRegistered");

    await expect(erdstall.registerTokenType(erc20Holder.address, "ERC20"))
      .to.emit(erdstall, "TokenTypeRegistered").withArgs("ERC20", erc20Holder.address);
    await expect(erdstall.registerToken(prn.address, erc20Holder.address))
      .to.emit(erdstall, "TokenRegistered");
  });

  it("deposit 1 Ether by Alice", async () => {
    const ethHolderAlice = ethHolder.connect(accounts[ALICE]);
    const amount = utils.parseEther("1");
    await expect(await ethHolderAlice.deposit({value: amount}))
      .to.changeEtherBalance(ethHolder, amount)
      .and.emit(erdstall, "Deposited");
  });

  it("deposit 10 PRN by Bob", async () => {
    const amount = utils.parseEther("10");

    const prnBob = prn.connect(accounts[BOB]);
    await expect(prnBob.approve(erc20Holder.address, amount))
      .to.emit(prn, "Approval");

    const erc20HolderBob = erc20Holder.connect(accounts[BOB]);
    await expect(await erc20HolderBob.deposit(prn.address, amount))
      .to.changeTokenBalance(prn, erc20Holder, amount)
      .and.emit(prn, "Transfer")
      .and.emit(erdstall, "Deposited");
  });

});
