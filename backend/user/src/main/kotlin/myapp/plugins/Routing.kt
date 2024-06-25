package myapp.plugins

import io.ktor.http.ContentType
import io.ktor.http.HttpStatusCode
import io.ktor.server.application.Application
import io.ktor.server.application.call
import io.ktor.server.application.install
import io.ktor.server.plugins.autohead.AutoHeadResponse
import io.ktor.server.plugins.statuspages.StatusPages
import io.ktor.server.request.receive
import io.ktor.server.response.respondText
import io.ktor.server.routing.get
import io.ktor.server.routing.post
import io.ktor.server.routing.route
import io.ktor.server.routing.routing
import myapp.model.UserLogin
import myapp.usecase.NewUserUsecase
import org.koin.ktor.ext.inject

fun Application.configureRouting() {
    install(AutoHeadResponse)
    install(StatusPages) {
        exception<Throwable> { call, cause ->
            call.respondText(text = "500: $cause", status = HttpStatusCode.InternalServerError)
        }
    }
    routing {
        val newUserUsecase by inject<NewUserUsecase>()
        get("/") {
            call.respondText("Hello World!", ContentType.Text.Plain)
        }
        route("/user") {
            post("/new") {
                val (name, password) = call.receive<UserLogin>()
                newUserUsecase.createUser(name, password)
            }
        }
    }
}
