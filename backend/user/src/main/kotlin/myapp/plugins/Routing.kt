package myapp.plugins

import arrow.core.Either
import io.ktor.http.ContentType
import io.ktor.http.HttpStatusCode
import io.ktor.server.application.Application
import io.ktor.server.application.call
import io.ktor.server.application.install
import io.ktor.server.plugins.autohead.AutoHeadResponse
import io.ktor.server.plugins.statuspages.StatusPages
import io.ktor.server.request.receive
import io.ktor.server.response.respond
import io.ktor.server.response.respondText
import io.ktor.server.routing.get
import io.ktor.server.routing.post
import io.ktor.server.routing.route
import io.ktor.server.routing.routing
import myapp.model.UserLogin
import myapp.model.UserPublic
import myapp.usecase.NewUserUsecase
import myapp.usecase.UserSigninUsecase
import org.koin.ktor.ext.inject

fun Application.configureRouting() {
    install(AutoHeadResponse)
    install(StatusPages) {
        exception<WrongRequestException> { call, error ->
            call.respond(HttpStatusCode.BadRequest, "${error.message}")
        }
        exception<Throwable> { call, cause ->
            call.respondText(text = "500: ${cause.message}", status = HttpStatusCode.InternalServerError)
        }
    }
    routing {
        val newUserUsecase by inject<NewUserUsecase>()
        val signinUsecase by inject<UserSigninUsecase>()
        get("/") {
            call.respondText("Hello World!", ContentType.Text.Plain)
        }

        route("/user") {
            post("/new") {
                val (name, password) = call.receive<UserLogin>()
                name.validate("User name length should between 4 and 32") { length in 4..32 }
                password.validate("Password length should between 6 and 64") { length in 8..64 }
                val user = newUserUsecase.createUser(name, password)
                call.respond(UserPublic(user.id, user.name))
            }
            post("/signin") {
                val (name, password) = call.receive<UserLogin>()
                when (val loginResult = signinUsecase.login(name, password)) {
                    is Either.Left -> {
                        when (loginResult.value) {
                            UserSigninUsecase.FailureReason.LOGIN_FAILED -> {
                                call.respondText("Username or Password is not correct", status = HttpStatusCode.BadRequest)
                            }
                            UserSigninUsecase.FailureReason.INTERNAL_ERROR -> {
                                call.respondText("Internal Server Error", status = HttpStatusCode.InternalServerError)
                            }
                        }
                    }
                    is Either.Right -> {
                        val (user, jwt) = loginResult.value
                        call.response.cookies.append(
                            "training-jwt",
                            jwt,
                            secure = true,
                            httpOnly = true,
                            extensions = mapOf("same-site" to "none"),
                        )
                        call.respond(UserPublic(user.id, user.name))
                    }
                }
            }
        }
    }
}

class WrongRequestException(message: String) : RuntimeException(message)

private inline fun <T> T.validate(message: String = "Request parameter is not valid", action: T.() -> Boolean) {
    runCatching {
        action()
    }.fold(onSuccess = {
        if (!it) throw WrongRequestException(message)
    }, onFailure = {
        throw WrongRequestException(message)
    })
}
