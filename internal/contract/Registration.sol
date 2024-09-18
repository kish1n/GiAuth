// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

contract Registration {
    mapping(address => bool) private users;

    event UserRegistered(address indexed user);
    event UserUnregistered(address indexed user);

    /// @notice Регистрирует отправителя как пользователя.
    function register() external {
        require(!users[msg.sender], "User is already registered");
        users[msg.sender] = true;
        emit UserRegistered(msg.sender);
    }

    /// @notice Удаляет регистрацию отправителя.
    function unregister() external {
        require(users[msg.sender], "User is not registered");
        delete users[msg.sender];
        emit UserUnregistered(msg.sender);
    }

    /// @notice Проверяет, зарегистрирован ли адрес.
    /// @param user Адрес для проверки.
    /// @return True, если пользователь зарегистрирован, иначе false.
    function isRegistered(address user) external view returns (bool) {
        return users[user];
    }
}
