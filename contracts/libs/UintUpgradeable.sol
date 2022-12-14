// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;


import "@openzeppelin/contracts-upgradeable/utils/math/SafeMathUpgradeable.sol";

library UintUpgradeable {
    using SafeMathUpgradeable for uint;
    function bp(uint value, uint bpValue) internal pure returns (uint) {
        return value.mul(bpValue).div(10000);
    }
}