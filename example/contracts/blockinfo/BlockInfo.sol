// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract BlockInfo {
    function getBlockInfo()
        public
        view
        returns (
            uint256 currentBlockNumber,
            bytes32 currentBlockHash,
            uint256 previousBlockNumber,
            bytes32 previousBlockHash,
            bytes32[10] memory lastTenBlockHashes
        )
    {
        currentBlockNumber = block.number;
        currentBlockHash = blockhash(block.number);
        previousBlockNumber = block.number - 1;
        previousBlockHash = blockhash(block.number - 1);

        for (uint i = 0; i < 10; i++) {
            if (block.number > i) {
                lastTenBlockHashes[i] = blockhash(block.number - i);
            } else {
                lastTenBlockHashes[i] = bytes32(0); // Fill with zeroes if block number is out of range
            }
        }
    }
}