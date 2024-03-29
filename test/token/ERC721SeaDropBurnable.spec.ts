import { expect } from "chai";
import { ethers, network } from "hardhat";

import { randomHex } from "./utils/encoding";
import { faucet } from "./utils/faucet";
import { VERSION } from "./utils/helpers";
import { whileImpersonating } from "./utils/impersonate";

import type { ERC721SeaDropBurnable, ISeaDrop } from "../../typechain-types";
import type { Wallet } from "ethers";

describe(`ERC721SeaDropBurnable (v${VERSION})`, function () {
  const { provider } = ethers;
  let seadrop: ISeaDrop;
  let token: ERC721SeaDropBurnable;
  let owner: Wallet;
  let creator: Wallet;
  let minter: Wallet;
  let approved: Wallet;

  after(async () => {
    await network.provider.request({
      method: "hardhat_reset",
    });
  });

  before(async () => {
    // Set the wallets
    owner = new ethers.Wallet(randomHex(32), provider);
    creator = new ethers.Wallet(randomHex(32), provider);
    minter = new ethers.Wallet(randomHex(32), provider);
    approved = new ethers.Wallet(randomHex(32), provider);

    // Add eth to wallets
    for (const wallet of [owner, minter, creator, approved]) {
      await faucet(wallet.address, provider);
    }

    // Deploy SeaDrop
    const SeaDrop = await ethers.getContractFactory("SeaDrop", owner);
    seadrop = await SeaDrop.deploy();
  });

  beforeEach(async () => {
    // Deploy token
    const ERC721SeaDropBurnable = await ethers.getContractFactory(
      "ERC721SeaDropBurnable",
      owner
    );
    token = await ERC721SeaDropBurnable.deploy("", "", [seadrop.address]);
  });

  it("Should only let the token owner burn their own token", async () => {
    await token.setMaxSupply(3);

    // Mint three tokens to the minter.
    await whileImpersonating(
      seadrop.address,
      provider,
      async (impersonatedSigner) => {
        await token.connect(impersonatedSigner).mintSeaDrop(minter.address, 3);
      }
    );

    expect(await token.ownerOf(1)).to.equal(minter.address);
    expect(await token.ownerOf(2)).to.equal(minter.address);
    expect(await token.ownerOf(3)).to.equal(minter.address);
    expect(await token.totalSupply()).to.equal(3);

    // Only the owner or approved of the minted token should be able to burn it.
    await expect(token.connect(owner).burn(1)).to.be.revertedWithCustomError(token,
      "TransferCallerNotOwnerNorApproved"
    );
    await expect(token.connect(approved).burn(1)).to.be.revertedWithCustomError(token,
      "TransferCallerNotOwnerNorApproved"
    );
    await expect(token.connect(approved).burn(2)).to.be.revertedWithCustomError(token,
      "TransferCallerNotOwnerNorApproved"
    );
    await expect(token.connect(owner).burn(3)).to.be.revertedWithCustomError(token,
      "TransferCallerNotOwnerNorApproved"
    );

    expect(await token.ownerOf(1)).to.equal(minter.address);
    expect(await token.ownerOf(2)).to.equal(minter.address);
    expect(await token.ownerOf(3)).to.equal(minter.address);
    expect(await token.totalSupply()).to.equal(3);

    await token.connect(minter).burn(1);

    expect(await token.totalSupply()).to.equal(2);

    await token.connect(minter).setApprovalForAll(approved.address, true);
    await token.connect(approved).burn(2);

    expect(await token.totalSupply()).to.equal(1);

    await token.connect(minter).setApprovalForAll(approved.address, false);
    await expect(token.connect(approved).burn(3)).to.be.revertedWithCustomError(token,
      "TransferCallerNotOwnerNorApproved"
    );

    await token.connect(minter).approve(owner.address, 3);
    await token.connect(owner).burn(3);

    expect(await token.totalSupply()).to.equal(0);

    await expect(token.ownerOf(1)).to.be.revertedWithCustomError(token,
      "OwnerQueryForNonexistentToken"
    );
    await expect(token.ownerOf(2)).to.be.revertedWithCustomError(token,
      "OwnerQueryForNonexistentToken"
    );
    expect(await token.totalSupply()).to.equal(0);

    // Should not be able to burn a nonexistent token.
    for (const tokenId of [0, 1, 2, 3]) {
      await expect(token.connect(minter).burn(tokenId)).to.be.revertedWithCustomError(token,
        "OwnerQueryForNonexistentToken"
      );
    }
  });
});
