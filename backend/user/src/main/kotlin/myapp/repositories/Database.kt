package myapp.repositories

import kotlinx.coroutines.Dispatchers
import org.jetbrains.exposed.sql.Database
import org.jetbrains.exposed.sql.transactions.experimental.newSuspendedTransaction

fun databaseConfig(
    dbUrl: String = System.getenv("DB_URL") ?: "jdbc:mysql://localhost:3306/training",
    user: String = System.getenv("DB_USER") ?: "root",
    password: String = System.getenv("DB_PASSWORD") ?: ""
) = Database.connect(
    url = dbUrl,
    user = user,
    password = password,
    driver = "com.mysql.cj.jdbc.Driver",
)

suspend fun <T> Database.query(block: suspend () -> T): T =
    newSuspendedTransaction(Dispatchers.IO, this) { block() }