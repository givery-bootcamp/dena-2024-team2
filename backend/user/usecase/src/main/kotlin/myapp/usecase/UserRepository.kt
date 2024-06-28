package myapp.usecase

import myapp.User

interface UserRepository {
    suspend fun createUser(name: String, password: String): User

    suspend fun findById(id: Int): User?

    suspend fun findByName(name: String): User?
}
