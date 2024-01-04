// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

library StateRootBuffer {
    struct BufferHeader {
        uint128 cursor;
        uint128 size;
    }

    struct Buffer {
        BufferHeader header;
        bytes32[] content;
    }

    function at(Buffer storage self, uint256 idx) internal view returns (bytes32) {
        uint128 _cursor = self.header.cursor;
        uint128 _size = self.header.size;

        require(idx < _size);

        return self.content[(idx + _cursor) % _size];
    }

    function atSlot(Buffer storage self, uint256 slot) internal view returns (bytes32) {
        uint128 _size = self.header.size;

        require(slot < _size);

        return self.content[slot];
    }

    function latestValue(Buffer storage self) internal view returns (uint256 slot, bytes32 value) {
        return (self.header.cursor, self.content[self.header.cursor]);
    }

    function insert(Buffer storage self, bytes32 value) internal {
        uint128 _cursor = self.header.cursor;
        uint128 _size = self.header.size;

        require(_size > 0);

        self.content[_cursor++] = value;

        if (_cursor >= _size) {
            self.header.cursor = 0;
        }
    }

    function size(Buffer storage self) internal view returns (uint256) {
        return self.header.size;
    }

    function cursor(Buffer storage self) internal view returns (uint256) {
        return self.header.cursor;
    }

    function initialize(Buffer storage self, uint128 newSize) internal {
        require(self.header.size == 0);

        self.header = BufferHeader(0, newSize);
    }
}
