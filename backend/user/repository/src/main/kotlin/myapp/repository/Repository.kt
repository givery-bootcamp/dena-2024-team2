package myapp.repository

import kotlinx.coroutines.Dispatchers
import myapp.usecase.interfaces.UserRepository
import org.jetbrains.exposed.sql.Database
import org.jetbrains.exposed.sql.transactions.experimental.newSuspendedTransaction
import org.koin.dsl.bind
import org.koin.dsl.module

val repositoryModule = module {
    single { databaseConfig() }
    single { UserRepositoryImpl(get()) }.bind(UserRepository::class)
}

fun databaseConfig(
    dbUrl: String = System.getenv("DB_URL") ?: "jdbc:mysql://localhost:3306/training",
    user: String = System.getenv("DB_USER") ?: "root",
    password: String = System.getenv("DB_PASSWORD") ?: "",
) = Database.connect(
    url = dbUrl,
    user = user,
    password = password,
    driver = "com.mysql.cj.jdbc.Driver",
)

internal suspend fun <T> Database.query(block: suspend () -> T): T = newSuspendedTransaction(Dispatchers.IO, this) { block() }
