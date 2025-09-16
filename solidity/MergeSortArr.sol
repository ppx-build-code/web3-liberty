
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract MergeSortArr {

    function merge(uint[] memory arr1, uint[] memory arr2) public returns ( uint[] memory) {
        uint len = arr1.length + arr2.length;
        uint[] memory marr = new uint[](len);
        uint l = 0;
        for (uint i=0; i < arr1.length; i ++) {
            marr[l] = arr1[i];
            l ++;
        }
        for (uint j = 0; j < arr2.length; j ++) {
            marr[l] = arr2[j];
            l ++;
        }
        return merge1(marr);
     }

    function compute(uint x) public pure returns (uint) {
        return x / 2;
    }

    function merge1(uint[] memory arr) private returns (uint[]memory) {
        uint l = arr.length / 2;
        if(l < 1){
            return arr;
        }
        uint[] memory left = new uint[](l);
        uint[] memory right = new uint[](arr.length-l);
        for(uint i = 0; i < arr.length; i++) {
            if(i < l) {
                left[i] = arr[i];
            }else {
                right[i-l] = arr[i];
            }
        }
        uint [] memory lr = merge1(left);
        uint [] memory rr = merge1(right);

        return doMerge(lr, rr);
    }

    function doMerge(uint[] memory arr1, uint[] memory arr2) public pure returns ( uint[] memory) {
        uint len = arr1.length + arr2.length;
        uint [] memory result = new uint[](len);
        uint i = 0;
        uint j = 0;
        uint k = 0;
        while(i < arr1.length && j < arr2.length ){
            if (arr1[i] < arr2[j]) {
                result[k] = arr1[i];
                i++;
            } else {
                result[k] = arr2[j];
                j++;
            }
            k++;
        }
        while(i < arr1.length) {
            result[k] = arr1[i];
            i++;
            k++;
        }
        while(j < arr2.length) {
            result[k] = arr2[j];
            j++;
            k++;
        }
        return result;
    }
}