pragma solidity ^0.8.17;

contract ERC721 {
    event Transfer(address indexed _from, bytes32 indexed _to, uint tokens);
    
    function transfer(bytes32 _to) public payable {      
      emit Transfer(msg.sender, _to, msg.value);
    }
}