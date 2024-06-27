package myapp.repositories

import myapp.model.User
import org.jetbrains.exposed.sql.*
import org.jetbrains.exposed.sql.transactions.transaction
import org.koin.core.annotation.Single

interface UserRepository {
    suspend fun createUser(name: String, password: String): User

    suspend fun findById(id: Int): User?

    suspend fun findByName(name: String): User?
}

@Single
@Suppress("unused")
class UserRepositoryImpl(
    private val db: Database,
) : UserRepository {
    init {
        transaction(db) {
            SchemaUtils.createMissingTablesAndColumns(UserTable)
        }
    }

    @Suppress("PARAMETER_NAME_CHANGED_ON_OVERRIDE")
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

private object UserTable : Table("users") {
    val id = integer("id").autoIncrement()
    val name = varchar("name", 32).uniqueIndex()
    val password = char("password", 60)

    override val primaryKey: PrimaryKey = PrimaryKey(id)
}
