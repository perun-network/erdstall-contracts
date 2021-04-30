import "@nomiclabs/hardhat-waffle";
import "@typechain/hardhat";

/**
 * @type import('hardhat/config').HardhatUserConfig
 */
export default {
  solidity: {
    version: "0.8.3",
    settings: {
      optimizer: {
        enabled: true,
        runs: 10000
      }
    }
  }
};
