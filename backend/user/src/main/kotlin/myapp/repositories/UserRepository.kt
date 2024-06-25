package myapp.repositories

import myapp.model.User
import org.jetbrains.exposed.sql.*
import org.jetbrains.exposed.sql.transactions.transaction
import org.koin.core.annotation.Single

interface UserRepository {
    suspend fun createUser(name: String, password: String): User
    suspend fun findById(id: UInt): User?
}

@Single
class UserRepositoryImpl(
    private val db: Database
) : UserRepository {
    init {
        transaction(db) {
            SchemaUtils.createMissingTablesAndColumns(UserTable)
        }
    }

    override suspend fun createUser(userName: String, userPassword: String): User {
        val newUser = db.query {
            UserTable.insert {
                it[name] = userName
                it[password] = userPassword
            }
        }
        return User(
            id = newUser[UserTable.id],
            name = newUser[UserTable.name],
            password = newUser[UserTable.password]
        )
    }

    override suspend fun findById(id: UInt) : User? = db.query {
        UserTable.select { UserTable.id eq id }.map {
            User(
                id = it[UserTable.id],
                name = it[UserTable.name],
                password = it[UserTable.password]
            )
        }.singleOrNull()
    }
}

private object UserTable : Table() {
    val id = uinteger("id").autoIncrement()
    val name = varchar("name", 32)
    val password = char("password", 60)

    override val primaryKey: PrimaryKey = PrimaryKey(id)
}
