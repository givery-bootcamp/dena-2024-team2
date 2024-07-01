package myapp.usecase

import arrow.core.Either
import arrow.core.left
import arrow.core.right
import at.favre.lib.crypto.bcrypt.BCrypt
import myapp.entity.User
import myapp.usecase.UserSigninUsecase.FailureReason
import myapp.usecase.interfaces.UserRepository

interface UserSigninUsecase {
    suspend fun login(username: String, password: String): Either<FailureReason, User>

    enum class FailureReason {
        LOGIN_FAILED,
        INTERNAL_ERROR,
    }
}

internal class UserSigninUsecaseImpl(
    private val userRepository: UserRepository,
) : UserSigninUsecase {
    private val verifyer = BCrypt.verifyer()

    override suspend fun login(username: String, password: String): Either<FailureReason, User> {
        try {
            val user = userRepository.findByName(username) ?: return FailureReason.LOGIN_FAILED.left()
            return if (verifyer.verify(password.toCharArray(), user.password).verified) {
                user.right()
            } else {
                FailureReason.LOGIN_FAILED.left()
            }
        } catch (e: Exception) {
            e.printStackTrace()
            return FailureReason.INTERNAL_ERROR.left()
        }
    }
}
