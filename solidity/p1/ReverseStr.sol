// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract ReverseStr {
    
    function reverseStr(string memory str) public pure returns (string memory) {
        bytes memory b = bytes(str);
        bytes memory nb = new bytes(b.length);
        
        for(uint i = b.length; i > 0; i --) {
            nb[b.length - i] = b[i-1];
            
        }
        return string(nb);
    }
}