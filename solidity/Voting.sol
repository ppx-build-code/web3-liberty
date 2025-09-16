// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Voting {

    mapping(address => uint) public tickets;
    address [] users;

    function vote(address to) public returns(bool) {
        uint ts = tickets[to];
        tickets[to] = ts + 1;
        users.push(to);
        return true;
    }

    function getVote(address to) public view returns(uint) {
        return tickets[to];
    }

    function resetVote() external {
        for(uint i =0; i < users.length; i ++){
            delete tickets[users[i]];
        }
    }
}