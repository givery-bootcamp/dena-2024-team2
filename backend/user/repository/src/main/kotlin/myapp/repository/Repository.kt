package myapp.repository

import kotlinx.coroutines.Dispatchers
import myapp.usecase.UserRepository
import org.jetbrains.exposed.sql.Database
import org.jetbrains.exposed.sql.transactions.experimental.newSuspendedTransaction
import org.koin.dsl.bind
import org.koin.dsl.module

val repositoryModuke = module {
    single { UserRepositoryImpl(get()) }.bind(UserRepository::class)
}

internal suspend fun <T> Database.query(block: suspend () -> T): T = newSuspendedTransaction(Dispatchers.IO, this) { block() }
