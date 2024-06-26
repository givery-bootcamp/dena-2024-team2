package myapp.usecase.interfaces

import myapp.entity.User

interface UserRepository {
    suspend fun createUser(name: String, password: String): User

    suspend fun findById(id: Int): User?

    suspend fun findByName(name: String): User?
}
