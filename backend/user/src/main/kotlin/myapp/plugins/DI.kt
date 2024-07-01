package myapp.plugins

import io.ktor.server.application.*
import myapp.repository.repositoryModule
import myapp.usecase.usecaseModule
import org.koin.ktor.plugin.Koin
import org.koin.logger.slf4jLogger

fun Application.configureKoin() {
    install(Koin) {
        slf4jLogger()
        modules(repositoryModule, usecaseModule)
    }
}
