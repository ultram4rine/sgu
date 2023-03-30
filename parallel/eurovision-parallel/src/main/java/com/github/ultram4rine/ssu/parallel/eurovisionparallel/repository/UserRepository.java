package com.github.ultram4rine.ssu.parallel.eurovisionparallel.repository;

import com.github.ultram4rine.ssu.parallel.eurovisionparallel.entity.User;
import org.springframework.data.jpa.repository.JpaRepository;

public interface UserRepository extends JpaRepository<User, Long> {
    User findByEmail(String email);
}
