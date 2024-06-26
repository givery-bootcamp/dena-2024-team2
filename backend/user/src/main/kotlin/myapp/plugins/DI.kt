package myapp.plugins

import io.ktor.server.application.*
import myapp.repositories.databaseConfig
import org.koin.core.annotation.ComponentScan
import org.koin.core.annotation.Module
import org.koin.core.annotation.Single
import org.koin.dsl.module
import org.koin.ksp.generated.module
import org.koin.ktor.plugin.Koin
import org.koin.logger.slf4jLogger

fun Application.configureKoin() {
    install(Koin) {
        slf4jLogger()
        modules(MyModule().module)
    }
}

@Module
@ComponentScan("myapp")
class MyModule {
    @Single
    fun database() = databaseConfig()
}
