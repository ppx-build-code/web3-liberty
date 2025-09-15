// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0

Contract Voting {

    mapping(address => uint16) public ticketmappings;

    function vote(address to) public {
        
        uint16 ticketNum = ticketmappings[to];
        if(ticketNum != null) {
            ticketmappings[to] = ticketNum + 1;
        }else {
            ticketmappings[to] = 1;
        }
    }

    function getVote(candicate address) public uint16 {
        return ticketmappings[candicate];
    }

    function resetVote() public {
        ticketmappings = new mapping(address => uint16)();
    }



}