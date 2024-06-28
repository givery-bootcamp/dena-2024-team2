package myapp.repository

import myapp.User
import myapp.usecase.interfaces.UserRepository
import org.jetbrains.exposed.sql.*
import org.jetbrains.exposed.sql.transactions.transaction

internal class UserRepositoryImpl(
    private val db: Database,
) : UserRepository {
    init {
        transaction(db) {
            SchemaUtils.createMissingTablesAndColumns(UserTable)
        }
    }

    @Suppress("PARAMETER_NAME_CHANGED_ON_OVERRIDE")
    override suspend fun createUser(userName: String, userPassword: String): User {
        val newUser =
            db.query {
                UserTable.insert {
                    it[name] = userName
                    it[password] = userPassword
                }
            }
        return User(
            id = newUser[UserTable.id],
            name = newUser[UserTable.name],
            password = newUser[UserTable.password],
        )
    }

    override suspend fun findById(id: Int): User? = db.query {
        UserTable.select { UserTable.id eq id }.map(::mapUser).singleOrNull()
    }

    override suspend fun findByName(name: String): User? = db.query {
        UserTable.select { UserTable.name eq name }.map(::mapUser).singleOrNull()
    }

    private fun mapUser(it: ResultRow) = User(
        id = it[UserTable.id],
        name = it[UserTable.name],
        password = it[UserTable.password],
    )
}
