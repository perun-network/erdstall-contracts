// SPDX-License-Identifier: Apache-2.0

import { ethers, waffle } from "hardhat";
import {
  Signer,
  constants, utils,
} from "ethers";
import chai from "chai";
import { solidity } from "ethereum-waffle";

import { Balance } from "@polycrypt/erdstall/api/responses";
import { Address } from "@polycrypt/erdstall/ledger";
import { Amount, Assets, Tokens } from "@polycrypt/erdstall/ledger/assets";

import {
  RevertToken__factory, RevertToken,
  PerunToken__factory, PerunToken,
  Erdstall__factory, Erdstall,
  ETHHolder__factory, ETHHolder,
  ERC20Holder__factory, ERC20Holder,
  ERC721MintableHolder__factory, ERC721MintableHolder,
} from "../typechain";

chai.use(solidity);
const { expect } = chai;


describe("Erdstall", () => {
  const EPOCHD = 12;
  const TEE = 0, OP = 1, ALICE = 2, BOB = 3;
  const BASE_URL = "https://api.erdstall.dev/token/";
  const REVERT_MOD = 618970019642690137449562111n; // 10th Mersenne Prime
  const provider = waffle.provider;

  let prn: PerunToken;
  let prnArt: RevertToken;
  let erdstall: Erdstall; // bound to Operator=Owner
  let ethHolder: ETHHolder; // bound to Operator=Owner
  let erc20Holder: ERC20Holder; // bound to Operator=Owner
  let erc721Holder: ERC721MintableHolder; // bound to Operator=Owner
  let accounts: Signer[];
  let teeAddr: string, aliceAddr: string, bobAddr: string;

  async function mineBlocks(n: number) {
    for (let i = 0; i < n; i++) {
      await provider.send("evm_mine", []);
    }
  }

  async function sealEpoch(epoch: number): Promise<number> {
    const bInit = (await erdstall.bigBang()).toNumber();
    const targetEpoch = epoch+2; // an epoch is sealed on-chain if the current block is two epochs further
    const targetBlock = bInit + targetEpoch * EPOCHD
    const currentBlock = await provider.getBlockNumber();
    const bdelta = targetBlock - currentBlock
    if (bdelta <= -EPOCHD)
      throw new Error(`Sealed epoch ${targetEpoch} already passed, current: ${(currentBlock-bInit)/EPOCHD}`)
    if (bdelta <= 0) {
      console.log(`Current sealed epoch already at ${targetEpoch}`);
      return bdelta;
    }
    await mineBlocks(bdelta);
    return bdelta;
  }

  before(async () => {
    accounts = await ethers.getSigners();
    teeAddr = await accounts[TEE].getAddress();
    aliceAddr = await accounts[ALICE].getAddress();
    bobAddr = await accounts[BOB].getAddress();
  });

  it("deploy the PerunToken ERC20", async () => {
    const prnFactory = (await ethers.getContractFactory("PerunToken", accounts[OP])) as PerunToken__factory;
    prn = await prnFactory.deploy(
      [aliceAddr, bobAddr],
      utils.parseEther("10000"),
    );
    await prn.deployed();
  });

  it("deploy Erdstall, TokenHolders, register them", async () => {
    const erdstallFactory = (await ethers.getContractFactory("Erdstall", accounts[OP])) as Erdstall__factory;
    erdstall = await erdstallFactory.deploy(teeAddr, EPOCHD);
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

  it("deploy and register the RevertToken ERC721Mintable", async function() {
    // 1. Deploy NFT Holder
    const erc721HolderFactory = (
      await ethers.getContractFactory(
        "ERC721MintableHolder",
        accounts[OP])
    ) as ERC721MintableHolder__factory;
    erc721Holder = await erc721HolderFactory.deploy(erdstall.address);
    await erc721Holder.deployed();

    // 2. Register NFT Holder for ERC721 tokens
    await expect(erdstall.registerTokenType(erc721Holder.address, "ERC721"))
      .to.emit(erdstall, "TokenTypeRegistered").withArgs("ERC721", erc721Holder.address);

    // 3. Deploy PerunArt with holder as minter
    const prnArtFactory = (await ethers.getContractFactory("RevertToken", accounts[OP])) as RevertToken__factory;
    prnArt = await prnArtFactory.deploy(BASE_URL, [erc721Holder.address], REVERT_MOD);
    await prnArt.deployed();

    // 4. Register PerunArt token contact
    await expect(erdstall.registerToken(prnArt.address, erc721Holder.address))
      .to.emit(erdstall, "TokenRegistered");
  });

  it("deposit 1 Ether [Alice]", async () => {
    const ethHolderAlice = ethHolder.connect(accounts[ALICE]);
    const amount = utils.parseEther("1");
    await expect(await ethHolderAlice.deposit({value: amount}))
      .to.changeEtherBalance(ethHolder, amount)
      .and.emit(erdstall, "Deposited");
  });

  it("deposit 10 PRN [Bob]", async () => {
    const amount = utils.parseEther("10");

    const prnBob = prn.connect(accounts[BOB]);
    await expect(prnBob.approve(erc20Holder.address, amount))
      .to.emit(prn, "Approval");

    const erc20HolderBob = erc20Holder.connect(accounts[BOB]);
    await expect(() =>
      expect(erc20HolderBob.deposit(prn.address, amount))
      .to.emit(prn, "Transfer")
      .and.emit(erdstall, "Deposited"))
      .to.changeTokenBalance(prn, erc20Holder, amount);
  });

  it("sealing epoch 0 should mine some blocks...", async function () {
    const bdelta = await sealEpoch(0);
    console.log(`Sealing 0 mined ${bdelta} blocks...`);
  });

  const aliceTknIds = [420n, 3000n];

  it("withdrawing 5 PRN and Erdstall-minted NFTs should mint them on-chain [Alice]", async function() {
    const amount = utils.parseEther("5");
    const bal = new Balance(
      0n, // epoch
      aliceAddr, // account
      true, // exit
      new Assets(
        { token: prnArt.address, asset: new Tokens(aliceTknIds) },
        { token: prn.address, asset: new Amount(amount.toBigInt()) },
      ),
    );
    const bp = await bal.sign(Address.fromString(erdstall.address), accounts[TEE]);

    const erdAlice = erdstall.connect(accounts[ALICE]);
    await expect(() =>
      expect(erdAlice.withdraw(...bp.toEthProof()))
      .to.emit(erdstall, "Withdrawn")
      .and.to.not.emit(erdstall, "WithdrawalException"))
      .to.changeTokenBalance(prn, accounts[ALICE], amount);

    for (const id of aliceTknIds) {
      expect(await prnArt.ownerOf(id)).to.equal(aliceAddr);
      expect(await prnArt.tokenURI(id)).to.equal(`${BASE_URL}${prnArt.address.toLowerCase()}/${id}`);
    }
  });

  it("deposit minted NFTs back into Erdstall [Alice]", async function() {
    const partAlice = prnArt.connect(accounts[ALICE]);
    for (const id of aliceTknIds) {
      await expect(partAlice.approve(erc721Holder.address, id))
        .to.emit(prnArt, "Approval");
    }

    const erc721HolderAlice = erc721Holder.connect(accounts[ALICE]);
    await expect(erc721HolderAlice.deposit(prnArt.address, aliceTknIds))
      .to.emit(prnArt, "Transfer")
      .and.emit(erdstall, "Deposited");
    for (const id of aliceTknIds) {
      expect(await prnArt.ownerOf(id)).to.equal(erc721Holder.address);
    }
  });

  it("withdrawing 5 PRN and Erdstall-minted reverting NFTs should only withdraw PRN [Bob]", async function() {
    const tknIds = [REVERT_MOD];
    const amount = utils.parseEther("5");
    const bal = new Balance(
      0n, // epoch
      bobAddr, // account
      true, // exit
      new Assets(
        { token: prnArt.address, asset: new Tokens(tknIds) },
        { token: prn.address, asset: new Amount(amount.toBigInt()) },
      ),
    );
    const bp = await bal.sign(Address.fromString(erdstall.address), accounts[TEE]);

    const erdBob = erdstall.connect(accounts[BOB]);
    await expect(() =>
      expect(erdBob.withdraw(...bp.toEthProof()))
      .to.emit(erdstall, "Withdrawn")
      .and.to.emit(erdstall, "WithdrawalException"))
      .to.changeTokenBalance(prn, accounts[BOB], amount);

    for (const id of tknIds) {
      expect(prnArt.ownerOf(id)).to.be.reverted;
    }
  });
});
