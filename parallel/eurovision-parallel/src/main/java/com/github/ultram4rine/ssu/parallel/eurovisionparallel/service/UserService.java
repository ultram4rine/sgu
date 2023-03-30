package com.github.ultram4rine.ssu.parallel.eurovisionparallel.service;

import com.github.ultram4rine.ssu.parallel.eurovisionparallel.dto.UserDto;
import com.github.ultram4rine.ssu.parallel.eurovisionparallel.entity.User;

import java.util.List;

public interface UserService {
    void saveUser(UserDto userDto);

    User findByEmail(String email);

    List<UserDto> findAllUsers();
}
