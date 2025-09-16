// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract RomanToInt {


    mapping(bytes1 => uint) dict;

    constructor(){
        dict["I"] = 1;
        dict["V"] = 5;
        dict["X"] = 10;
        dict["L"] = 50;
        dict["C"] = 100;
        dict["D"] = 500;
        dict["M"] = 1000;

    }

    function convert(string memory roman) public view returns (uint) {
        bytes memory b = bytes(roman);
        uint result = 0;
        for(uint i = 0; i < b.length; i ++){
            uint v = dict[b[i]];
            result += v;
        }
        return result;
    }
}