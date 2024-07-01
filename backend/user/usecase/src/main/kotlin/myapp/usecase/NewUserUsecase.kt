package myapp.usecase

import at.favre.lib.crypto.bcrypt.BCrypt
import myapp.entity.User
import myapp.usecase.interfaces.UserRepository
import org.koin.core.annotation.Single

interface NewUserUsecase {
    suspend fun createUser(name: String, password: String): User
}

@Single
internal class NewUserUsecaseImpl(
    private val userRepository: UserRepository,
) : NewUserUsecase {
    private val hasher = BCrypt.withDefaults()

    override suspend fun createUser(name: String, password: String): User {
        val passwordEncrypt = hasher.hashToString(10, password.toCharArray())
        return userRepository.createUser(name, passwordEncrypt)
    }
}
