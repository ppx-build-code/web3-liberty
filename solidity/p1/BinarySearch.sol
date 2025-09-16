// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract BinarySearch {

    function binarySearch(uint[] memory nums, uint target) public pure returns (uint) {

        uint start = 0;
        uint end = nums.length - 1;
        while (start < end) {
            uint mid = (end - start) / 2 + start;
            if(target == nums[mid]) {
                return mid;
            } else if (target > nums[mid]) {
                start = mid + 1;
            }else {
                end = mid;
            }
        }
        return 0;
    }
}