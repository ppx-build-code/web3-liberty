// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract IntToRoman {

    mapping(uint => bytes1) dict;
    uint [] nums;

    constructor() {
        dict[1] = "I";
        dict[5] = bytes1('V');
        dict[10] = bytes1('X');
        dict[50] = bytes1('L');
        dict[100] = bytes1('C');
        dict[500] = bytes1('D');
        dict[1000] = bytes1('M');
        nums.push(1000);
        nums.push( 500);
        nums.push(100);
        nums.push(50);
        nums.push(10);
        nums.push(5);
        nums.push(1);
    }

    function intToRoman(uint num) public view returns (string memory) {
        bytes memory result = new bytes(100);
        uint v = num;
        uint ct = 0;
        for (uint8 i = 0; i< nums.length; i ++){
            uint romanCount = v / nums[i];
            v = v % nums[i];
            for(uint j =0; j < romanCount; j ++) {
                bytes1 c = dict[nums[i]];
                if (ct >= result.length) {
                    bytes memory result1 = result;
                    result = new bytes(result.length + 100);
                    for(uint k = 0; k < result1.length; k ++){
                        result[k] = result1[k];
                    }
                }
                result[ct] = c;
                ct ++;
            }
        }
        bytes memory fresult = new bytes(ct);
        for(uint i =0; i < ct; i ++){
            fresult[i] = result[i];
        }
        return string(fresult);
    }
}