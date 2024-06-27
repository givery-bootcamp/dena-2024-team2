package myapp.usecase

import arrow.core.Either
import arrow.core.left
import arrow.core.right
import at.favre.lib.crypto.bcrypt.BCrypt
import com.auth0.jwt.JWT
import com.auth0.jwt.algorithms.Algorithm
import myapp.model.User
import myapp.repositories.UserRepository
import myapp.usecase.UserSigninUsecase.FailureReason
import org.koin.core.annotation.Single
import java.time.Instant

interface UserSigninUsecase {
    suspend fun login(username: String, password: String): Either<FailureReason, Pair<User, String>>

    enum class FailureReason {
        LOGIN_FAILED,
        INTERNAL_ERROR,
    }
}

@Single
@Suppress("unused")
class UserSigninUsecaseImpl(
    private val userRepository: UserRepository,
) : UserSigninUsecase {
    private val verifyer = BCrypt.verifyer()
    private val jwtSecret = requireNotNull(System.getenv("JWT_SECRET"))

    override suspend fun login(username: String, password: String): Either<FailureReason, Pair<User, String>> {
        try {
            val user = userRepository.findByName(username) ?: return FailureReason.LOGIN_FAILED.left()
            return if (verifyer.verify(password.toCharArray(), user.password).verified) {
                (user to genJwt(user)).right()
            } else {
                FailureReason.LOGIN_FAILED.left()
            }
        } catch (e: Exception) {
            e.printStackTrace()
            return FailureReason.INTERNAL_ERROR.left()
        }
    }

    private fun genJwt(user: User): String = JWT.create()
        .withSubject(user.id.toString())
        .withIssuer("Kodee")
        .withIssuedAt(Instant.now())
        .withExpiresAt(Instant.now().plusMillis(3_600_000)) // 1 hour
        .withClaim("username", user.name)
        .sign(Algorithm.HMAC256(jwtSecret))
}
